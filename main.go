package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"net/http"
	"offergo/connect"
	"offergo/log"
	_ "offergo/routers"
	_ "net/http/pprof"
	"offergo/ws"
	"time"
)

func init() {
	//开启协程即时监听端口
	go func() {
		fmt.Println("pprof start...")
		fmt.Println(http.ListenAndServe(":9876", nil))
	}()
}

func init() {
	//开启websocket协程
	go func() {
		http.HandleFunc("/ws/server", ws.WebsocketStart)
		err := http.ListenAndServe(":5678", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()
}
func init() {
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
}
func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	//hkok数据库连接
	defer startDb()()
	//大k卡数据库连接
	defer startTelecomDb()()
	defer startRedis()()
	//是否开启打印sql
	//connect.Getdb().LogMode(true)
	//connect.Gettdb().LogMode(true)
	beego.Run()
}

func startDb() func() {
	err := connect.Dbconnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.Dbexit()
	}
}

func startTelecomDb() func() {
	err := connect.Tdbconnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.Tdbexit()
	}
}

func startRedis() func() {
	err := connect.InitRedis()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.CloseRedis()
	}
}
