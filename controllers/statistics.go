package controllers

import (
	"math"
	"offergo/lib"
	"offergo/models"
	"reflect"
	"sync"
)

type StatisticsController struct {
	baseController
}
type statisticalFunc func(*lib.GetUserStatisticalResponseData) (bool, interface{})

func (s *StatisticsController) GetUserStatistical() {
	//当前请求需要用到goroutine的方法集合
	var funcs = map[string]statisticalFunc{
		"WxUser":                       s.getWechatUserInfo,                   //微信用户信息
		"AppUser":                      s.getAllAppUserInfo,                   //总APP用户信息
		"IosUser":                      s.getIosUserInfo,                      //ios用户信息
		"AndroidUser":                  s.getAndroidUserInfo,                  //android用户信息
		"JobBookUser":                  s.getJobUserInfo,                      //工作板块用户信息
		"ActiveBookUser":               s.getActivityUserInfo,                 //活动板块用户信息
		"MbUser":                       s.getParcelUserInfo,                   //集运信息
		"SecondHandUser":               s.getSecondHandUserInfo,               //二手用户信息
		"OfflineCommissionPublishUser": s.getOfflineCommissionPublishUserInfo, //待寄待取发布者用户信息
		"OfflineCommissionPayUser":     s.getOfflineCommissionPayUserInfo,     //待寄待取付费者用户信息
		"TelecomCardActivationedUser":  s.getTelecomCardActivationedUserInfo,  //大K卡已激活总人数信息
		"TelecomUsingUser":             s.getTelecomCardUseingUserInfo,        //正在使用的大k卡用户信息
		"TelecomNewUser":               s.getTelecomCardNewUserInfo,           //新申请的大k卡用户信息
		"KPlusUser":                    s.getKPlusUserInfo,                    //k_plus会员用户信息
	}
	//定义响应结构体
	var result lib.GetUserStatisticalResponseData
	//获取结构体字段数量
	resultNum := reflect.TypeOf(result).NumField()
	/**
	错误接收需要内存安全所以使用通道
	设置有缓存通道，缓存数量与响应结构体相同
	**/
	errorChannel := make(chan string, resultNum)
	// 声明一个等待组
	var wg sync.WaitGroup
	// 设置等待组数量
	wg.Add(resultNum)

	//遍历方法集合
	for _, v := range funcs {
		//开启goroutine
		go func(v statisticalFunc) {
			defer wg.Done()
			ok, data := v(&result)
			if !ok {
				//输出错误信息给通道
				err := data.(string)
				errorChannel <- err
			}
		}(v)
	}

	wg.Wait()
	//关闭通道防止无限for
	close(errorChannel)
	//查看通道是否有错误
	if len(errorChannel) > 0 {
		//获取所有错误信息
		//存放错误字符串
		var err string
		for data := range errorChannel {
			err += data
		}
		s.responseError(err)

	}
	s.responseSuccess(result)
}

//二手市场用户信息
func (s *StatisticsController) getSecondHandUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取二手用户数据
	//获取当前二手用户数量
	sel := []string{"openid"}
	where := models.SecondHandInfo{}
	//拿到当前二手用户数量
	var currentUser int
	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["openid IS NOT NULL and openid != ?"] = ""
	option["wheres"] = wheres
	option["group"] = "openid"
	option["count"] = &currentUser
	result := s.getDBSecondHandInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号二手用户数量

	sel = []string{"openid"}
	where = models.SecondHandInfo{}
	//拿到当月1号二手用户数量
	var currentMonthUser int
	//option
	wheres["post_at <= ?"] = lib.MonthOneDay()
	wheres["openid IS NOT NULL and openid != ?"] = ""
	option["wheres"] = wheres
	option["group"] = "openid"
	option["count"] = &currentMonthUser

	result = s.getDBSecondHandInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取上个月1号二手数量
	sel = []string{"openid"}
	where = models.SecondHandInfo{}
	//拿到上个月1号二手用户数量
	var lastMonthUser int
	//option
	wheres["post_at <= ?"] = lib.LastMonthOneDay()
	wheres["openid IS NOT NULL and openid != ?"] = ""
	option["wheres"] = wheres
	option["group"] = "openid"
	option["count"] = &lastMonthUser
	result = s.getDBSecondHandInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.SecondHandUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "二手市场用户",
	}
	return true, "ok"
}

