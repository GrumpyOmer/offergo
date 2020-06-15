package controllers

import (
	"offergo/lib"
	"offergo/models"
)

//statistics.go's
type getUserStatisticalStruct struct {
	CurrentUser      int
	CurrentMonthUser int
	LastMonthUser    int
	Percentage       float64
	Text             string
}
type getUserStatisticalResponseData struct {
	WxUser                       getUserStatisticalStruct //微信用户信息
	AppUser                      getUserStatisticalStruct //总APP用户信息
	IosUser                      getUserStatisticalStruct //ios用户信息
	AndroidUser                  getUserStatisticalStruct //android用户信息
	JobBookUser                  getUserStatisticalStruct //工作板块用户信息
	ActiveBookUser               getUserStatisticalStruct //活动板块用户信息
	MbUser                       getUserStatisticalStruct //集运信息
	SecondHandUser               getUserStatisticalStruct //二手用户信息
	OfflineCommissionPublishUser getUserStatisticalStruct //待寄待取发布者用户信息
	OfflineCommissionPayUser     getUserStatisticalStruct //待寄待取付费者用户信息
	TelecomCardActivationedUser  getUserStatisticalStruct //大K卡已激活总人数信息
	TelecomUsingUser             getUserStatisticalStruct //正在使用的大k卡用户信息
	TelecomNewUser               getUserStatisticalStruct //新申请的大k卡用户信息
}

//page's struct
type PageResult struct {
	Data interface{}
	Page models.PageStruct
}

//shenZhouInvite's api struct
type shenZhouInviteApiResult struct {
	Code int
	Msg  string
	Data inviteStruct
}

type inviteStruct struct {
	HongKong  []lib.HongKongStruct  `json:"hong_kong"`
	TakePoint []lib.TakePointStruct `json:"take_point"`
}

//telecom.go's
