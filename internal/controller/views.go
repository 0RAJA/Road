package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
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
	pageSize := app.GetPageSize(ctx)
	views := make([]logic.View, pageSize)
	for i := range views {
		views[i] = logic.View{
			ID:         utils.RandomInt(1, 100),
			ViewsNum:   utils.RandomInt(1, 1000),
			CreateTime: time.Now(),
		}
	}
	response.ToResponseList(views, len(views))
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
	response.ToResponse(utils.RandomInt(1, 1000))
}
