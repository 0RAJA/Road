package db

import (
	"context"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreateView(t *testing.T) {
	err := TestQueries.CreateView(context.Background(), utils.RandomInt(1, 100))
	require.NoError(t, err)
}
func testCreateView(t *testing.T) {
	err := TestQueries.CreateView(context.Background(), utils.RandomInt(1, 100))
	require.NoError(t, err)
}

func TestQueries_ListViewByCreateTime(t *testing.T) {
	st := time.Now()
	time.Sleep(time.Second)
	n := 10
	for i := 0; i < n; i++ {
		testCreateView(t)
	}
	views, err := TestQueries.ListViewByCreateTime(context.Background(), ListViewByCreateTimeParams{
		CreateTime:   st,
		CreateTime_2: time.Now(),
		Offset:       0,
		Limit:        int32(n),
	})
	require.NoError(t, err)
	require.Len(t, views, n)
	for i := 0; i < n-1; i++ {
		require.True(t, !views[i].CreateTime.Before(views[i+1].CreateTime))
	}
}
