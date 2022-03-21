package db

import (
	"context"
	"database/sql"
	"github.com/0RAJA/Bank/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_GetUserByUsername(t *testing.T) {
	arg := CreateUserParams{
		Username:      util.RandomOwner(),
		AvatarUrl:     util.RandomOwner(),
		DepositoryUrl: util.RandomOwner(),
		Address:       sql.NullString{},
	}
	testCreateUser(t, arg)
	user, err := testQueries.GetUserByUsername(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AvatarUrl, arg.AvatarUrl)
	require.Equal(t, user.DepositoryUrl, arg.DepositoryUrl)
	require.Equal(t, user.Address, arg.Address)
}

func testGetUserByUsername(t *testing.T, username string) User {
	user, err := testQueries.GetUserByUsername(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func testCreateUser2(t *testing.T) User {
	arg := CreateUserParams{
		Username:      util.RandomOwner(),
		AvatarUrl:     util.RandomOwner(),
		DepositoryUrl: util.RandomOwner(),
		Address:       sql.NullString{},
	}
	err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	user := testGetUserByUsername(t, arg.Username)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AvatarUrl, arg.AvatarUrl)
	require.Equal(t, user.DepositoryUrl, arg.DepositoryUrl)
	require.Equal(t, user.Address, arg.Address)
	return user
}

func testCreateUser(t *testing.T, arg CreateUserParams) {
	err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	user := testGetUserByUsername(t, arg.Username)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AvatarUrl, arg.AvatarUrl)
	require.Equal(t, user.DepositoryUrl, arg.DepositoryUrl)
	require.Equal(t, user.Address, arg.Address)
}

func TestQueries_CreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:      util.RandomOwner(),
		AvatarUrl:     util.RandomOwner(),
		DepositoryUrl: util.RandomOwner(),
		Address:       sql.NullString{},
	}
	err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	user := testGetUserByUsername(t, arg.Username)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AvatarUrl, arg.AvatarUrl)
	require.Equal(t, user.DepositoryUrl, arg.DepositoryUrl)
	require.Equal(t, user.Address, arg.Address)
}

func TestQueries_UpdateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:      util.RandomOwner(),
		AvatarUrl:     util.RandomOwner(),
		DepositoryUrl: util.RandomOwner(),
		Address:       sql.NullString{},
	}
	testCreateUser(t, arg)
	arg2 := UpdateUserParams{
		Username:      util.RandomOwner(),
		AvatarUrl:     util.RandomOwner(),
		DepositoryUrl: util.RandomOwner(),
		Address:       sql.NullString{},
	}
	err := testQueries.UpdateUser(context.Background(), arg2)
	require.NoError(t, err)
	user := testGetUserByUsername(t, arg.Username)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.AvatarUrl, arg.AvatarUrl)
	require.Equal(t, user.DepositoryUrl, arg.DepositoryUrl)
	require.Equal(t, user.Address, arg.Address)
}

func TestQueries_ListUser(t *testing.T) {
	n := 10
	//st := time.Now()
	users := make([]CreateUserParams, n)
	for i := 0; i < n; i++ {
		users[i] = CreateUserParams{
			Username:      util.RandomOwner(),
			AvatarUrl:     util.RandomOwner(),
			DepositoryUrl: util.RandomOwner(),
			Address:       sql.NullString{},
		}
		testCreateUser(t, users[i])
	}
	users2, err := testQueries.ListUser(context.Background(), ListUserParams{
		Offset: int32(n / 2),
		Limit:  int32(n / 2),
	})
	require.NoError(t, err)
	require.Len(t, users2, n/2)
}

//必须在user位空时测试
func TestQueries_ListUserByCreateTime(t *testing.T) {
	n := 10
	users := make([]CreateUserParams, n)
	st := time.Now()
	time.Sleep(time.Second)
	for i := 0; i < n; i++ {
		users[i] = CreateUserParams{
			Username:      util.RandomOwner(),
			AvatarUrl:     util.RandomOwner(),
			DepositoryUrl: util.RandomOwner(),
			Address:       sql.NullString{},
		}
		testCreateUser(t, users[i])
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second)
	users2, err := testQueries.ListUserByCreateTime(context.Background(), ListUserByCreateTimeParams{
		CreateTime:   st,
		CreateTime_2: time.Now(),
		Offset:       0,
		Limit:        5,
	})
	require.NoError(t, err)
	require.Len(t, users2, 5)
}
