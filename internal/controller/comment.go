package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// AddComment
// @Summary 创建评论
// @Description 创建一个对于帖子的评论或对评论的回复
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param username body string true "用户名"
// @Param content body string true "评论内容 1<=len<=100"
// @Param to_comment_id body int64 true "回复的评论ID"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /comment [post]
func AddComment(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// DeleteComment
// @Summary 删除评论
// @Description 删除一个评论
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param comment_id path int64 true "需要删除的评论ID"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /comment/{comment_id} [delete]
func DeleteComment(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ModifyComment
// @Summary 修改评论
// @Description 修改一个评论的内容
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param comment_id body int64 true "需要修改的评论ID"
// @Param content body string true "评论修改后的内容 1<=len<=100"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /comment [put]
func ModifyComment(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ListComments
// @Summary 显示评论
// @Description 根据post_id和偏移量显示一个帖子的部分评论的内容
// @Tags 评论
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param PostID query int64 true "需要显示评论的帖子ID"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {array} logic.ListCommentByPostIDReply "返回评论的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /comment [get]
func ListComments(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	comments := make([]logic.Comment, pageSize)
	for i := range comments {
		comments[i] = logic.Comment{
			ID:            utils.RandomInt(1, 100),
			PostID:        utils.RandomInt(1, 100),
			Username:      utils.RandomOwner(),
			Content:       utils.RandomString(100),
			ToCommentID:   utils.RandomInt(1, 100),
			CreateTime:    time.Now(),
			ModifyTime:    time.Now(),
			DepositoryUrl: utils.RandomString(10),
		}
	}
	response.ToResponseList(comments, len(comments))
}

/*
评论:
    增加
    删除
        通过ID删除
    修改
        修改内容
    查询
        通过post_id查
        通过id查 //测试
*/
