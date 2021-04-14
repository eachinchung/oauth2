// @Title        oauth2
// @Description	 oauth2 逻辑处理
// @Author       Eachin
// @Date         2021/4/5 6:00 下午

package logic

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"oauth/constant/exception"
	"oauth/dao"
	"oauth/lib/oss"
	"oauth/lib/requests"
	"oauth/lib/sonyFlake"
	"oauth/logger"
	"oauth/model/network"
	"oauth/setting"
	"strconv"
	"strings"
	"time"
)

func genOauth2Sign(userID uint64) string {
	var build strings.Builder
	expiresAt := strconv.FormatInt(time.Now().Add(time.Minute).Unix(), 10)
	uID := strconv.FormatUint(userID, 10)
	hash := md5.Sum([]byte(expiresAt + uID + setting.Conf.Secret.Oauth2))
	build.WriteString(expiresAt)
	build.WriteString("-")
	build.WriteString(uID)
	build.WriteString("-")
	build.WriteString(fmt.Sprintf("%x", hash))
	return build.String()
}

// Oauth2ByWechatLoginRequiredSubscribe 已关注公众号的微信用户登录接口
func Oauth2ByWechatLoginRequiredSubscribe(
	ctx *gin.Context,
	param *network.Oauth2ByWechatLoginGetParam,
) (string, error) {
	// 登录的重定向地址
	var redirect string
	var oauth2System map[string]string
	config, err := dao.GetNormalConfigByKeyAndVersion(ctx, "oauth2_system", 1)
	if err != nil {
		return redirect, err
	}
	_ = json.Unmarshal(config, &oauth2System)
	if redirect, ok := oauth2System[param.Service]; !ok {
		return redirect, exception.ErrorServerBusy
	}
	redirect += param.Redirect + "?sign="

	// 获取微信 openid
	oauth2AccessToken, err := GetWechatOauth2AccessToken(ctx, param.Code)
	if err != nil {
		return redirect, err
	}

	// 查询用户
	wechatUsers, err := dao.GetWechatUserByOpenID(ctx, oauth2AccessToken.Openid)
	if err == nil {
		sign := genOauth2Sign(wechatUsers.UserID)
		return redirect + sign, nil
	}

	// 向微信请求用户信息
	userInfo, err := GetWechatUserInfoFerCgi(ctx, oauth2AccessToken.Openid)
	if err != nil {
		return redirect, err
	}
	if userInfo.Subscribe == 0 {
		return redirect, exception.ErrorWechatNotSubscribe
	}

	userID, err := sonyFlake.GenID(ctx)
	if err != nil {
		return redirect, err
	}

	// 存储用户头像
	if userInfo.HeadImgURL != "" {
		avatar := strings.Replace(userInfo.HeadImgURL, "/132", "/0", -1)
		res, err := requests.Get(ctx, avatar, map[string]interface{}{})
		if err != nil {
			return redirect, err
		}
		body, err := res.Body()
		if err != nil {
			log := logger.NewSentry(ctx)
			log.Error("获取微信头像错误", err)
			return redirect, exception.ErrorServerBusy
		}
		hash := md5.Sum(body)
		fileName := fmt.Sprintf("%x", hash)

		ossKey, err := oss.ByteUploader(ctx, "ek-studio", "avatar/", fileName, body)
		if err != nil {
			return redirect, err
		}

		userInfo.HeadImgURL = ossKey
	}

	sign := genOauth2Sign(userID)
	err = dao.CreateUserByWechat(ctx, userID, userInfo.Nickname, userInfo.HeadImgURL, userInfo.Openid, userInfo.Unionid)
	return redirect + sign, err
}
