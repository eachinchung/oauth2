// @Title        errors
// @Description
// @Author       Eachin
// @Date         2021/4/8 11:07 下午

package exception

import (
	"oauth/constant"
)

var (
	ErrorServerBusy         = NewHTTPError(constant.ErrCodeServerBusy)
	ErrorUserExist          = NewHTTPError(constant.ErrCodeUserNotExist)
	ErrorConfigNotExist     = NewHTTPError(constant.ErrCodeConfigNotExist)
	ErrorWechatNotSubscribe = NewHTTPError(constant.ErrCodeWechatNotSubscribe)
	ErrorWechatUnauthorized = NewHTTPError(constant.ErrCodeWechatUnauthorized)
)
