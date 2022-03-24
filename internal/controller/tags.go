package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// AddTag
// @Summary 增加一个标签
// @Description 增加一个标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_id body int64 true "标签ID"
// @Param tag_name body string true "标签名"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [post]
func AddTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// DeleteTag
// @Summary 删除一个标签
// @Description 删除一个标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_id path int64 true "标签ID"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag/{tag_id} [delete]
func DeleteTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// UpdateTag
// @Summary 修改一个标签的名字
// @Description 修改一个标签的名字
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param tag_id body int64 true "标签ID"
// @Param tag_name body string true "更改后名字"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [put]
func UpdateTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ListTags
// @Summary 列出标签
// @Description 列出标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListTagsReply "返回帖子对应的所有标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [get]
func ListTags(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	tags := make([]logic.Tag, pageSize)
	for i := range tags {
		tags[i] = logic.Tag{
			ID:         utils.RandomInt(1, 100),
			TagName:    utils.RandomOwner(),
			CreateTime: time.Now(),
		}
	}
	response.ToResponseList(tags, len(tags))
}