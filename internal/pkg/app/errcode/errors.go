package errcode

import "github.com/0RAJA/Bank/pkg/app/errcode"

//以 7,8 开头表示个人项目错误码

var (
	ErrCommentNotFind       = errcode.NewError(70000000, "评论不存在")
	ErrPostNotFind          = errcode.NewError(70000001, "不存在此文章")
	ErrPostNotEqual         = errcode.NewError(70000002, "回复的评论不在回复的文章下")
	ErrUsernameNotFind      = errcode.NewError(70000003, "用户不存在")
	ErrPasswordNotEqual     = errcode.NewError(70000004, "密码错误")
	ErrPasswordEncodeFailed = errcode.NewError(70000005, "加密密码失败")
	ErrPasswordRepeat       = errcode.NewError(70000006, "密码与原密码相同")
	ErrDeleteManagerSelf    = errcode.NewError(70000007, "不能删除自己")
	ErrStateRepeat          = errcode.NewError(70000008, "状态设置重复")
	ErrDeletedState         = errcode.NewError(70000009, "删除状态异常")
	ErrListPostInfosOptions = errcode.NewError(70000010, "列出帖子信息选项异常")
	ErrAuthorizationFailed  = errcode.NewError(70000011, "授权失败")
	ErrTokenNotExpired      = errcode.NewError(70000012, "token未过期")
)

var (
	ExtErr              = errcode.NewError(70000001, "file suffix is not supported")
	FileSizeErr         = errcode.NewError(70000002, "exceeded maximum file limit")
	CreatePathErr       = errcode.NewError(70000003, "failed to create save directory")
	CompetenceErr       = errcode.NewError(70000004, "insufficient file permissions")
	RepeatedFileTypeErr = errcode.NewError(70000005, "DuplicateFileType")
)
