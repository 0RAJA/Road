package logic

import (
	"github.com/0RAJA/Road/internal/middleware"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/conversion"
	"github.com/0RAJA/Road/internal/pkg/singleflight"
	"github.com/0RAJA/Road/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	postKey     = "post:"
	postInfoKey = "postInfo:"
)

var doOnce = singleflight.NewGroup()

func getPostInfoKey(postID int64) string {
	return postInfoKey + conversion.Int64toA(postID)
}

func getPostKey(postID int64) string {
	return postKey + conversion.Int64toA(postID)
}

func getUsername(ctx *gin.Context) (string, error) {
	payload, ok := ctx.Get(middleware.AuthorizationKey)
	if !ok {
		return "", errcode.UnauthorizedAuthNotExistErr
	}
	return payload.(*token.Payload).UserName, nil
}

func getRoot(ctx *gin.Context) bool {
	_, ok := ctx.Get(middleware.RootKey)
	return ok
}
