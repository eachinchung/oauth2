// @Title        not_found
// @Description  找不到资源 404
// @Author       Eachin
// @Date         2021/4/2 10:33 下午

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth/constant"
	"time"
)

// NotFound 404错误
func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"err_code": constant.ErrCodeNotFound,
		"message":  constant.ErrCodeNotFound.Msg(),
		"detail":   gin.H{},
		"now_ts":   time.Now().Unix(),
	})
}
