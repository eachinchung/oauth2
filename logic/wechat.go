// @Title        wechat
// @Description  微信有关操作
// @Author       Eachin
// @Date         2021/4/4 7:46 下午

package logic

import (
	"github.com/gin-gonic/gin"
	"oauth/constant/exception"
	"oauth/lib/redis"
	"oauth/lib/requests"
	"oauth/lib/validator"
	"oauth/logger"
	"oauth/setting"
	"time"
)

type WechatUserInfo interface {
	GetUserInfo() *wechatUserInfoJson
}

type wechatAccessTokenJson struct {
	AccessToken string `json:"access_token" validate:"required"`
	ExpiresIn   int    `json:"expires_in" validate:"required"`
}

type wechatOauth2AccessTokenJson struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	ExpiresIn    int    `json:"expires_in" validate:"required"`
	Openid       string `json:"openid" validate:"required"`
	Scope        string `json:"scope" validate:"required"`
}

type wechatUserInfoJson struct {
	Openid     string   `json:"openid" validate:"required"`
	Nickname   string   `json:"nickname" validate:"required"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func (u *wechatUserInfoJson) GetUserInfo() *wechatUserInfoJson {
	return u
}

type wechatUserInfoJsonForSns struct {
	wechatUserInfoJson
}

type wechatUserInfoJsonForCgi struct {
	wechatUserInfoJson
	Subscribe      int    `json:"subscribe"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

// GetAccessToken 获取微信访问 token
func GetAccessToken(ctx *gin.Context) (string, error) {
	key := "wechat:access_token"
	if accessToken, err := redis.Get(ctx, key); err == nil {
		return accessToken, nil
	}

	res, err := requests.Get(ctx, "https://api.weixin.qq.com/cgi-bin/token", map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      setting.Conf.WechatConfig.AppID,
		"secret":     setting.Conf.WechatConfig.AppSecret,
	})

	if err != nil {
		return "", err
	}

	accessToken := &wechatAccessTokenJson{}

	if err = res.Json(accessToken); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("获取 accessToken 失败", err)
		return "", exception.ErrorServerBusy
	}

	if err = validator.Validate.Struct(accessToken); err != nil {
		return "", exception.ErrorServerBusy
	}

	_ = redis.Set(ctx, key, accessToken.AccessToken, time.Second*7000)
	return accessToken.AccessToken, err
}

// GetWechatOauth2AccessToken 获取 oauth2 的访问 token
func GetWechatOauth2AccessToken(ctx *gin.Context, code string) (*wechatOauth2AccessTokenJson, error) {
	oauth2AccessToken := &wechatOauth2AccessTokenJson{}

	res, err := requests.Get(ctx, "https://api.weixin.qq.com/sns/oauth2/access_token", map[string]interface{}{
		"grant_type": "authorization_code",
		"appid":      setting.Conf.WechatConfig.AppID,
		"secret":     setting.Conf.WechatConfig.AppSecret,
		"code":       code,
	})

	if err != nil {
		return oauth2AccessToken, err
	}

	if err = res.Json(oauth2AccessToken); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("获取 oauth2 的访问 token 失败", err)
		return oauth2AccessToken, exception.ErrorServerBusy
	}

	if err = validator.Validate.Struct(oauth2AccessToken); err != nil {
		return oauth2AccessToken, exception.ErrorServerBusy
	}

	return oauth2AccessToken, nil
}

// GetWechatUserInfoFerSns 获取微信用户信息
func GetWechatUserInfoFerSns(
	ctx *gin.Context,
	oauth2AccessToken *wechatOauth2AccessTokenJson,
) (*wechatUserInfoJsonForSns, error) {
	wechatUserInfo := &wechatUserInfoJsonForSns{}

	if oauth2AccessToken.Scope != "snsapi_userinfo" {
		return wechatUserInfo, exception.ErrorWechatUnauthorized
	}

	res, err := requests.Get(ctx, "https://api.weixin.qq.com/sns/userinfo", map[string]interface{}{
		"lang":         "zh_CN",
		"access_token": oauth2AccessToken.AccessToken,
		"openid":       oauth2AccessToken.Openid,
	})

	if err != nil {
		return wechatUserInfo, err
	}

	if err = res.Json(wechatUserInfo); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("获取用户信息失败", err)
		return wechatUserInfo, exception.ErrorServerBusy
	}

	if err = validator.Validate.Struct(wechatUserInfo); err != nil {
		return wechatUserInfo, exception.ErrorServerBusy
	}

	return wechatUserInfo, nil
}

// GetWechatUserInfoFerCgi 获取微信用户信息
func GetWechatUserInfoFerCgi(
	ctx *gin.Context,
	openID string,
) (*wechatUserInfoJsonForCgi, error) {
	wechatUserInfo := &wechatUserInfoJsonForCgi{}

	accessToken, err := GetAccessToken(ctx)
	if err != nil {
		return wechatUserInfo, err
	}

	res, err := requests.Get(ctx, "https://api.weixin.qq.com/cgi-bin/user/info", map[string]interface{}{
		"lang":         "zh_CN",
		"access_token": accessToken,
		"openid":       openID,
	})

	if err != nil {
		return wechatUserInfo, err
	}

	if err = res.Json(wechatUserInfo); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("获取用户信息失败", err)
		return wechatUserInfo, exception.ErrorServerBusy
	}

	if err = validator.Validate.Struct(wechatUserInfo); err != nil {
		return wechatUserInfo, exception.ErrorServerBusy
	}

	return wechatUserInfo, nil
}
