package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"oauth/dao"
	"oauth/lib/redis"
	"oauth/lib/sonyFlake"
	"oauth/lib/validator"
	"oauth/logger"
	"oauth/router"
	"oauth/setting"
)

// @title Oauth2
// @description Oauth2后端API接口文档
// @version 0.0.1
func main() {
	if err := setting.Init(); err != nil {
		fmt.Printf("Config initialization failed:%v\n", err)
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         setting.Conf.Sentry.Dsn,
		Release:     fmt.Sprintf("%s@%s", setting.Conf.Name, setting.Conf.Version),
		Environment: setting.Conf.Mode,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
		return
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("Logger initialization failed:%v\n", err)
		return
	}

	if err := sonyFlake.Init("2021-04-01", setting.Conf.MachineID); err != nil {
		fmt.Printf("Sonyflake initialization failed:%v\n", err)
		return
	}

	if err := dao.SetupMySQL(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("Mysql initialization failed:%v\n", err)
		return
	}

	if err := redis.SetupRedis(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("Redis initialization failed:%v\n", err)
		return
	}
	defer redis.CloseRedis()

	if err := validator.Init(); err != nil {
		fmt.Printf("Validator initialization failed, err:%v\n", err)
		return
	}

	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("Server initialization failed:%v\n", err)
		return
	}
}
