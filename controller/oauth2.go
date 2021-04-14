// @Title        oauth2
// @Description  oauth2 登录接口
// @Author       Eachin
// @Date         2021/4/5 12:24 上午

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oauth/lib/response"
	"oauth/logic"
	"oauth/model/network"
)

// Oauth2ByWechatLoginRequiredSubscribe 微信登录，必须关注公众号
// @Summary 微信登录接口
// @Tags Oauth2
// @Accept application/json
// @Produce application/json
// @Param object query network.Oauth2ByWechatLoginGetParam true "查询参数"
// @Success 302 {string} string
// @Failure 400,422,500 {object} network.HTTPErrResponse
// @Router /oauth2/wechat/required/subscribe [get]
func Oauth2ByWechatLoginRequiredSubscribe(ctx *gin.Context) {
	var param network.Oauth2ByWechatLoginGetParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.ErrorInvalidParam(ctx, err)
		return
	}

	redirect, err := logic.Oauth2ByWechatLoginRequiredSubscribe(ctx, &param)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	ctx.Redirect(http.StatusFound, redirect)
}
