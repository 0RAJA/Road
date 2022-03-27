package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/0RAJA/Road/internal/pkg/conversion"
	"github.com/gin-gonic/gin"
)

// ListViewsByCreateTime
// @Summary 获取时间段内的访问数量
// @Description 获取时间段内的访问数量
// @Tags views
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param start_time query time.Time true "起始时间 (2002-03-26)"
// @Param end_time query time.Time true "结束时间 (2002-03-26)"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListViewsByCreateTimeReply  "指定时间内的所有访问数"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /views/all [get]
func ListViewsByCreateTime(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.SearchPostInfosByCreateTimeParam{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPage(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.ListViewsByCreateTime(ctx, params.StartTime, params.EndTime, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// GetGrowViewsByPostID
// @Summary 获取指定帖子的新增访问量
// @Description 获取指定帖子的新增访问量
// @Tags views
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {int64} int64  "返回对应帖子的新增访问量"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /views/post/{post_id} [get]
func GetGrowViewsByPostID(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postID := conversion.AtoInt64Must(app.GetPath(ctx, "post_id"))
	if postID <= 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	num, err := logic.GetGrowViewsByPostID(ctx, postID)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(num)
}
