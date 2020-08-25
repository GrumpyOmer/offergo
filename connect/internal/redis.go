package internal

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

var redisInstance *redis.Conn
var redis_host = beego.AppConfig.String("redis_host")
var redis_pass = beego.AppConfig.String("redis_pass")

func RedisConnect() error {
	var errors error
	redisInstance, errors = redis_connect()
	return errors
}

//关闭redis实例
func CloseRedis() {
	(*redisInstance).Close()
}

func redis_connect() (*redis.Conn, error) {
	option := redis.DialPassword(redis_pass)
	c, err := redis.Dial("tcp", redis_host, option)
	if err != nil {
		fmt.Println("connect to redis error:" + err.Error())
	}

	return &c, err
}

func GetRedisInstance() *redis.Conn {
	return redisInstance
}
