package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/dao/redis"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

func DeleteUserStar(ctx *gin.Context, postID int64) error {
	username, _ := getUsername(ctx)
	err := redis.Query.SetPostStar(ctx, postID, username, false)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func UserStarPost(ctx *gin.Context, postID int64) error {
	username, _ := getUsername(ctx)
	err := redis.Query.SetPostStar(ctx, postID, username, true)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func GetUserStar(ctx *gin.Context, postID int64) (bool, error) {
	username, _ := getUsername(ctx)
	result, err := redis.Query.GetPostStarByPostIDAndUserName(ctx, postID, username)
	if err != nil {
		if redis.IsNil(err) {
			_, err = mysql.Query.GetUser_StarByUserNameAndPostId(ctx, db.GetUser_StarByUserNameAndPostIdParams{
				Username: username,
				PostID:   postID,
			})
			if err != nil {
				if mysql.IsNil(err) {
					err = redis.Query.SetPostStar(ctx, postID, username, false)
					if err != nil {
						global.Logger.Error(err.Error())
					}
					return false, nil
				}
			}
		}
		global.Logger.Error(err.Error())
		return false, errcode.ServerErr
	}
	return result, nil
}

func EndurancePostStar(ctx *gin.Context) {
	postStars, err := redis.Query.ListAllPostStarAndSetZero(ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	for postID := range postStars {
		for username, ok := range postStars[postID] {
			err = mysql.Query.DeleteUser_StarByUserNameAndPostID(ctx, db.DeleteUser_StarByUserNameAndPostIDParams{
				Username: username,
				PostID:   postID,
			})
			if err != nil {
				global.Logger.Error(err.Error())
			}
			if ok {
				err = mysql.Query.CreateUser_Star(ctx, db.CreateUser_StarParams{
					Username: username,
					PostID:   postID,
				})
				if err != nil {
					global.Logger.Error(err.Error())
				}
			}
		}
	}
}
