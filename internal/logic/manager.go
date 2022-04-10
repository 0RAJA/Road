package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	errcode "github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/password"
	"github.com/gin-gonic/gin"
)

// CheckManagerName 获取是否存在对应管理员名
func CheckManagerName(ctx *gin.Context, username string) (bool, *errcode.Error) {
	_, err := mysql.Query.GetManagerByUsername(ctx, username)
	if err != nil {
		if mysql.IsNil(err) {
			return false, nil
		}
		global.Logger.Error(err.Error())
		return false, errcode.ServerErr
	}
	return true, nil
}

func LoginManager(ctx *gin.Context, params LoginManagerParams) (LoginManagerReply, *errcode.Error) {
	manager, err := mysql.Query.GetManagerByUsername(ctx, params.Username)
	if err != nil {
		if mysql.IsNil(err) {
			return LoginManagerReply{}, errcode.ErrUsernameNotFind
		}
		global.Logger.Error(err.Error())
		return LoginManagerReply{}, errcode.ServerErr
	}
	err = password.CheckPassword(params.Password, manager.Password)
	if err != nil {
		return LoginManagerReply{}, errcode.ErrPasswordNotEqual
	}
	token, refreshToken, err1 := generateToken(manager.Username)
	if err1 != nil {
		return LoginManagerReply{}, errcode.UnauthorizedTokenGenerateErr
	}
	return LoginManagerReply{
		Manager: Manager{
			Username:  manager.Username,
			AvatarUrl: manager.AvatarUrl,
		},
		Token:   token,
		ReToken: refreshToken,
	}, nil
}

func AddManager(ctx *gin.Context, request AddManagerRequest) *errcode.Error {
	hashPassword, err := password.HashPassword(request.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrPasswordEncodeFailed.WithDetails(err.Error())
	}
	err = mysql.Query.CreateManager(ctx, db.CreateManagerParams{
		Username:  request.Username,
		Password:  hashPassword,
		AvatarUrl: request.AvatarUrl,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	arg := db.CreateUserParams{
		Username:      request.Username,
		AvatarUrl:     request.AvatarUrl,
		DepositoryUrl: "#",
	}
	if ipaddr, ok := ctx.RemoteIP(); ok {
		arg.Address = ipaddr.String()
	} else {
		arg.Address = ctx.Request.RemoteAddr
	}
	err = mysql.Query.CreateUser(ctx, arg)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return nil
}

func UpdateManagerPassword(ctx *gin.Context, Password string) *errcode.Error {
	username, _ := getUsername(ctx)
	manager, err := mysql.Query.GetManagerByUsername(ctx, username)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrUsernameNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	hashPassword, err := password.HashPassword(Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ErrPasswordEncodeFailed.WithDetails(err.Error())
	}
	if err := password.CheckPassword(hashPassword, manager.Password); err == nil {
		return errcode.ErrPasswordRepeat
	}
	err = mysql.Query.UpdateManagerPassword(ctx, db.UpdateManagerPasswordParams{
		Password: hashPassword,
		Username: username,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func UpdateManagerAvatar(ctx *gin.Context, AvatarUrl string) *errcode.Error {
	username, _ := getUsername(ctx)
	err := mysql.Query.UpdateManagerAvatar(ctx, db.UpdateManagerAvatarParams{
		AvatarUrl: AvatarUrl,
		Username:  username,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func DeleteManager(ctx *gin.Context, username string) *errcode.Error {
	ownUsername, _ := getUsername(ctx)
	if ownUsername == username {
		return errcode.ErrDeleteManagerSelf
	}
	if err := mysql.Query.DeleteManager(ctx, username); err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func ListManagers(ctx *gin.Context, offset, limit int32) ([]db.ListManagerRow, *errcode.Error) {
	managers, err := mysql.Query.ListManager(ctx, db.ListManagerParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return managers, nil
}
