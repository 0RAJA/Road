package controller

import (
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

}

// RefreshToken
// @Summary 刷新token
// @Description 将过期的token和未过期的retoken换取新的token
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} logic.RefreshTokenReply "返回用户信息和新的token"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /token/refresh [get]
func RefreshToken(ctx *gin.Context) {

}
