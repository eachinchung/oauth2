// @Title        wechat_users
// @Description  微信用户有关的数据库操作
// @Author       Eachin
// @Date         2021/4/5 4:05 下午

package dao

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
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
) (wu *mysql.WechatUsers, err error) {
	key := "w:u:" + openID
	if wuc, err := redis.Get(ctx, key); err == nil {
		_ = json.Unmarshal([]byte(wuc), wu)
		return wu, nil
	}

	if err = engine.Where("open_id = ?", openID).First(&wu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.ErrorUserExist
		} else {
			log := logger.NewSentry(ctx)
			log.Error("查询微信用户错误", err)
			err = exception.ErrorServerBusy
		}
		return
	}

	wuc, _ := json.Marshal(wu)
	_ = redis.Set(ctx, key, string(wuc), 10*time.Minute)
	return
}

// CreateWechatUser 添加一个用户
func CreateWechatUser(
	ctx *gin.Context,
	userID uint64,
	openID string,
	unionID string,
) error {
	return createWechatUser(ctx, db, userID, openID, unionID)
}

// GetWechatUserByOpenID 通过 openid 获取微信用户
func GetWechatUserByOpenID(ctx *gin.Context, openID string) (*mysql.WechatUsers, error) {
	return getWechatUserByOpenID(ctx, db, openID)
}
