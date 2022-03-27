package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/0RAJA/Road/internal/pkg/conversion"
	"github.com/gin-gonic/gin"
)

// UserStarPost
// @Summary 用户对某个帖子点赞或取消点赞
// @Description 用户对某个帖子点赞或取消点赞
// @Tags user_star
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param state body bool true "点赞状态 Enums(true,false)"
// @Success 200 {string} string  ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /star [put]
func UserStarPost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.UserStarPostParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	var err *errcode.Error
	if params.State {
		err = logic.UserStarPost(ctx, params.PostID)
	} else {
		err = logic.DeleteUserStar(ctx, params.PostID)
	}
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// GetUserStar
// @Summary 获取用户对于某个帖子的点赞情况
// @Description 通过post_id获取用户对于某个帖子的点赞情况
// @Tags user_star
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {bool} bool  "true表示已经点过赞，false表示没有"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /star/{post_id} [get]
func GetUserStar(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postID := conversion.AtoInt64Must(app.GetPath(ctx, "post_id"))
	if postID <= 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	result, err := logic.GetUserStar(ctx, postID)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(result)
}
