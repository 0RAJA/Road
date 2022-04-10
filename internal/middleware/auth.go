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

func Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		if _, ok := ctx.Get(AuthorizationKey); !ok {
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExistErr)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func Auth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		if len(authorizationHeader) == 0 {
			ctx.Next()
			return
		}
		fields := strings.SplitN(authorizationHeader, " ", 2)
		if len(fields) != 2 || strings.ToLower(fields[0]) != global.AllSetting.Token.AuthorizationType {
			ctx.Next()
			return
		}
		accessToken := fields[1]
		payload, err := global.Maker.VerifyToken(accessToken)
		if err != nil {
			ctx.Next()
			return
		}
		ctx.Set(AuthorizationKey, payload)
		if err = redis.Query.AddVisitedUserName(ctx, payload.UserName); err != nil {
			global.Logger.Error(err.Error())
		}
		if _, ok := ctx.Get(RootKey); !ok {
			_, err := mysql.Query.GetManagerByUsername(ctx, payload.UserName)
			ctx.Set(RootKey, err == nil)
		}
		ctx.Next()
	}
}

func ManagerAuth() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := app.NewResponse(ctx)
		if _, ok := ctx.Get(RootKey); !ok {
			response.ToErrorResponse(errcode.InsufficientPermissionsErr)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
