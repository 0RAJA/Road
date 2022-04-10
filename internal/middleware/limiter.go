package middleware

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func Limiter(l limiter.Iface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok { //查询是否需要限流
			cnt := bucket.TakeAvailable(1)
			if cnt == 0 {
				global.Logger.Info(errcode.TooManyRequestsErr.Error() + c.Request.RequestURI)
				app.NewResponse(c).ToErrorResponse(errcode.TooManyRequestsErr)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
