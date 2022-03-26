package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

func AddPostTag(ctx *gin.Context, params PostTagParams) *errcode.Error {
	err := mysql.Query.CreatePost_Tag(ctx, db.CreatePost_TagParams{
		PostID: params.PostID,
		TagID:  params.TagID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func DeletePostTag(ctx *gin.Context, params DeletePostTagParams) *errcode.Error {
	err := mysql.Query.DeletePost_Tag(ctx, db.DeletePost_TagParams{
		PostID: params.PostID,
		TagID:  params.TagID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func ListTagsByPostID(ctx *gin.Context, postID int64, offset, limit int32) ([]db.Tag, *errcode.Error) {
	tags, err := mysql.Query.ListTagByPostID(ctx, db.ListTagByPostIDParams{
		PostID: postID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return tags, nil
}

func ListPostInfosByTagID(ctx *gin.Context, tagID int64, offset, limit int32) ([]PostInfo, *errcode.Error) {
	isRoot := getRoot(ctx)
	posts, err := mysql.Query.ListPostByTagID(ctx, db.ListPostByTagIDParams{
		TagID:  tagID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	result := make([]PostInfo, 0, len(posts))
	for i := range posts {
		if posts[i].Public || isRoot {
			result = append(result, PostInfo(posts[i]))
		}
	}
	return result, nil
}
