// @Title        wechat_users
// @Description  微信用户有关的数据库操作
// @Author       Eachin
// @Date         2021/4/5 4:05 下午

package dao

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"oauth/constant/exception"
	"oauth/lib/redis"
	"oauth/logger"
	"oauth/model/mysql"
	"time"
)

func createWechatUser(
	ctx *gin.Context,
	engine *gorm.DB,
	userID uint64,
	openID string,
	unionID string,
) (err error) {
	wechatUsers := mysql.WechatUsers{
		UserID:  userID,
		OpenID:  openID,
		UnionID: unionID,
	}
	if err = engine.Create(&wechatUsers).Error; err != nil {
		log := logger.NewSentry(ctx)
		log.Error("创建用户失败", err)
		err = exception.ErrorServerBusy
	}
	return
}

func getWechatUserByOpenID(
	ctx *gin.Context,
	engine *gorm.DB,
	openID string,
) (*mysql.WechatUsers, error) {
	w := &mysql.WechatUsers{}
	key := "w:w:" + openID
	if cache, err := redis.Get(ctx, key); err == nil {
		_ = json.Unmarshal([]byte(cache), w)
		zap.L().Debug("从缓存获取微信用户", zap.Uint64("uid", w.UserID))
		return w, nil
	}

	if err := engine.Where("open_id = ?", openID).First(w).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.ErrorUserExist
		} else {
			log := logger.NewSentry(ctx)
			log.Error("查询微信用户错误", err)
			err = exception.ErrorServerBusy
		}
		return w, err
	}

	cache, _ := json.Marshal(w)
	_ = redis.Set(ctx, key, string(cache), 10*time.Minute)
	zap.L().Debug("从MySQL获取微信用户", zap.Uint64("uid", w.UserID))
	return w, nil
}

// GetWechatUserByOpenID 通过 openid 获取微信用户
func GetWechatUserByOpenID(ctx *gin.Context, openID string) (*mysql.WechatUsers, error) {
	return getWechatUserByOpenID(ctx, db, openID)
}
