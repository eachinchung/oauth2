// @Title        users
// @Description  用户表相关操作
// @Author       Eachin
// @Date         2021/4/2 10:33 下午

package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"oauth/constant/exception"
	"oauth/logger"
	"oauth/model/mysql"
)

func createUser(
	ctx *gin.Context,
	engine *gorm.DB,
	userID uint64,
	username string,
	avatar string,
) (err error) {
	user := mysql.Users{
		UserID:   userID,
		Username: username,
		Avatar:   avatar,
	}

	if err = engine.Create(&user).Error; err != nil {
		log := logger.NewSentry(ctx)
		log.Error("创建用户失败", err)
		err = exception.ErrorServerBusy
	}
	return
}

// CreateUser 添加一个用户
func CreateUser(
	ctx *gin.Context,
	userID uint64,
	username string,
	avatar string,
) error {
	return createUser(ctx, db, userID, username, avatar)
}

// CreateUserByWechat 通过微信添加一个新用户
func CreateUserByWechat(
	ctx *gin.Context,
	userID uint64,
	username string,
	avatar string,
	openID string,
	unionID string,
) error {
	err := db.Transaction(func(tx *gorm.DB) (err error) {
		if err = createUser(ctx, tx, userID, username, avatar); err != nil {
			return err
		}
		err = createWechatUser(ctx, tx, userID, openID, unionID)
		return err
	})
	if err != nil {
		log := logger.NewSentry(ctx)
		log.Error("通过微信创建用户失败", err)
		return exception.ErrorServerBusy
	}
	return nil
}
