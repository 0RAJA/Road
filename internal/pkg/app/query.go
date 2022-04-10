package app

import (
	"github.com/0RAJA/Road/internal/pkg/times"
	"github.com/gin-gonic/gin"
	"time"
)

func GetPath(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}

func GetTime(ctx *gin.Context, key string) (time.Time, error) {
	return time.Parse(times.LayoutDateTime, ctx.Query(key))
}
