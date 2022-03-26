package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"time"
)

func GetUserInfo(ctx *gin.Context, username string) (User, error) {
	user, err := mysql.Query.GetUserByUsername(ctx, username)
	if err != nil {
		if mysql.IsNil(err) {
			return User{}, errcode.ErrUsernameNotFind
		}
		global.Logger.Error(err.Error())
		return User{}, errcode.ServerErr
	}
	return User(user), nil
}

func ListUsers(ctx *gin.Context, offsite, limit int32) ([]User, error) {
	users, err := mysql.Query.ListUser(ctx, db.ListUserParams{
		Offset: offsite,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	results := make([]User, len(users))
	for i := range users {
		results[i] = User(users[i])
	}
	return results, nil
}

func ListUsersByCreateTime(ctx *gin.Context, startTime, endTime time.Time, offset, limit int32) ([]User, error) {
	users, err := mysql.Query.ListUserByCreateTime(ctx, db.ListUserByCreateTimeParams{
		CreateTime:   startTime,
		CreateTime_2: endTime,
		Offset:       offset,
		Limit:        limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, err
	}
	results := make([]User, len(users))
	for i := range users {
		results[i] = User(users[i])
	}
	return results, nil
}
