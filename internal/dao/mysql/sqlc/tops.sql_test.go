package db

import (
	"context"
	"database/sql"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreateTop(t *testing.T) Top {
	post := testCreatePost(t)
	arg := CreateTopParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
	}
	err := TestQueries.CreateTop(context.Background(), arg)
	require.NoError(t, err)
	top, err := TestQueries.GetTopByPostID(context.Background(), arg.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, top)
	require.Equal(t, top.ID, arg.ID)
	return top
}

func TestQueries_CreateTop(t *testing.T) {
	post := testCreatePost(t)
	arg := CreateTopParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
	}
	err := TestQueries.CreateTop(context.Background(), arg)
	require.NoError(t, err)
	top, err := TestQueries.GetTopByPostID(context.Background(), arg.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, top)
	require.Equal(t, top.ID, arg.ID)
}

func testDeleteTopByTopID(t *testing.T, postID int64) {
	err := TestQueries.DeleteTopByPostID(context.Background(), postID)
	require.NoError(t, err)
	_, err = testGetTopByPostID(postID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestQueries_DeleteTopByPostID(t *testing.T) {
	top := testCreateTop(t)
	err := TestQueries.DeleteTopByPostID(context.Background(), top.PostID)
	require.NoError(t, err)
	_, err = testGetTopByPostID(top.PostID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func testGetTopByPostID(postID int64) (Top, error) {
	return TestQueries.GetTopByPostID(context.Background(), postID)
}

func TestQueries_GetTopByTopID(t *testing.T) {
	top := testCreateTop(t)
	top1, err := TestQueries.GetTopByPostID(context.Background(), top.PostID)
	require.NoError(t, err)
	require.NotEmpty(t, top1)
	require.Equal(t, top1.ID, top.ID)
}
