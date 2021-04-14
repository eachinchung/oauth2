// @Title        init
// @Description
// @Author       Eachin
// @Date         2021/4/4 11:44 下午

package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
	"go.uber.org/zap"
	"net"
	"net/http"
	"oauth/constant/exception"
	"oauth/logger"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var client = HttpRequest.Transport(transport)

// Get 发送 get 请求
func Get(ctx *gin.Context, url string, param map[string]interface{}) (*HttpRequest.Response, error) {
	if res, err := client.Get(url, param); err != nil {
		log := logger.NewSentry(ctx)
		log.Error("HttpRequest error", err)
		return res, exception.ErrorServerBusy
	} else {
		zap.L().Info(
			"HttpRequest",
			zap.String("url", res.Url()),
			zap.String("method", "Get"),
			zap.Int("status", res.StatusCode()),
		)
		return res, nil
	}
}
