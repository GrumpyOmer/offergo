package controllers

import (
	"offergo/lib"
	"offergo/models"
	"strconv"
	"strings"
)

type TelecomController struct {
	baseController
}

//获取白卡列表
func (t *TelecomController) GetWriteCardList() {
	//获取请求参数
	startNumber := t.Input().Get("startNumber")
	endNumber := t.Input().Get("endNumber")
	isWhiteCard := t.Input().Get("isWhiteCard")
	whiteCardChannel := t.Input().Get("whiteCardChannel")
	isUse := t.Input().Get("isUse")
	page := t.Input().Get("page")
	pageNum := t.Input().Get("pageNum")
	//必传参数验证
	validation := map[string]interface{}{
		"page": page,
	}
	t.requestFilter(validation)
	sel := []string{
		"id",
		"iccid",
		"status",
		"white_card",
		"channel_id", //关联查询必须要的字段
	}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	where := models.TelecomCard{}
	if isWhiteCard != "" {
		wheres["white_card = ?"] = isWhiteCard
	}
	if isUse != "" {
		wheres["status = ?"] = isUse
	}
	if startNumber != "" {
		wheres["iccid >= ?"] = startNumber
	}
	if endNumber != "" {
		wheres["iccid <= ?"] = endNumber
	}
	if whiteCardChannel != "" {
		wheres["channel_id = ?"] = whiteCardChannel
	}
	option["wheres"] = wheres
	option["page"] = map[string]interface{}{
		"page":    page,
		"pageNum": pageNum,
	}
	option["pageInfo"] = nil
	result := t.getDBTelecomCardInfo(sel, where, &option)
	if result.Code == 400 {
		t.responseError(result.Msg)
	}
	pageInfo, _ := option["pageInfo"].(lib.PageStruct)
	res := lib.PageResult{
		Data: result.Data,
		Page: pageInfo,
	}
	t.responseSuccess(res)
}

//获取白卡渠道列表
func (t *TelecomController) GetWhiteCardChannel() {
	page := t.Input().Get("page")
	status := t.Input().Get("status")
	pageNum := t.Input().Get("pageNum")
	//必传参数验证
	validation := map[string]interface{}{
		"page": page,
	}
	t.requestFilter(validation)
	sel := []string{
		"*",
	}
	where := models.WhiteCardChannel{}
	option := make(map[string]interface{})
	wheres := make(map[string]interface{})
	if status != "" {
		wheres["status = ?"] = status
	}
	option["wheres"] = wheres
	option["page"] = map[string]interface{}{
		"page":    page,
		"pageNum": pageNum,
	}
	option["pageInfo"] = nil
	option["orderBy"] = "created_at DESC"
	result := t.getWhiteCardChannelInfo(sel, where, &option)
	if result.Code == 400 {
		t.responseError(result.Msg)
	}
	pageInfo, _ := option["pageInfo"].(lib.PageStruct)
	res := lib.PageResult{
		Data: result.Data,
		Page: pageInfo,
	}
	t.responseSuccess(res)
}

//新增白卡渠道
func (t *TelecomController) AddWhiteCardChannel() {
	name := t.Input().Get("name")
	balance := t.Input().Get("balance")
	//必传参数验证
	validation := map[string]interface{}{
		"name":    name,
		"balance": balance,
	}
	t.requestFilter(validation)
	balanceToInt, err := strconv.Atoi(balance)
	//不是整形字符串驳回请求
	if err != nil {
		t.responseError(err.Error())
	}
	balanceToInt32 := int32(balanceToInt)
	data := models.WhiteCardChannel{
		Name:    name,
		Balance: balanceToInt32,
	}
	result := t.addWhiteCardChannelInfo(&data)
	if result.Code == 400 {
		t.responseError(result.Msg)
	}
	t.responseSuccess(nil)
}

//修改白卡渠道
func (t *TelecomController) TagWriteCardChannel() {
	//获取请求参数
	id := t.Input().Get("id")
	status := t.Input().Get("status")
	//必传参数验证
	validation := map[string]interface{}{
		"id":     id,
		"status": status,
	}
	t.requestFilter(validation)
	where := make(map[string]interface{})
	update := make(map[string]interface{})
	where["id = ?"] = id
	update["status"] = status
	result := t.updateWhiteCardChannelInfo(where, &update)
	if result.Code == 400 {
		t.responseError(result.Msg)
	}
	t.responseSuccess(nil)
}

//标记白卡
func (t *TelecomController) TagWriteCard() {
	//获取请求参数
	id := t.Input().Get("id")
	isWhiteCard := t.Input().Get("isWhiteCard")
	whiteCardId := t.Input().Get("whiteCardId")
	//必传参数验证
	validation := map[string]interface{}{
		"id":          id,
		"isWhiteCard": isWhiteCard,
	}
	t.requestFilter(validation)
	//get id slice
	ids := strings.Split(id, ",")
	where := make(map[string]interface{})
	update := make(map[string]interface{})
	where["id In (?)"] = ids
	update["white_card"] = isWhiteCard
	if isWhiteCard == "0" {
		//取消标记 清除whiteCardId
		whiteCardId = "0"
	}
	update["channel_id"] = whiteCardId
	result := t.updateDBTelecomCard(where, &update)
	if result.Code == 400 {
		t.responseError(result.Msg)
	}
	t.responseSuccess(nil)
}

func (t *TelecomController) getDBTelecomCardInfo(sel []string, where models.TelecomCard, option *map[string]interface{}) result {
	//获取大K卡信息
	//result TelecomUserCard
	var resultStruct []models.TelecomCard
	result, ok := new(models.TelecomCard).GetTelecomCard(&resultStruct, &where, sel, option)
	if !ok {
		return t.error(result)
	}
	return t.success(resultStruct)
}

func (t *TelecomController) updateDBTelecomCard(where map[string]interface{}, update *map[string]interface{}) result {
	//修改卡状态
	result, ok := new(models.TelecomCard).UpdateTelecomCard(where, update)
	if ok != true {
		t.error(result)
	}
	return t.success(nil)
}

func (t *TelecomController) getWhiteCardChannelInfo(sel []string, where models.WhiteCardChannel, option *map[string]interface{}) result {
	//获取白卡渠道信息
	//result WhiteCardChannel
	var resultStruct []models.WhiteCardChannel
	result, ok := new(models.WhiteCardChannel).GetWhiteCardChannel(&resultStruct, &where, sel, option)
	if !ok {
		t.error(result)
	}
	return t.success(resultStruct)
}

func (t *TelecomController) addWhiteCardChannelInfo(data *models.WhiteCardChannel) result {
	//新增白卡渠道信息
	result, ok := new(models.WhiteCardChannel).CreateWhiteCardChannel(data)
	if !ok {
		t.error(result)
	}
	return t.success(nil)
}

func (t *TelecomController) updateWhiteCardChannelInfo(where map[string]interface{}, update *map[string]interface{}) result {
	//修改白卡渠道状态
	result, ok := new(models.WhiteCardChannel).UpdateWhiteCardChannel(where, update)
	if ok != true {
		t.error(result)
	}
	return t.success(nil)
}
