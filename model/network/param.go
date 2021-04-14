// @Title        param
// @Description  接口参数校验的模型
// @Author       Eachin
// @Date         2021/4/5 12:28 上午

package network

type Oauth2ByWechatLoginGetParam struct {
	Code     string `json:"code" form:"code" binding:"required"`
	State    string `json:"state" form:"state"`
	Service  string `json:"service" form:"service" binding:"required,oneof=university"` // 重定向：university
	Redirect string `json:"redirect" form:"redirect" binding:"required,uri"`            // 重定向的uri
}

type GenLoginTokenByOauth2PostParam struct {
	Sign string `json:"sign" binding:"required"  example:"1902390173-1765458873155585-WRVFULNCSV-fffc12a5f1d57536c90ecbc8b7bc91ca"` // 登录签名
}

type RefreshLoginTokenPutParam struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新token
}
