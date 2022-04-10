// +build doc

package routing

import (
	_ "github.com/0RAJA/Road/docs"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}

//go build -tags "doc" -o main main.go 即可编译
