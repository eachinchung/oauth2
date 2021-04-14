// @Title        sentry
// @Description  整合 sentry 与 zap
// @Author       Eachin
// @Date         2021/4/2 10:29 下午

package logger

import (
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Sentry struct {
	Context *gin.Context
}

func NewSentry(ctx *gin.Context) *Sentry {
	return &Sentry{Context: ctx}
}

func (sentry *Sentry) Warn(msg string) {
	hub := sentryGin.GetHubFromContext(sentry.Context)
	hub.CaptureMessage(msg)
	logger.Warn(msg)
}

func (sentry *Sentry) Error(msg string, exception error) {
	hub := sentryGin.GetHubFromContext(sentry.Context)
	hub.CaptureException(exception)
	logger.Error(msg, zap.Error(exception))
}
