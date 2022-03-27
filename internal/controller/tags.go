package controller

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/0RAJA/Road/internal/pkg/conversion"
	"github.com/gin-gonic/gin"
)

// AddTag
// @Summary 增加一个标签
// @Description 增加一个标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_name body string true "标签名"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [post]
func AddTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.AddTagParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	if err := logic.AddTag(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
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
	tagID := conversion.AtoInt64Must(app.GetPath(ctx, "tag_id"))
	if tagID <= 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	if err := logic.DeleteTag(ctx, tagID); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// UpdateTag
// @Summary 修改一个标签的名字
// @Description 修改一个标签的名字
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_id body int64 true "标签ID"
// @Param tag_name body string true "更改后名字"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [put]
func UpdateTag(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.UpdateTagParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	if err := logic.UpdateTag(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// ListTags
// @Summary 列出标签
// @Description 列出标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListTagsReply "返回帖子对应的所有标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag [get]
func ListTags(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.Pagination{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPage(ctx),
	}
	reply, err := logic.ListTags(ctx, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

func checkTagName(tagName string) *errcode.Error {
	if len(tagName) <= 0 || len(tagName) > global.AllSetting.Rule.TagLen {
		return errcode.InvalidParamsErr
	}
	return nil
}

// CheckTagName
// @Summary 判断标签名是否存在
// @Description 判断标签名是否存在
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param tag_name path string true "标签名"
// @Success 200 {object} logic.ListTagsReply "返回帖子对应的所有标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tag/check{tag_name} [get]
func CheckTagName(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	tagName := app.GetPath(ctx, "tag_name")
	if err := checkTagName(tagName); err != nil {
		response.ToErrorResponse(err)
		return
	}
	result, err := logic.CheckTagName(ctx, tagName)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(result)
}
