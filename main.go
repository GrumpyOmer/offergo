package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"offergo/connect"
	_ "offergo/routers"
	"offergo/ws"
	"time"
)

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
func main() {
	today := time.Now().Format("2006-01-02")
	fileName := "logs/" + today + ".log"
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+fileName+`", "level":7, "daily":true, "maxdays":3}`)
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	//hkok数据库连接
	defer startdb()()
	//大k卡数据库连接
	defer starttdb()()
	defer startredis()()
	//是否开启打印sql
	connect.Getdb().LogMode(true)
	//connect.Gettdb().LogMode(true)
	beego.Run()
}

func startdb() func() {
	err := connect.Dbconnect()
	if err != nil {
		logs.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.Dbexit()
	}
}

func starttdb() func() {
	err := connect.Tdbconnect()
	if err != nil {
		logs.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.Tdbexit()
	}
}

func startredis() func() {
	err := connect.InitRedis()
	if err != nil {
		logs.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		connect.CloseRedis()
	}
}
