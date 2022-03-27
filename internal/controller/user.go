package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/gin-gonic/gin"
)

// GetUserInfo
// @Summary 获取用户信息
// @Description 获取用户的用户名，头像链接，仓库链接，ip地址，创建时间，修改时间
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username path string true "用户名"
// @Success 200 {object} logic.User  "用户信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /user/{username} [get]
func GetUserInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	username := app.GetPath(ctx, "username")
	if err := checkUsername(username); err != nil {
		response.ToErrorResponse(err)
		return
	}
	reply, err := logic.GetUserInfo(ctx, username)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(reply)
}

// ListUsers
// @Summary 列出用户信息
// @Description 列出用户的用户名，头像链接，仓库链接，ip地址，创建时间，修改时间
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListUsersReply  "用户信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /user/users [get]
func ListUsers(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.Pagination{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPage(ctx),
	}
	reply, err := logic.ListUsers(ctx, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// ListUsersByCreateTime
// @Summary 通过时间段来检索用户信息
// @Description 通过时间段来检索用户信息
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param start_time query time.Time true "起始时间 (2002-03-26)"
// @Param end_time query time.Time true "结束时间 (2002-03-26)"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListUsersReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /user/createTime [get]
func ListUsersByCreateTime(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ListUsersByCreateTimeParams{
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
	reply, err := logic.ListUsersByCreateTime(ctx, params.StartTime, params.EndTime, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}
