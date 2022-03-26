package middleware

import (
	"fmt"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/email"
	"github.com/0RAJA/Road/internal/pkg/times"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

//异常捕获处理

//我们需要针对我们的公司内部情况或生态圈定制 Recovery 中间件，确保异常在被正常捕抓之余，要及时的被识别和处理
//自定义 Recovery

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.AllSetting.Email.Host,
		Port:     global.AllSetting.Email.Port,
		IsSSL:    global.AllSetting.Email.IsSSL,
		UserName: global.AllSetting.Email.UserName,
		Password: global.AllSetting.Email.Password,
		From:     global.AllSetting.Email.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				err1 := defailtMailer.SendMail( //短信通知
					global.AllSetting.Email.To,
					fmt.Sprintf("异常抛出，发生时间: %v", time.Now().Format(times.LayoutDateTime)),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err1 != nil {
					global.Logger.Error(fmt.Sprintf("mail.SendMail Error: %v", err1.Error()))
				}

				// Check for a broken connection
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					global.Logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					global.Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