//工作板块用户信息
func (s *StatisticsController) getJobUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取工作板块用户数据
	sel := []string{"user_id"}
	where := models.User{}
	var CurrentUser int
	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["job_subscriber = ?"] = "1"
	wheres["openid IS NOT NUll"] = nil
	option["wheres"] = wheres
	option["count"] = &CurrentUser
	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}
	data.JobBookUser = lib.GetUserStatisticalStruct{
		CurrentUser: CurrentUser,
		Text:        "工作板块订阅用户",
	}
	return true, "ok"
}

//活动板块用户信息
func (s *StatisticsController) getActivityUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取工作板块用户数据
	sel := []string{"user_id"}
	where := models.User{}
	var CurrentUser int
	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["activity_subscriber = ?"] = "1"
	wheres["openid IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &CurrentUser
	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}
	data.ActiveBookUser = lib.GetUserStatisticalStruct{
		CurrentUser: CurrentUser,
		Text:        "活动板块订阅用户",
	}
	return true, "ok"
}

//ios用户信息
func (s *StatisticsController) getIosUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取ios用户数据
	//获取当前ios用户数量
	sel := []string{"user_id"}
	where := models.User{}
	//拿到当前ios用户数量
	var currentUser int

	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["iosdeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser

	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号ios用户数量

	sel = []string{"user_id"}
	where = models.User{}
	//拿到当月1号ios用户数量
	var currentMonthUser int
	//option
	wheres["created_at <= ?"] = lib.MonthOneDay()
	wheres["iosdeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentMonthUser

	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取上个月1号ios数量
	sel = []string{"user_id"}
	where = models.User{}
	//拿到上个月1号ios用户数量
	var lastMonthUser int
	//option
	wheres["created_at <= ?"] = lib.LastMonthOneDay()
	wheres["iosdeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &lastMonthUser

	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}
	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.IosUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "IOS用户",
	}
	return true, "ok"
}

//android用户信息
func (s *StatisticsController) getAndroidUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取android用户数据
	//获取当前android用户数量
	sel := []string{"user_id"}
	where := models.User{}
	//拿到当前android用户数量
	var currentUser int
	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["androiddeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号android用户数量

	sel = []string{"user_id"}
	where = models.User{}
	//拿到当月1号android用户数量
	var currentMonthUser int
	//option
	wheres["created_at <= ?"] = lib.MonthOneDay()
	wheres["androiddeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取上个月1号android数量
	sel = []string{"user_id"}
	where = models.User{}
	//拿到上个月1号android用户数量
	var lastMonthUser int
	//option
	wheres["created_at <= ?"] = lib.LastMonthOneDay()
	wheres["androiddeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &lastMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.AndroidUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "Android用户",
	}
	return true, "ok"
}

//微信用户信息
func (s *StatisticsController) getWechatUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	//获取微信用户数据
	//获取当前用户数量
	sel := []string{"user_id"}
	where := models.User{}
	//拿到当前微信用户数量
	var currentUser int
	//option
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	wheres["openid IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取当月1号微信用户数量

	sel = []string{"user_id"}
	where = models.User{}
	//拿到当月1号微信用户数量
	var currentMonthUser int
	//option
	wheres["created_at <= ?"] = lib.MonthOneDay()
	wheres["openid IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取上个月1号微信数量
	sel = []string{"user_id"}
	where = models.User{}
	//拿到上个月1号微信用户数量
	var lastMonthUser int
	//option
	wheres["created_at <= ?"] = lib.LastMonthOneDay()
	wheres["openid IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &lastMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.WxUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "微信用户",
	}
	return true, "ok"
}

//获取集运用户信息
func (s *StatisticsController) getParcelUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {

	result := (&models.SzUser{}).GetStatisticsSzUser(lib.MonthOneDay(), lib.LastMonthOneDay())
	//字典转换为结构体
	//转换为结构体赋值
	tmp := interface{}(&data.MbUser)
	lib.MapToStruct(result, &tmp)
	return true, "ok"
}

//获取app总用户数量
func (s *StatisticsController) getAllAppUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"user_id"}
	where := models.User{}
	//当前用户数量
	var currentUser int
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取当前最新App用户数量
	//option
	wheres["iosdeviceToken IS NOT NULL or androiddeviceToken IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号APP用户数量

	sel = []string{"user_id"}
	where = models.User{}
	//拿到当月1号App用户数量
	var currentMonthUser int
	//option
	wheres["created_at <= ?"] = lib.MonthOneDay()
	option["wheres"] = wheres
	option["count"] = &currentMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取上个月1号APP用户数量
	sel = []string{"user_id"}
	where = models.User{}
	//拿到上个月1号APP用户数量
	var lastMonthUser int
	//option
	wheres["created_at <= ?"] = lib.LastMonthOneDay()
	option["wheres"] = wheres
	option["count"] = &lastMonthUser
	result = s.getDBUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.AppUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "App用户",
	}
	return true, "ok"
}

