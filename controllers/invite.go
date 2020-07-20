package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"offergo/lib"
	"offergo/models"
)

type InviteController struct {
	baseController
}

//获取自取点列表
func (i *InviteController) GetInviteList() {
	page := i.Input().Get("page")
	status, status_err := i.GetInt8("status")
	invite_area, invite_areas_err := i.GetInt8("invite_area")
	key_word := i.GetString("key_word")
	pageNum := i.Input().Get("pageNum")
	//必传参数验证
	validation := map[string]interface{}{
		"page": page,
	}
	i.requestFilter(validation)
	sel := []string{
		"*",
	}
	option := make(map[string]interface{})
	where := models.Invite{}
	option["page"] = map[string]interface{}{
		"page":    page,
		"pageNum": pageNum,
	}
	option["pageInfo"] = nil
	wheres := make(map[string]interface{})
	if key_word != "" {
		wheres["invite_id = '"+key_word+"' or invite_name = '"+key_word+"'"] = nil
	}
	if status_err == nil {
		wheres["status = ?"] = status
	}
	if invite_areas_err == nil {
		wheres["invite_area = ?"] = invite_area
	}
	option["wheres"] = wheres
	result := i.getInviteList(sel, where, &option)
	if result.Code == 400 {
		i.responseError(result.Msg)
	}
	pageInfo, _ := option["pageInfo"].(lib.PageStruct)
	res := lib.PageResult{
		Data: result.Data,
		Page: pageInfo,
	}
	i.responseSuccess(res)
}

//更新/获取神州集运数据
func (i *InviteController) UpdateShenZhouInviteData() {
	invite_id, id_err := i.GetInt32("invite_id")
	status := i.Input().Get("status")
	user_describle := i.Input().Get("user_describle")
	if id_err != nil {
		i.responseError(id_err.Error())
	}
	//必传参数验证
	validation := map[string]interface{}{
		"invite_id": invite_id,
	}
	i.requestFilter(validation)
	if status != "" {
		//修改显示状态 是否显示 0 否 1 是
		where := make(map[string]interface{})
		update := make(map[string]interface{})
		where["invite_id = ?"] = invite_id
		update["status"] = status
		result := i.updateInviteInfo(where, &update)
		if result.Code == 400 {
			i.responseError(result.Msg)
		}
	} else if user_describle != "" {
		//修改/新增自定义自取点描述
		//判断是否存在 存在就修改 不存在就新增
		where := make(map[string]interface{})
		update := make(map[string]interface{})
		where["invite_id = ?"] = invite_id
		data := models.InviteDescrible{}
		searchResult := i.getInviteDescribleInfo(where, &data)
		if searchResult.Code == 400 && searchResult.Msg == "无数据" {
			//不存在 直接新增
			data.InviteID = invite_id
			data.UserDescrible = user_describle
			result := i.addInviteDescribleInfo(&data)
			if result.Code == 400 {
				i.responseError(result.Msg)
			}
		} else {
			//存在 直接更新
			update["user_describle"] = user_describle
			result := i.updateInviteDescribleInfo(where, &update)
			if result.Code == 400 {
				i.responseError(result.Msg)
			}
		}
	} else {
		i.responseError("无任何更新")
	}
	i.responseSuccess("ok")

}

