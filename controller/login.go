// @Title        token
// @Description  登录token相关
// @Author       Eachin
// @Date         2021/4/13 12:36 下午

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth/lib/response"
	"oauth/logic"
	"oauth/model/network"
)

// GenLoginTokenByOauth2 Oauth2的登录接口
// @Summary 通过签名获取token
// @Tags Login
// @Accept application/json
// @Produce application/json
// @Param object body network.GenLoginTokenByOauth2PostParam true "请求参数"
// @Success 200 {object} network.LoginResponse
// @Failure 400,403,422,500 {object} network.HTTPErrResponse
// @Router /login/oauth2 [post]
func GenLoginTokenByOauth2(ctx *gin.Context) {
	var param network.GenLoginTokenByOauth2PostParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ErrorInvalidParam(ctx, err)
		return
	}

	tokens, err := logic.GenLoginTokenBySign(ctx, param.Sign)
	if err != nil {
		response.ErrorSetHTTPStatus(ctx, http.StatusForbidden, err)
		return
	}

	response.Success(ctx, &network.LoginResponse{Detail: tokens})
}

// RefreshLoginToken 刷新Token
// @Summary 通过刷新token获取token
// @Tags Login
// @Accept application/json
// @Produce application/json
// @Param object body network.RefreshLoginTokenPutParam true "请求参数"
// @Success 200 {object} network.LoginResponse
// @Failure 400,403,422,500 {object} network.HTTPErrResponse
// @Router /login [put]
func RefreshLoginToken(ctx *gin.Context) {
	var param network.RefreshLoginTokenPutParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.ErrorInvalidParam(ctx, err)
		return
	}

	tokens, err := logic.RefreshLoginToken(ctx, param.RefreshToken)
	if err != nil {
		response.ErrorSetHTTPStatus(ctx, http.StatusForbidden, err)
		return
	}

	response.Success(ctx, &network.LoginResponse{Detail: tokens})
}
