// @Title        token
// @Description  登录token相关
// @Author       Eachin
// @Date         2021/4/13 12:36 下午

package controller

import (
	"github.com/gin-gonic/gin"
	"oauth/lib/response"
	"oauth/model/network"
)

// GenLoginTokenByOauth2 Oauth2的登录接口
// @Summary 通过签名获取token
// @Tags Login
// @Accept application/json
// @Produce application/json
// @Param object body network.GenLoginTokenByOauth2PostParam true "请求参数"
// @Success 200 {object} network.LoginResponse
// @Failure 400,422,500 {object} network.HTTPErrResponse
// @Router /login/oauth2 [post]
func GenLoginTokenByOauth2(ctx *gin.Context) {
	var param network.GenLoginTokenByOauth2PostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ErrorInvalidParam(ctx, err)
		return
	}
}