//从APi更新/获取神州集运数据
func (i *InviteController) GetShenZhouInviteData() {
	res, err := http.Get(beego.AppConfig.String("ShenZhouInviteApi"))
	if err != nil {
		i.responseError(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		i.responseError(res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		i.responseError(err.Error())
	}
	//解析结果
	var requestResult lib.ShenZhouInviteApiResult
	json.Unmarshal(body, &requestResult)
	err = i.decodeApiResult(&requestResult)
	if err != nil {
		i.responseError(err.Error())
	}
	//按需求过滤 method_id 为 114、1239 和 method_type = 1的数据
	i.filterResult(&requestResult)
	//先删除所有数据，再重新添加进去
	i.deleteInviteList()
	result := i.addInviteList(&requestResult.Data.TakePoint)
	if result.Code == 200 {
		i.responseSuccess("OK")
	}
	i.responseError(result.Msg)
}

//解析神州api响应结果
func (i *InviteController) decodeApiResult(apiResult *lib.ShenZhouInviteApiResult) error {
	if apiResult.Code != 200 {
		//神州请求出错
		return errors.New(apiResult.Msg)
	}
	return nil
}

//结果过滤
func (i *InviteController) filterResult(apiResult *lib.ShenZhouInviteApiResult) {
	index := 0
	for k, v := range apiResult.Data.TakePoint {
		if v.MethodId == 114 || v.MethodId == 1239 || v.MethodType == 1 {
			//直接切片内部交换 无需扩容 copy
			temp := apiResult.Data.TakePoint[k]
			apiResult.Data.TakePoint[k] = apiResult.Data.TakePoint[index]
			apiResult.Data.TakePoint[index] = temp
			index++
		}
	}
	apiResult.Data.TakePoint = apiResult.Data.TakePoint[index:]
}

func (i *InviteController) getInviteList(sel []string, where models.Invite, option *map[string]interface{}) result {
	//获取自取点列表信息
	//result Invite
	var resultStruct []models.Invite
	result, ok := new(models.Invite).GetInviteList(&resultStruct, &where, sel, option)
	if !ok {
		return i.error(result)
	}
	return i.success(resultStruct)
}

func (i *InviteController) deleteInviteList() {
	new(models.Invite).DeleteInviteList()
}

//批量添加自取点列表
func (i *InviteController) addInviteList(invite *[]lib.TakePointStruct) result {
	sql := "INSERT INTO `invite` (`invite_id`, `invite_name`, `invite_address`, `invite_area`, `api_describle`, `invite_location`) VALUES "
	//循环切片，组合sql语句
	for k, v := range *invite {
		if len(*invite)-1 == k {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,'%s','%s','%d','%s','%d');", v.MethodId, v.MethodName, v.TakePointAddress, v.TakePointArea, v.MethodDescription, v.TakePointLocation)
		} else {
			sql += fmt.Sprintf("(%d,'%s','%s','%d','%s','%d'),", v.MethodId, v.MethodName, v.TakePointAddress, v.TakePointArea, v.MethodDescription, v.TakePointLocation)
		}
	}
	result, ok := new(models.Invite).InsertManyRecords(sql)
	if !ok {
		return i.error(result)
	}
	return i.success("")
}

//修改自取点信息
func (i *InviteController) updateInviteInfo(where map[string]interface{}, update *map[string]interface{}) result {
	//修改自取点信息
	result, ok := new(models.Invite).UpdateInviteInfo(where, update)
	if ok != true {
		return i.error(result)
	}
	return i.success(nil)
}

//修改自取点自定义信息
func (i *InviteController) updateInviteDescribleInfo(where map[string]interface{}, update *map[string]interface{}) result {
	//修改自取点自定义信息
	result, ok := new(models.InviteDescrible).UpdateInviteDescribleInfo(where, update)
	if ok != true {
		return i.error(result)
	}
	return i.success(nil)
}

//获取自取点自定义信息
func (i *InviteController) getInviteDescribleInfo(where map[string]interface{}, describle *models.InviteDescrible) result {
	//修改自取点自定义信息
	result, ok := new(models.InviteDescrible).GetInviteDescribleInfo(where, describle)
	if ok != true {
		return i.error(result)
	}
	return i.success(nil)
}

//新增单条自取点自定义信息
func (i *InviteController) addInviteDescribleInfo(describle *models.InviteDescrible) result {
	//修改自取点自定义信息
	result, ok := new(models.InviteDescrible).AddInviteDescribleInfo(describle)
	if ok != true {
		return i.error(result)
	}
	return i.success(nil)
}
