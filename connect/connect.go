package connect

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"offergo/connect/internal"
	"offergo/log"
)

func InitHkokDb() func() {
	err := internal.HkokDbConnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		internal.HkokDbExit()
	}
}

func InitTelecomDb() func() {
	err := internal.TelecomDbConnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		internal.TelecomDbExit()
	}
}

func InitHkokJobDb() func() {
	err := internal.HkokJobDbConnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		internal.HkokJobDbExit()
	}
}

//redis连接实例
func InitRedis() func() {
	err := internal.RedisConnect()
	if err != nil {
		log.LogInfo.Info(err.Error())
		panic(err.Error())
	}
	return func() {
		internal.CloseRedis()
	}
}

//香不香港DB连接实例
func GetHkokDb() *gorm.DB {
	return internal.GetHkokDb()
}

//大K卡DB连接实例
func GetTelecomDb() *gorm.DB {
	return internal.GetTelecomDb()
}

//redis连接实例
func GetRedisInstance() *redis.Conn {
	return internal.GetRedisInstance()
}

//香不香港工作板块DB连接实例
func GetHkokJobDb() *gorm.DB {
	return internal.GetHkokJobDb()
}
