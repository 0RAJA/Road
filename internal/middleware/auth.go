package middleware

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	AuthorizationKey = "payload"
	RootKey          = "Root"
)

func AuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		authorizationHeader := ctx.GetHeader(AuthorizationKey)
		if len(authorizationHeader) == 0 {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
			ctx.Abort()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.AllSetting.Token.AuthorizationType {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
			ctx.Abort()
			return
		}
		accessToken := fields[1]
		payload, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr.WithDetails(err.Error()))
			ctx.Abort()
			return
		}
		ctx.Set(AuthorizationKey, payload)
		//TODO:检查是不是管理员
		panic(nil)
		ctx.Next()
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
