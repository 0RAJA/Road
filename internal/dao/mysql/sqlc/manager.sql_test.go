package db

import (
	"context"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_ListManager(t *testing.T) {
	sum, err := TestQueries.ListManager(context.Background(), ListManagerParams{
		Offset: 0,
		Limit:  1000,
	})
	n := 10
	require.NoError(t, err)
	for i := 0; i < n; i++ {
		testCreateManager2(t)
	}
	result, err := TestQueries.ListManager(context.Background(), ListManagerParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(sum)+10)
}

func TestQueries_GetManagerByUsername(t *testing.T) {
	arg := CreateManagerParams{
		Username:  utils.RandomOwner(),
		Password:  utils.RandomString(10),
		AvatarUrl: utils.RandomString(10),
	}
	testCreateManager(t, arg)
	manager, err := TestQueries.GetManagerByUsername(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, manager)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
	require.Equal(t, manager.AvatarUrl, arg.AvatarUrl)
}

func testGetManagerByUsername(t *testing.T, username string) Manager {
	manager, err := TestQueries.GetManagerByUsername(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, manager)
	return manager
}

func testCreateManager2(t *testing.T) Manager {
	arg := CreateManagerParams{
		Username:  utils.RandomOwner(),
		Password:  utils.RandomString(10),
		AvatarUrl: utils.RandomString(10),
	}
	err := TestQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
	require.Equal(t, manager.AvatarUrl, arg.AvatarUrl)
	return manager
}

func testCreateManager(t *testing.T, arg CreateManagerParams) {
	err := TestQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
	require.Equal(t, manager.AvatarUrl, arg.AvatarUrl)
}

func TestQueries_CreateManager(t *testing.T) {
	arg := CreateManagerParams{
		Username:  utils.RandomOwner(),
		Password:  utils.RandomString(10),
		AvatarUrl: utils.RandomString(10),
	}
	err := TestQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
	require.Equal(t, manager.AvatarUrl, arg.AvatarUrl)
}

func TestQueries_UpdateManager(t *testing.T) {
	manager := testCreateManager2(t)
	arg := UpdateManagerPasswordParams{
		Username: manager.Username,
		Password: utils.RandomOwner(),
	}
	err := TestQueries.UpdateManagerPassword(context.Background(), arg)
	require.NoError(t, err)
	manager2 := testGetManagerByUsername(t, manager.Username)
	require.Equal(t, manager2.Password, arg.Password)
	arg2 := UpdateManagerAvatarParams{
		AvatarUrl: utils.RandomString(10),
		Username:  manager.Username,
	}
	err = TestQueries.UpdateManagerAvatar(context.Background(), arg2)
	require.NoError(t, err)
	manager3 := testGetManagerByUsername(t, manager.Username)
	require.Equal(t, manager3.AvatarUrl, arg2.AvatarUrl)
}