//获取待寄待取发布者用户信息
func (s *StatisticsController) getOfflineCommissionPublishUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"id"}
	//获取发布者数量
	var currentUser int
	where := models.OfflineCommissionVerify{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取所有发布者数量信息
	where.VerifyStatus = 3
	wheres["completion_times > ?"] = 0
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBOfflineCommissionPublishUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	data.OfflineCommissionPublishUser = lib.GetUserStatisticalStruct{
		CurrentUser: currentUser,
		Text:        "待寄待取发布者",
	}
	return true, "ok"
}

//获取待寄待取付费者用户信息
func (s *StatisticsController) getOfflineCommissionPayUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"distinct(user_id)"}
	//获取当前所有付费者数量信息
	var currentUser int
	where := models.OfflineCommissionOrder{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取当前所有付费者数量信息
	wheres["status in (?)"] = []int{1, 2}
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBOfflineCommissionPayUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取当月1号所有付费者数量信息
	//获取当月1号所有付费者数量信息
	var currentMonthUser int
	wheres["status in (?)"] = []int{1, 2}
	wheres["created_at <= ?"] = lib.MonthOneDayUnix()
	where = models.OfflineCommissionOrder{}
	option["wheres"] = wheres
	option["count"] = &currentMonthUser
	result = s.getDBOfflineCommissionPayUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取上个月1号所有付费者数量信息
	//拿到上个月1号所有付费者数量信息
	var lastMonthUser int
	where = models.OfflineCommissionOrder{}
	//option
	wheres["created_at <= ?"] = lib.LastMonthOneDayUnix()
	option["wheres"] = wheres
	option["count"] = &lastMonthUser
	result = s.getDBOfflineCommissionPayUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.OfflineCommissionPayUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "待寄待取付费用户",
	}
	return true, "ok"
}

//获取大K卡正在使用的人数
func (s *StatisticsController) getTelecomCardUseingUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"id"}
	where := models.TelecomUserCard{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取当前所有付费者数量数量
	var currentUser int
	//获取当前所有付费者数量信息
	where.CardStatus = 1
	wheres["activate_date IS NOT NULL"] = nil
	wheres["type != ?"] = "TEST"
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBTelecomCardUsingUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}
	data.TelecomUsingUser = lib.GetUserStatisticalStruct{
		CurrentUser: currentUser,
		Text:        "大K卡正在使用人数",
	}
	return true, "ok"
}

//获取大K卡新申请的人数
func (s *StatisticsController) getTelecomCardNewUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"user_card.id"}
	where := models.TelecomUserCard{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	join := make(map[string]interface{})
	//获取当前新申请的人数数量
	var currentUser int
	//获取当前新申请的人数信息
	startTime := lib.MonthOneDay()
	//当月最后一天时间戳
	endTime := lib.MonthLastDay()
	wheres["applications.apply_status >= ?"] = 3
	wheres["applications.apply_status != ?"] = -1
	wheres["user_card.type = ?"] = "BUY"
	wheres["applications.created_at >= ?"] = startTime
	wheres["applications.created_at <= ?"] = endTime

	join["join `applications` on applications.user_id = user_card.user_id"] = nil
	option["wheres"] = wheres
	option["join"] = join
	option["count"] = &currentUser
	result := s.getDBTelecomCardUsingUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	data.TelecomNewUser = lib.GetUserStatisticalStruct{
		CurrentUser: currentUser,
		Text:        "大K卡新申请人数",
	}
	return true, "ok"
}

//获取大K卡已激活的总人数
func (s *StatisticsController) getTelecomCardActivationedUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"user_card.id"}
	where := models.TelecomUserCard{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取当前已激活大k卡的总人数信息
	//获取当前已激活的的总人数数量
	var currentUser int
	wheres["type != ?"] = "TEST"
	wheres["activate_date IS NOT NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBTelecomCardUsingUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号已激活大k卡用户数量
	var currentMonthUser int
	wheres["created_at <= ?"] = lib.MonthOneDayUnix()
	option["count"] = &currentMonthUser
	result = s.getDBTelecomCardUsingUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取上月1号已激活大k卡用户数量
	var lastMonthUser int
	wheres["created_at <= ?"] = lib.LastMonthOneDayUnix()
	option["count"] = &lastMonthUser
	result = s.getDBTelecomCardUsingUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}


	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.TelecomCardActivationedUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "大K卡已激活总人数",
	}
	return true, "ok"
}

