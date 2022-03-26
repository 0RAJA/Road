package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/gin-gonic/gin"
)

func CheckTagName(ctx *gin.Context, tagName string) (bool, error) {
	_, err := mysql.Query.GetTagByName(ctx, tagName)
	if err != nil {
		if mysql.IsNil(err) {
			return false, nil
		}
		global.Logger.Error(err.Error())
		return false, errcode.ServerErr
	}
	return true, nil
}

func AddTag(ctx *gin.Context, params AddTagParams) error {
	err := mysql.Query.CreateTag(ctx, db.CreateTagParams{
		ID:      snowflake.GetID(),
		TagName: params.TagName,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func DeleteTag(ctx *gin.Context, tagID int64) error {
	err := mysql.Query.DeleteTagByTagID(ctx, tagID)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func UpdateTag(ctx *gin.Context, params UpdateTagParams) error {
	err := mysql.Query.UpdateTag(ctx, db.UpdateTagParams{
		TagName: params.TagName,
		ID:      params.TagID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func ListTags(ctx *gin.Context, offsite, limit int32) ([]db.Tag, error) {
	tags, err := mysql.Query.ListTag(ctx, db.ListTagParams{
		Offset: offsite,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return tags, nil
}
