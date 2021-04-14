// @Title        exception
// @Description  自定义错误类型
// @Author       Eachin
// @Date         2021/4/8 1:13 上午

package exception

import "oauth/constant"

type HTTPError struct {
	Err     string
	ErrCode constant.ErrCode
}

// NewHTTPError 新建一个 HTTPError
func NewHTTPError(errCode constant.ErrCode) error {
	return &HTTPError{errCode.Msg(), errCode}
}

func (e *HTTPError) Error() string {
	return e.Err
}
