// @Title        err_code
// @Description  错误码
// @Author       Eachin
// @Date         2021/4/2 10:33 下午

package constant

type ErrCode int

const (
	ErrCodeSuccess ErrCode = 1000 + iota
	ErrCodeInvalidParam
	ErrCodeServerBusy
	ErrCodeNotFound
	ErrCodeInternalServerError

	ErrCodeNeedLogin
	ErrCodeInvalidSign
	ErrCodeInvalidToken

	ErrCodeUserExist
	ErrCodeUserNotExist
	ErrCodeInvalidPassword

	ErrCodeConfigNotExist
	ErrCodeConfigWechatUserExist
)

const (
	ErrCodeWechatUnauthorized ErrCode = 5001 + iota
	ErrCodeWechatNotSubscribe
)

var errCodeMsgMap = map[ErrCode]string{
	ErrCodeSuccess:             "success",
	ErrCodeInvalidParam:        "请求参数错误",
	ErrCodeServerBusy:          "服务器繁忙",
	ErrCodeNotFound:            "资源不存在",
	ErrCodeInternalServerError: "服务器内部错误",

	ErrCodeNeedLogin:    "需要登录",
	ErrCodeInvalidSign:  "无效的签名",
	ErrCodeInvalidToken: "无效的token",

	ErrCodeUserExist:       "用户已存在",
	ErrCodeUserNotExist:    "用户不存在",
	ErrCodeInvalidPassword: "用户名或密码错误",

	ErrCodeConfigNotExist:        "配置不存在",
	ErrCodeConfigWechatUserExist: "微信用户不存在",

	ErrCodeWechatUnauthorized: "用户未授权",
	ErrCodeWechatNotSubscribe: "用户未关注公众号",
}

func (c ErrCode) Msg() string {
	msg, ok := errCodeMsgMap[c]
	if !ok {
		msg = errCodeMsgMap[ErrCodeServerBusy]
	}
	return msg
}
