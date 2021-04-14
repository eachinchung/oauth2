// @Title        response
// @Description  一些关于响应体的封装
// @Author       Eachin
// @Date         2021/4/2 10:16 下午

package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"oauth/constant"
	"oauth/constant/exception"
	vt "oauth/lib/validator"
	"oauth/logger"
	"time"
)

type jsonResponse interface {
	SetErrCode(errCode constant.ErrCode)
}

// BaseModel 响应体基类
type BaseModel struct {
	ErrCode   constant.ErrCode  `json:"err_code"`
	Message   string            `json:"message"`
	Detail    map[string]string `json:"detail"`
	Timestamp int64             `json:"now_ts"`
}

// SetErrCode 设置错误码
func (res *BaseModel) SetErrCode(errCode constant.ErrCode) {
	res.ErrCode = errCode
	res.Message = errCode.Msg()
	res.Timestamp = time.Now().Unix()
}

// Success 成功响应
func Success(ctx *gin.Context, res jsonResponse) {
	res.SetErrCode(constant.ErrCodeSuccess)
	ctx.JSON(http.StatusOK, &res)
}

// Error 失败响应
func Error(ctx *gin.Context, err error) {
	res := BaseModel{}
	res.Detail = map[string]string{}

	HTTPErr, ok := err.(*exception.HTTPError)
	if !ok {
		log := logger.NewSentry(ctx)
		log.Error("捕获到未处理的错误", err)
		res.SetErrCode(constant.ErrCodeServerBusy)
		ctx.JSON(http.StatusBadRequest, &res)
		return
	}

	res.SetErrCode(HTTPErr.ErrCode)
	ctx.JSON(http.StatusBadRequest, &res)
}

// ErrorInvalidParam 失败响应，参数不合法
func ErrorInvalidParam(ctx *gin.Context, err error) {
	res := BaseModel{}
	res.ErrCode = constant.ErrCodeInvalidParam
	res.Message = constant.ErrCodeInvalidParam.Msg()
	res.Timestamp = time.Now().Unix()

	errs, ok := err.(validator.ValidationErrors)
	if ok {
		res.Detail = vt.RemoveTopStruct(errs.Translate(vt.Trans))
	} else {
		res.Detail = map[string]string{}
	}

	ctx.JSON(http.StatusUnprocessableEntity, &res)
}
