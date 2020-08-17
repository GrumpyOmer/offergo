package internal

import (
	// "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	//导入mysql驱动，这是必须的
	_ "github.com/go-sql-driver/mysql"
)

//香不香港DB连接配置
var mysql_user = beego.AppConfig.String("mysql_user")
var mysql_password = beego.AppConfig.String("mysql_password")
var mysql_host = beego.AppConfig.String("mysql_host")
var mysql_dbname = beego.AppConfig.String("mysql_dbname")
//工作板块库名
var mysql_job_dbname = beego.AppConfig.String("mysql_job_dbname")

//大K卡DB连接配置
var mysql_telecom_user = beego.AppConfig.String("mysql_telecom_user")
var mysql_telecom_password = beego.AppConfig.String("mysql_telecom_password")
var mysql_telecom_host = beego.AppConfig.String("mysql_telecom_host")
var mysql_telecom_dbname = beego.AppConfig.String("mysql_telecom_dbname")

var hkokDB *gorm.DB
var telecomDB *gorm.DB
var hkokJobDB *gorm.DB

func HkokDbConnect() error {
	var errors error
	//连接香不香港DB实例
	hkokDB, errors = gorm.Open("mysql", mysql_user+":"+mysql_password+"@"+"("+mysql_host+")/"+mysql_dbname+"?charset=utf8&parseTime=True&loc=Local")
	return errors
}

func TelecomDbConnect() error {
	var errors error
	//连接大K卡DB实例
	telecomDB, errors = gorm.Open("mysql", mysql_telecom_user+":"+mysql_telecom_password+"@"+"("+mysql_telecom_host+")/"+mysql_telecom_dbname+"?charset=utf8&parseTime=True&loc=Local")
	return errors
}

func HkokJobDbConnect() error {
	var errors error
	//连接香不香港工作板块DB实例
	hkokJobDB, errors = gorm.Open("mysql", mysql_user+":"+mysql_password+"@"+"("+mysql_host+")/"+mysql_job_dbname+"?charset=utf8&parseTime=True&loc=Local")
	return errors
}

//香不香港DB连接实例
func GetHkokDb() *gorm.DB {
	return hkokDB
}

//大K卡DB连接实例
func GetTelecomDb() *gorm.DB {
	return telecomDB
}

//香不香港工作板块DB连接实例
func GetHkokJobDb() *gorm.DB {
	return hkokJobDB
}

//关闭香不香港DB连接
func HkokDbExit() {
	hkokDB.Close()
}

//关闭大K卡DB连接
func TelecomDbExit() {
	telecomDB.Close()
}

//关闭香不香港工作板块DB连接
func HkokJobDbExit() {
	hkokJobDB.Close()
}