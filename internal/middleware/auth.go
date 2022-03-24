package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//response := app.NewResponse(ctx)
		//authorizationHeader := ctx.GetHeader(token.AuthorizationKey)
		//if len(authorizationHeader) == 0 {
		//	response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
		//	ctx.Abort()
		//	return
		//}
		//fields := strings.SplitN(authorizationHeader, " ", 2)
		//if len(fields) != 2 || strings.ToLower(fields[0]) != global.AllSetting.Token.AuthorizationType {
		//	response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
		//	ctx.Abort()
		//	return
		//}
		//accessToken := fields[1]
		//payload, err := global.Maker.VerifyToken(accessToken)
		//if err != nil {
		//	response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr.WithDetails(err.Error()))
		//	ctx.Abort()
		//	return
		//}
		//ctx.Set(token.AuthorizationKey, payload)
		//ctx.Next()
	}
}

func ManagerAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func NoLogin() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
