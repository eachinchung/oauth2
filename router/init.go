// @Title        init
// @Description  初始化 Gin 路由相关服务
// @Author       Eachin
// @Date         2021/4/2 10:31 下午

package router

import (
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"oauth/controller"
	_ "oauth/docs"
	"oauth/logger"
	"oauth/setting"
)

// SetupRouter 安装路由
func SetupRouter() *gin.Engine {
	if setting.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery())
	r.Use(sentryGin.New(sentryGin.Options{Repanic: true}))
	r.GET("/docs/*any", ginSwagger.DisablingWrapHandler(
		swaggerFiles.Handler, "ENV"),
	)

	r.GET("/oauth2/wechat/required/subscribe", controller.Oauth2ByWechatLoginRequiredSubscribe)
	r.POST("/login/oauth2", controller.GenLoginTokenByOauth2)

	r.NoRoute(controller.NotFound)
	return r
}
