package controller

import (
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/gin-gonic/gin"
)

// DeleteUserStar
// @Summary 取消用户对于某个帖子的点赞
// @Description 通过post_id取消用户对于某个帖子的点赞
// @Tags user_star
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {string} string  ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /star/{post_id} [delete]
func DeleteUserStar(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// UserStarPost
// @Summary 用户对某个帖子点赞
// @Description 设置用户对于某个帖子的点赞
// @Tags user_star
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {string} string  ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /star/{post_id} [put]
func UserStarPost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
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
	response.ToResponse(true)
}