//获取k_plus会员数量
func (s *StatisticsController) getKPlusUserInfo(data *lib.GetUserStatisticalResponseData) (bool, interface{}) {
	sel := []string{"id"}
	//获取当前已有的k+plus总人数数量
	var currentUser int
	where := models.KPlusOrder{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	//获取当前的k+plus会员人数信息
	wheres["pay_status = ?"] = 1
	wheres["deleted_at IS NULL"] = nil
	option["wheres"] = wheres
	option["count"] = &currentUser
	result := s.getDBKPlusUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取当月1号k+plus用户数量
	var currentMonthUser int
	wheres["created_at <= ?"] = lib.MonthOneDayUnix()
	option["count"] = &currentMonthUser
	result = s.getDBKPlusUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取上月1号k+plus用户数量
	var lastMonthUser int
	wheres["created_at <= ?"] = lib.LastMonthOneDayUnix()
	option["count"] = &lastMonthUser
	result = s.getDBKPlusUserInfo(sel, where, option)
	if result.Code == 400 {
		return false, result.Msg
	}

	//获取增长率
	percentage := s.getChance(currentMonthUser, lastMonthUser)
	data.KPlusUser = lib.GetUserStatisticalStruct{
		CurrentUser:      currentUser,
		CurrentMonthUser: currentMonthUser,
		LastMonthUser:    lastMonthUser,
		Percentage:       percentage,
		Text:             "k+plus会员总人数",
	}
	return true, "ok"
}

func (s *StatisticsController) getDBUserInfo(sel []string, where models.User, option map[string]interface{}) result {
	//获取微信用户信息
	//result user
	var resultStruct []models.User

	result, ok := new(models.User).GetUser(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

func (s *StatisticsController) getDBSecondHandInfo(sel []string, where models.SecondHandInfo, option map[string]interface{}) result {
	//获取二手用户信息
	//result SecondHandInfo
	var resultStruct []models.SecondHandInfo

	result, ok := new(models.SecondHandInfo).GetSecondHandUser(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

func (s *StatisticsController) getDBOfflineCommissionPublishUserInfo(sel []string, where models.OfflineCommissionVerify, option map[string]interface{}) result {
	//获取待寄待取发布者用户信息
	//result OfflineCommissionVerify
	var resultStruct []models.OfflineCommissionVerify

	result, ok := new(models.OfflineCommissionVerify).GetOfflineCommission(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

func (s *StatisticsController) getDBOfflineCommissionPayUserInfo(sel []string, where models.OfflineCommissionOrder, option map[string]interface{}) result {
	//获取待寄待取付费者用户信息
	//result OfflineCommissionOrder
	var resultStruct []models.OfflineCommissionOrder

	result, ok := new(models.OfflineCommissionOrder).GetOfflineCommission(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

func (s *StatisticsController) getDBTelecomCardUsingUserInfo(sel []string, where models.TelecomUserCard, option map[string]interface{}) result {
	//获取大K卡用户信息
	//result TelecomUserCard
	var resultStruct []models.TelecomUserCard

	result, ok := new(models.TelecomUserCard).GetTelecomUser(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

func (s *StatisticsController) getDBKPlusUserInfo(sel []string, where models.KPlusOrder, option map[string]interface{}) result {
	//获取k+plus用户信息
	//result plus
	var resultStruct []models.KPlusOrder

	result, ok := new(models.KPlusOrder).GetKPlusUser(&resultStruct, &where, sel, &option)
	if !ok {
		return s.error(result)
	}
	return s.success(resultStruct)
}

//计算增长率
func (s *StatisticsController) getChance(current int, last int) float64 {
	//增长数据目前是计算 30 天前的数据和当前数据对比，现在需要修改为计算上个月份 1 日和当前月份 1 日之间的增长率。
	differ := current - last
	//获取增长率
	percentage := math.Round(float64(differ)/float64(last)*10000) / 100
	//防止无效结果值报错
	// NaN 代表 不是一个数 Not a number
	isNaN := math.IsNaN(percentage)
	//Inf 代表 阶码溢出，前面的加减符号代表高地位溢出，说白了就是小数点位后面无限大，再别的地方使用不能很好的序列化
	isInf := math.IsInf(percentage, 0)
	if isNaN || isInf {
		percentage = 0
	}
	return percentage
}
