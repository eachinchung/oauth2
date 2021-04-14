// @Title        redis
// @Description  安装 Redis
// @Author       Eachin
// @Date         2021/4/2 10:16 下午

package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"oauth/constant/exception"
	"oauth/logger"
	"oauth/setting"
	"time"
)

var client *redis.Client

// SetupRedis 初始化 Redis 连接
func SetupRedis(config *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})

	if _, err = client.Ping().Result(); err != nil {
		return err
	}
	return nil
}

func checkKeyIsNotExist(ctx *gin.Context, err error) bool {
	if err == redis.Nil {
		return true
	}

	log := logger.NewSentry(ctx)
	log.Error("RedisErr", err)
	return false
}

func Set(ctx *gin.Context, key string, value interface{}, expiration time.Duration) error {
	if err := client.Set(key, value, expiration).Err(); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("RedisErr", err)
		return exception.ErrorServerBusy
	} else {
		return nil
	}
}

func Get(ctx *gin.Context, key string) (string, error) {
	if val, err := client.Get(key).Result(); err != nil {
		checkKeyIsNotExist(ctx, err)
		return "", exception.ErrorServerBusy
	} else {
		return val, nil
	}
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() {
	_ = client.Close()
}
