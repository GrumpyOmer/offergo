package routers

import (
	"github.com/astaxie/beego"
	"offergo/controllers"
)

func init() {
	//hkokadmin
	hkokadmin := beego.NewNamespace("/hkokadmin",
		//获取控制台用户数据
		beego.NSRouter("/statistics/getUserStatistical", &controllers.StatisticsController{}, "Get:GetUserStatistical"),
		//获取白卡列表页数据
		beego.NSRouter("/telecom/getWriteCardList", &controllers.TelecomController{}, "Get:GetWriteCardList"),
		//标记白卡
		beego.NSRouter("/telecom/tagWriteCard", &controllers.TelecomController{}, "Post:TagWriteCard"),
		//获取白卡渠道页
		beego.NSRouter("/telecom/getWhiteCardChannel", &controllers.TelecomController{}, "Get:GetWhiteCardChannel"),
		//新增白卡渠道页
		beego.NSRouter("/telecom/addWhiteCardChannel", &controllers.TelecomController{}, "Post:AddWhiteCardChannel"),
		//修改白卡渠道状态
		beego.NSRouter("/telecom/tagWhiteCardChannel", &controllers.TelecomController{}, "Post:TagWriteCardChannel"),
		//获取自取点列表
		beego.NSRouter("/invite/getInviteList", &controllers.InviteController{}, "Get:GetInviteList"),
		//API更新/获取神州自取点
		beego.NSRouter("/invite/getShenZhouInviteData", &controllers.InviteController{}, "Get:GetShenZhouInviteData"),
		//更新/获取神州自取点
		beego.NSRouter("/invite/updateShenZhouInviteData", &controllers.InviteController{}, "Post:UpdateShenZhouInviteData"),
	)
	//hkokapp

	//websocket
	websocket := beego.NewNamespace("/ws",
		//client
		beego.NSRouter("/client", &controllers.WebsocketController{}, "Get:Index"),
	)

	//document
	document := beego.NewNamespace("/document",
		//replaceDocument
		beego.NSRouter("/replaceDocument", &controllers.GoQueryController{}, "Post:ReplaceDocument"),
		// SearchDocument
		beego.NSRouter("/searchDocument", &controllers.GoQueryController{}, "Post:SearchDocument"),
		// ReplaceSearchDocument
		beego.NSRouter("/replaceSearchDocument", &controllers.GoQueryController{}, "Post:ReplaceSearchDocument"),
	)
	//registerRouter
	beego.AddNamespace(hkokadmin)
	beego.AddNamespace(websocket)
	beego.AddNamespace(document)

}
