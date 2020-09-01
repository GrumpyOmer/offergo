package main

import (
	"github.com/astaxie/beego"
	"offergo/connect"
	_ "offergo/init" //程序初始化 （路由/子协程/配置）
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	//初始化连接项 defer用来进程退出时执行关闭连接操作
	{
		//hkok数据库连接
		defer connect.InitHkokDb()()
		//大k卡数据库连接
		defer connect.InitTelecomDb()()
		//redis连接
		defer connect.InitRedis()()
		//hkok工作板块数据库连接
		defer connect.InitHkokJobDb()()
		//是否开启打印sql
		connect.GetHkokDb().LogMode(true)
		//connect.GetTelecomDb().LogMode(true)
		//connect.GetHkokJobDb().LogMode(true)
	}
	beego.Run()
}
