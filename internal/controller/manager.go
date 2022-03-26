package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// CheckManagerName
// @Summary 检查管理员名是否存在
// @Description 检查管理员名是否存在
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param username path string true "用户名 3<=len<=50"
// @Success 200 {bool} bool "返回是否存在此管理员"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/check/{username} [get]
func CheckManagerName(ctx *gin.Context) {

}

// LoginManager
// @Summary 管理员登录
// @Description 用于管理员使用账号和密码进行登录
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param username body string true "用户名 3<=len<=50"
// @Param password body string true "密码 6<=len<=32"
// @Success 200 {object} logic.LoginManagerReply "返回用户信息和授权码"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/login [post]
func LoginManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	token := logic.Token{
		Token:     utils.RandomString(10),
		ExpiredAt: time.Now(),
	}
	reToken := logic.ReToken{
		RefreshToken: utils.RandomString(10),
		ExpiredAt:    time.Now(),
	}
	manager := logic.Manager{
		Username: utils.RandomOwner(),
	}
	ret := logic.LoginManagerReply{
		Manager: manager,
		Token:   token,
		ReToken: reToken,
	}
	response.ToResponse(ret)
}

// AddManager
// @Summary 添加管理员
// @Description 添加一个管理员的信息
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username body int64 true "用户名 3<=len<=50"
// @Param password body string true "密码 6<=len<=32"
// @Param avatar_url body string true "头像链接"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/create [post]
func AddManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// UpdateManager
// @Summary 修改管理员头像和密码
// @Description 修改管理员头像和密码，不输入表示不修改
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param password body string true "新密码 6<=len<=32"
// @Param avatar_url body string true "新头像链接"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager [put]
func UpdateManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// DeleteManager
// @Summary 删除管理员
// @Description 删除一个管理员
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username path string true "用户名"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/{username} [post]
func DeleteManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ListManagers
// @Summary 列出管理员
// @Description 列出管理员
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListManagerReply "返回管理员的用户名"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager [get]
func ListManagers(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	managers := make([]logic.Manager, pageSize)
	for i := range managers {
		managers[i] = logic.Manager{
			Username:  utils.RandomOwner(),
			AvatarUrl: "",
		}
	}
	response.ToResponseList(managers, len(managers))
}

/*
	管理员:
		登录
		添加(管理员添加管理员)
		更新
		删除
		列出所有管理员
*/
