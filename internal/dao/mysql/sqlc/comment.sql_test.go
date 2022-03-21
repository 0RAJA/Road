package db

import (
	"context"
	"database/sql"
	"github.com/0RAJA/Bank/db/util"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func testCreateComment2(t *testing.T, post Post) Comment {
	user := testCreateUser2(t)
	arg := CreateCommentParams{
		ID:       snowflake.GetID(),
		PostID:   post.ID,
		Username: user.Username,
		Content:  util.RandomString(10),
	}
	st := time.Now()
	err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	comment, err := testGetCommentByCommentID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, comment.Username, arg.Username)
	require.Equal(t, comment.PostID, arg.PostID)
	require.Equal(t, comment.ToCommentID, arg.ToCommentID)
	require.WithinDuration(t, comment.CreateTime, st, time.Second)
	return comment
}
func testCreateComment(t *testing.T) Comment {
	post := testCreatePost(t)
	user := testCreateUser2(t)
	arg := CreateCommentParams{
		ID:       snowflake.GetID(),
		PostID:   post.ID,
		Username: user.Username,
		Content:  util.RandomString(10),
	}
	st := time.Now()
	err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	comment, err := testGetCommentByCommentID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, comment.Username, arg.Username)
	require.Equal(t, comment.PostID, arg.PostID)
	require.Equal(t, comment.ToCommentID, arg.ToCommentID)
	require.WithinDuration(t, comment.CreateTime, st, time.Second)
	return comment
}

func TestQueries_CreateComment(t *testing.T) {
	post := testCreatePost(t)
	user := testCreateUser2(t)
	arg := CreateCommentParams{
		ID:       snowflake.GetID(),
		PostID:   post.ID,
		Username: user.Username,
		Content:  util.RandomString(10),
	}
	st := time.Now()
	err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	comment, err := testGetCommentByCommentID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, comment.Username, arg.Username)
	require.Equal(t, comment.PostID, arg.PostID)
	require.Equal(t, comment.ToCommentID, arg.ToCommentID)
	require.WithinDuration(t, comment.CreateTime, st, time.Second)
}

func TestQueries_DeleteCommentByCommentID(t *testing.T) {
	comment := testCreateComment(t)
	err := testQueries.DeleteCommentByCommentID(context.Background(), comment.ID)
	require.NoError(t, err)
	comment1, err := testGetCommentByCommentID(comment.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, comment1)
}

func testGetCommentByCommentID(commentID int64) (Comment, error) {
	return testQueries.GetCommentByCommentID(context.Background(), commentID)
}

func TestQueries_GetCommentByCommentID(t *testing.T) {
	comment := testCreateComment(t)
	comment1, err := testGetCommentByCommentID(comment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment.ID, comment1.ID)
}

func TestQueries_ListCommentByPostID(t *testing.T) {
	post := testCreatePost(t)
	num, err := testQueries.ListCommentByPostID(context.Background(), ListCommentByPostIDParams{
		PostID: post.ID,
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, num, 0)
	for i := 0; i < 10; i++ {
		testCreateComment2(t, post)
	}
	num, err = testQueries.ListCommentByPostID(context.Background(), ListCommentByPostIDParams{
		PostID: post.ID,
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, num, 10)
}

func TestQueries_UpdateCommentByCommentID(t *testing.T) {
	comment := testCreateComment(t)
	arg := UpdateCommentByCommentIDParams{
		Content: util.RandomString(10),
		ID:      comment.ID,
	}
	mt := time.Now()
	err := testQueries.UpdateCommentByCommentID(context.Background(), arg)
	require.NoError(t, err)
	comment, err = testGetCommentByCommentID(comment.ID)
	require.NoError(t, err)
	require.Equal(t, comment.Content, arg.Content)
	require.WithinDuration(t, comment.ModifyTime, mt, time.Second)
}
