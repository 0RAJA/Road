package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreateUserStar(t *testing.T, arg CreateUser_StarParams) {
	err := testQueries.CreateUser_Star(context.Background(), arg)
	require.NoError(t, err)
	result, err := testGetUserStarByUserNameAndPostID(GetUser_StarByUserNameAndPostIdParams{
		Username: arg.Username,
		PostID:   arg.PostID,
	})
	require.NoError(t, err)
	require.True(t, result > 0)
}

func TestQueries_CreateUser_Star(t *testing.T) {
	user := testCreateUser2(t)
	post := testCreatePost(t)
	err := testQueries.CreateUser_Star(context.Background(), CreateUser_StarParams{
		Username: user.Username,
		PostID:   post.ID,
	})
	require.NoError(t, err)
	result, err := testGetUserStarByUserNameAndPostID(GetUser_StarByUserNameAndPostIdParams{
		Username: user.Username,
		PostID:   post.ID,
	})
	require.NoError(t, err)
	require.True(t, result > 0)
}

func TestQueries_DeleteUser_StarByUserNameAndPostID(t *testing.T) {
	user := testCreateUser2(t)
	post := testCreatePost(t)
	arg := CreateUser_StarParams{
		Username: user.Username,
		PostID:   post.ID,
	}
	testCreateUserStar(t, arg)
	id, err := testGetUserStarByUserNameAndPostID(GetUser_StarByUserNameAndPostIdParams{
		Username: arg.Username,
		PostID:   arg.PostID})
	require.NoError(t, err)
	require.True(t, id > 0)
	err = testQueries.DeleteUser_StarByUserNameAndPostID(context.Background(), DeleteUser_StarByUserNameAndPostIDParams{
		Username: arg.Username,
		PostID:   arg.PostID,
	})
	require.NoError(t, err)
	id, err = testGetUserStarByUserNameAndPostID(GetUser_StarByUserNameAndPostIdParams{
		Username: arg.Username,
		PostID:   arg.PostID})
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Zero(t, id)
}

func testGetUserStarByUserNameAndPostID(arg GetUser_StarByUserNameAndPostIdParams) (int32, error) {
	return testQueries.GetUser_StarByUserNameAndPostId(context.Background(), arg)
}

func TestQueries_GetUser_StarByUserNameAndPostId(t *testing.T) {
	user := testCreateUser2(t)
	post := testCreatePost(t)
	err := testQueries.CreateUser_Star(context.Background(), CreateUser_StarParams{
		Username: user.Username,
		PostID:   post.ID,
	})
	require.NoError(t, err)
	result, err := testGetUserStarByUserNameAndPostID(GetUser_StarByUserNameAndPostIdParams{
		Username: user.Username,
		PostID:   post.ID,
	})
	require.NoError(t, err)
	require.True(t, result > 0)
}
