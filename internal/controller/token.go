package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/gin-gonic/gin"
)

// GetToken
// @Summary 获取token
// @Description 用户将被重定向到github授权，服务器会得到授权后返回用户信息和token,retoken
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} logic.GetTokenReply "返回用户信息和token以及retoken"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /token/get [get]
func GetToken(ctx *gin.Context) {
	logic.GetToken(ctx)
}

func TokenRedirect(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.TokenRedirectParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	if len(params.Code) == 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	reply, err := logic.TokenRedirect(ctx, params.Code)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(reply)
}

// RefreshToken
// @Summary 刷新token
// @Description 将过期的token和未过期的retoken换取新的token
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param token body string true "过期的token"
// @Param re_token body string true "未过期的刷新token"
// @Success 200 {object} logic.RefreshTokenReply "返回用户信息和新的token"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /token/refresh [put]
func RefreshToken(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.RefreshTokenReplyParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.RefreshToken(ctx, params)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(reply)
}
