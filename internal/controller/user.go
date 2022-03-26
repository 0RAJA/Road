package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
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
	userInfo := logic.UserInfo{
		Username: utils.RandomOwner(),
	}
	response.ToResponse(userInfo)
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
	pageSize := app.GetPageSize(ctx)
	users := make([]logic.User, pageSize)
	for i := range users {
		users[i] = logic.User{
			Username:   utils.RandomOwner(),
			CreateTime: time.Now(),
			ModifyTime: time.Now(),
		}
	}
	response.ToResponseList(users, len(users))
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
	pageSize := app.GetPageSize(ctx)
	users := make([]logic.User, pageSize)
	for i := range users {
		users[i] = logic.User{
			Username:   utils.RandomOwner(),
			CreateTime: time.Now(),
			ModifyTime: time.Now(),
		}
	}
	response.ToResponseList(users, len(users))
}
