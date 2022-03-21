package db

import (
	"context"
	"github.com/0RAJA/Bank/db/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_ListManager(t *testing.T) {
	sum, err := testQueries.ListManager(context.Background(), ListManagerParams{
		Offset: 0,
		Limit:  1000,
	})
	n := 10
	require.NoError(t, err)
	for i := 0; i < n; i++ {
		testCreateManager2(t)
	}
	result, err := testQueries.ListManager(context.Background(), ListManagerParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(sum)+10)
}

func TestQueries_GetManagerByUsername(t *testing.T) {
	arg := CreateManagerParams{
		Username: util.RandomOwner(),
		Password: util.RandomString(10),
	}
	testCreateManager(t, arg)
	manager, err := testQueries.GetManagerByUsername(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, manager)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
}

func testGetManagerByUsername(t *testing.T, username string) Manager {
	manager, err := testQueries.GetManagerByUsername(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, manager)
	return manager
}

func testCreateManager2(t *testing.T) Manager {
	arg := CreateManagerParams{
		Username: util.RandomOwner(),
		Password: util.RandomString(10),
	}
	err := testQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
	return manager
}

func testCreateManager(t *testing.T, arg CreateManagerParams) {
	err := testQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
}

func TestQueries_CreateManager(t *testing.T) {
	arg := CreateManagerParams{
		Username: util.RandomOwner(),
		Password: util.RandomString(10),
	}
	err := testQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	manager := testGetManagerByUsername(t, arg.Username)
	require.Equal(t, manager.Username, arg.Username)
	require.Equal(t, manager.Password, arg.Password)
}

func TestQueries_UpdateManager(t *testing.T) {
	manager := testCreateManager2(t)
	arg := UpdateManagerParams{
		Username: manager.Username,
		Password: util.RandomOwner(),
	}
	err := testQueries.UpdateManager(context.Background(), arg)
	require.NoError(t, err)
	manager2 := testGetManagerByUsername(t, manager.Username)
	require.Equal(t, manager2.Password, arg.Password)
}
