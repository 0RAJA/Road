package controller

import (
	"github.com/0RAJA/Road/internal/global"
	logic "github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/gin-gonic/gin"
)

func checkUsername(username string) *errcode.Error {
	if len(username) <= 0 || len(username) > global.AllSetting.Rule.UsernameLen {
		return errcode.ErrUsernameLengthErr
	}
	return nil
}
func checkPassword(password string) *errcode.Error {
	if len(password) <= 0 || len(password) > global.AllSetting.Rule.PasswordLen {
		return errcode.ErrUsernameLengthErr
	}
	return nil
}

// CheckManagerName
// @Summary 检查管理员名是否存在
// @Description 检查管理员名是否存在
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username path string true "用户名 3<=len<=50"
// @Success 200 {bool} bool "返回是否存在此管理员"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/check/{username} [get]
func CheckManagerName(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	username := app.GetPath(ctx, "username")
	if err := checkUsername(username); err != nil {
		response.ToErrorResponse(err)
		return
	}
	ok, err := logic.CheckManagerName(ctx, username)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(ok)
}

// LoginManager
// @Summary 管理员登录
// @Description 用于管理员使用账号和密码进行登录
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param username body string true "用户名 3<=len<=50"
// @Param password body string true "密码 6<=len<=32"
// @Success 200 {object} logic.LoginManagerReply "返回用户信息和授权码"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /manager/login [post]
func LoginManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.LoginManagerParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := checkUsername(params.Username); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := checkPassword(params.Password); err != nil {
		response.ToErrorResponse(err)
		return
	}
	reply, err := logic.LoginManager(ctx, params)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(reply)
}

// AddManager
// @Summary 添加管理员
// @Description 添加一个管理员的信息
// @Tags 管理员
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
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
	params := logic.AddManagerRequest{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := checkUsername(params.Username); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := checkPassword(params.Password); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := logic.AddManager(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// UpdateManager
// @Summary 修改管理员头像和密码
// @Description 修改管理员头像和密码，空字符串表示不修改
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
	params := logic.UpdateManagerRequest{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	var err *errcode.Error
	if params.Password != "" {
		if err = checkPassword(params.Password); err != nil {
			response.ToErrorResponse(err)
			return
		}
		if err = logic.UpdateManagerPassword(ctx, params.Password); err != nil {
			response.ToErrorResponse(err)
			return
		}
	}
	if params.AvatarUrl != "" {
		if err = logic.UpdateManagerAvatar(ctx, params.AvatarUrl); err != nil {
			response.ToErrorResponse(err)
			return
		}
	}
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
// @Router /manager/{username} [delete]
func DeleteManager(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	username := app.GetPath(ctx, "username")
	if err := checkUsername(username); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := logic.DeleteManager(ctx, username); err != nil {
		response.ToErrorResponse(err)
		return
	}
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
	params := logic.Pagination{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPageSize(ctx),
	}
	reply, err := logic.ListManagers(ctx, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

/*
	管理员:
		登录
		添加(管理员添加管理员)
		更新
		删除
		列出所有管理员
*/
