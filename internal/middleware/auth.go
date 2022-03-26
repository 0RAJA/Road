package middleware

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	"github.com/0RAJA/Road/internal/dao/redis"
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
		if err = redis.Query.AddVisitedUserName(ctx, payload.UserName); err != nil {
			global.Logger.Error(err.Error())
		}
		ctx.Set(AuthorizationKey, payload)
		if _, err := mysql.Query.GetManagerByUsername(ctx, payload.UserName); err == nil {
			ctx.Set(RootKey, true)
		}
		ctx.Next()
	}
}

func ManagerAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		_, ok := ctx.Get(RootKey)
		if !ok {
			response.ToErrorResponse(errcode.InsufficientPermissionsErr)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func NoLogin() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		authorizationHeader := ctx.GetHeader(AuthorizationKey)
		if len(authorizationHeader) == 0 {
			ctx.Next()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.AllSetting.Token.AuthorizationType {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
			ctx.Abort()
			return
		}
		accessToken := fields[1]
		_, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr.WithDetails(err.Error()))
			ctx.Abort()
			return
		}
		response.ToErrorResponse(errcode.ErrLoggedIn)
		ctx.Abort()
		return
	}
}
