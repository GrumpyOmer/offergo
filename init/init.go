package init

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"net/http"
	"offergo/controllers"
	"offergo/log"
	"offergo/ws"
	"time"
)

//程序初始化单元

func init() {
	//开启pprof协程 port: 9876
	go func() {
		fmt.Println("pprof start...")
		fmt.Println(http.ListenAndServe(":9876", nil))
	}()

	//开启websocket协程 port: 5678
	go func() {
		http.HandleFunc("/ws/server", ws.WebsocketStart)
		err := http.ListenAndServe(":5678", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	//初始化日志配置
	// JSONFormatter格式
	log.LogInfo.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     false,                 //格式化
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	})

	// 输出文件设置，默认为os.stderr
	today := time.Now().Format("2006-01-02")
	fileName := "logs/" + today + ".log"

	//添加log打印行号
	log.LogInfo.SetReportCaller(true)

	// 设置日志等级 只输出不低于当前级别是日志数据
	log.LogInfo.SetLevel(logrus.TraceLevel)

	//添加钩子
	log.LogInfo.AddHook(log.NewLfsHook(fileName, 5, 1))

	//路由初始化
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
