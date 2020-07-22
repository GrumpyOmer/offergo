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

		//是否开启打印sql
		//connect.Getdb().LogMode(true)
		//connect.Gettdb().LogMode(true)
	}
	beego.Run()
}
