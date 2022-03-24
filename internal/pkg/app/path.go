package app

import "github.com/gin-gonic/gin"

func GetPath(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}
