package logic

import (
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/gin-gonic/gin"
)

func AddComment(ctx *gin.Context, params AddCommentParams) error {
	username, err := getUsername(ctx)
	if err != nil {
		return err
	}
	if username != params.Username {
		return errcode.InsufficientPermissionsErr
	}
	toPost, err := mysql.Query.GetPostByPostID(ctx, params.PostID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if params.ToCommentID > 0 {
		toComment, err := mysql.Query.GetCommentByCommentID(ctx, params.ToCommentID)
		if err != nil {
			if mysql.IsNil(err) {
				return errcode.ErrCommentNotFind
			}
			global.Logger.Error(err.Error())
			return errcode.ServerErr
		}
		if toComment.PostID != toPost.ID {
			return errcode.ErrPostNotEqual
		}
	}
	err = mysql.Query.CreateComment(ctx, db.CreateCommentParams{
		ID:          snowflake.GetID(),
		PostID:      params.PostID,
		Username:    params.Username,
		Content:     params.Content,
		ToCommentID: params.ToCommentID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func DeleteComment(ctx *gin.Context, commentID int64) error {
	username, err := getUsername(ctx)
	if err != nil {
		return err
	}
	isRoot := getRoot(ctx)
	comment, err := mysql.Query.GetCommentByCommentID(ctx, commentID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if comment.Username != username && !isRoot {
		return errcode.InsufficientPermissionsErr
	}
	err = mysql.Query.DeleteCommentByCommentID(ctx, commentID)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func ModifyComment(ctx *gin.Context, params ModifyCommentParams) error {
	username, err := getUsername(ctx)
	if err != nil {
		return err
	}
	isRoot := getRoot(ctx)
	comment, err := mysql.Query.GetCommentByCommentID(ctx, params.CommentID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if comment.Username != username && !isRoot {
		return errcode.InsufficientPermissionsErr
	}
	err = mysql.Query.UpdateCommentByCommentID(ctx, db.UpdateCommentByCommentIDParams{
		Content: params.Content,
		ID:      params.CommentID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func ListComments(ctx *gin.Context, postID int64, offset, limit int32) ([]db.ListCommentByPostIDRow, error) {
	comments, err := mysql.Query.ListCommentByPostID(ctx, db.ListCommentByPostIDParams{
		PostID: postID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return comments, nil
}
