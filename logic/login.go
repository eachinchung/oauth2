// @Title        login
// @Description  登录相关逻辑
// @Author       Eachin
// @Date         2021/4/14 9:37 下午

package logic

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"oauth/constant/exception"
	"oauth/lib/jwt"
	"oauth/model/network"
	"oauth/setting"
	"strconv"
	"strings"
	"time"
)

// GenLoginTokenBySign 生成登录的token
func GenLoginTokenBySign(ctx *gin.Context, sign string) (*network.LoginToken, error) {
	var t network.LoginToken
	signArr := strings.Split(sign, "-")
	if len(signArr) != 4 {
		return nil, exception.ErrorInvalidSign
	}
	expiresAt, err := strconv.ParseInt(signArr[0], 10, 64)
	if err != nil {
		return &t, exception.ErrorInvalidSign
	}
	userID, err := strconv.ParseInt(signArr[1], 10, 64)
	if err != nil {
		return &t, exception.ErrorInvalidSign
	}
	if time.Now().Unix() > expiresAt {
		return &t, exception.ErrorInvalidSign
	}
	var build strings.Builder
	build.WriteString(signArr[0])
	build.WriteString(signArr[1])
	build.WriteString(signArr[2])
	build.WriteString(setting.Conf.Secret.Oauth2)
	hash := md5.Sum([]byte(build.String()))
	if fmt.Sprintf("%x", hash) != signArr[3] {
		return &t, exception.ErrorInvalidSign
	}

	claims := jwt.NewClaims(userID)
	if t.AccessToken, err = jwt.GenToken(ctx, claims, setting.Conf.Secret.AccessToken, time.Hour); err != nil {
		return nil, err
	}
	if t.RefreshToken, err = jwt.GenToken(ctx, claims, setting.Conf.Secret.RefreshToken, 15*24*time.Hour); err != nil {
		return nil, err
	}

	return &t, nil
}

// RefreshLoginToken 刷新Token
func RefreshLoginToken(ctx *gin.Context, refreshToken string) (*network.LoginToken, error) {
	var t network.LoginToken
	var claims jwt.BaseClaims
	err := jwt.ParseToken(refreshToken, &claims, setting.Conf.Secret.RefreshToken)
	if err != nil {
		return nil, err
	}

	if t.AccessToken, err = jwt.GenToken(ctx, &claims, setting.Conf.Secret.AccessToken, time.Hour); err != nil {
		return nil, err
	}
	if t.RefreshToken, err = jwt.GenToken(ctx, &claims, setting.Conf.Secret.RefreshToken, 15*24*time.Hour); err != nil {
		return nil, err
	}

	return &t, nil
}
