package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/gin-gonic/gin"
)

// AddPostTag
// @Summary 给一个帖子加上一个标签
// @Description 给一个帖子加上一个标签
// @Tags 文章和标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param tag_id body int64 true "标签ID"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /postTag/add [post]
func AddPostTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.PostTagParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := logic.AddPostTag(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// DeletePostTag
// @Summary 删除一个帖子的标签
// @Description 删除一个帖子的标签
// @Tags 文章和标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param tag_id body int64 true "标签ID"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /postTag/delete [delete]
func DeletePostTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.DeletePostTagParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := logic.DeletePostTag(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// ListTagsByPostID
// @Summary 列出一个帖子对应的所有标签
// @Description 列出一个帖子对应的所有标签
// @Tags 文章和标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id query int64 true "帖子ID"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListTagsReply "返回帖子对应的所有标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /postTag/tags [get]
func ListTagsByPostID(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ListTagsByPostIDParams{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPageSize(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	reply, err := logic.ListTagsByPostID(ctx, params.PostID, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// ListPostInfosByTagID
// @Summary 列出一个标签对应的所有帖子简介信息
// @Description 列出一个标签对应的所有帖子简介信息
// @Tags 文章和标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_id query int64 true "标签ID"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosByTagIDReply "返回帖子对应的所有标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /postTag/infos [get]
func ListPostInfosByTagID(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ListPostInfosByTagIDParams{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPageSize(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	reply, err := logic.ListPostInfosByTagID(ctx, params.TagID, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

/*
	给帖子添加标签
	给帖子去除某标签
	通过postID 查对应标签
	通过tagID 查对应帖子
*/
