// @Title        configs
// @Description
// @Author       Eachin
// @Date         2021/4/4 9:42 下午

package dao

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oauth/constant"
	"oauth/constant/exception"
	"oauth/lib/redis"
	"oauth/logger"
	"oauth/model/mysql"
	"time"
)

func getNormalConfigByKeyAndVersion(
	ctx *gin.Context,
	engine *gorm.DB,
	configKey string,
	version int8,
) ([]byte, error) {
	key := fmt.Sprintf("config:%s:version:%d", configKey, version)
	if config, err := redis.Get(ctx, key); err == nil {
		return []byte(config), nil
	}

	config := mysql.Configs{}

	if err := engine.Where(
		"config_key = ? AND version = ? AND status = ?",
		configKey, version, constant.StatusNormal,
	).First(&config).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrorConfigNotExist
		}
		log := logger.NewSentry(ctx)
		log.Error("查询configs错误", err)
		return nil, exception.ErrorServerBusy
	}

	_ = redis.Set(ctx, key, string(config.Config), time.Hour)
	return config.Config, nil
}

// GetNormalConfigByKeyAndVersion 获取正常状态的配置
func GetNormalConfigByKeyAndVersion(
	ctx *gin.Context,
	configKey string,
	version int8,
) ([]byte, error) {
	return getNormalConfigByKeyAndVersion(ctx, db, configKey, version)
}
