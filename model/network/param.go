// @Title        param
// @Description
// @Author       Eachin
// @Date         2021/4/5 12:28 上午

package network

type Oauth2ByWechatLoginGetParam struct {
	Code     string `json:"code" form:"code" binding:"required"`
	State    string `json:"state" form:"state"`
	Service  string `json:"service" form:"service" binding:"required,oneof=university" example:"university"` // 重定向：university
	Redirect string `json:"redirect" form:"redirect" binding:"required,uri"`                                 // 重定向的uri
}

type GenLoginTokenByOauth2PostParam struct {
	Sign string `json:"sign" binding:"required"` // 登录签名
}
