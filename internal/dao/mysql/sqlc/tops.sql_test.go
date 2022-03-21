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
	err := testQueries.CreateTop(context.Background(), arg)
	require.NoError(t, err)
	top, err := testQueries.GetTopByTopID(context.Background(), arg.ID)
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
	err := testQueries.CreateTop(context.Background(), arg)
	require.NoError(t, err)
	top, err := testQueries.GetTopByTopID(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, top)
	require.Equal(t, top.ID, arg.ID)
}

func testDeleteTopByTopID(t *testing.T, topID int64) {
	err := testQueries.DeleteTopByTopID(context.Background(), topID)
	require.NoError(t, err)
	_, err = testGetTopByID(topID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestQueries_DeleteTopByTopID(t *testing.T) {
	top := testCreateTop(t)
	err := testQueries.DeleteTopByTopID(context.Background(), top.ID)
	require.NoError(t, err)
	_, err = testGetTopByID(top.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func testGetTopByID(topID int64) (Top, error) {
	return testQueries.GetTopByTopID(context.Background(), topID)
}

func TestQueries_GetTopByTopID(t *testing.T) {
	top := testCreateTop(t)
	top1, err := testQueries.GetTopByTopID(context.Background(), top.ID)
	require.NoError(t, err)
	require.NotEmpty(t, top1)
	require.Equal(t, top1.ID, top.ID)
}
