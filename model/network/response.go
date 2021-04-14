// @Title        response
// @Description
// @Author       Eachin
// @Date         2021/4/5 3:32 下午

package network

import (
	"oauth/lib/response"
)

type LoginToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// HTTPErrResponse 错误响应
type HTTPErrResponse struct {
	response.BaseModel
}

// LoginResponse 登录响应
type LoginResponse struct {
	response.BaseModel
	Detail *LoginToken `json:"detail"`
}
