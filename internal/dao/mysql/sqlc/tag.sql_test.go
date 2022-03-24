package db

import (
	"context"
	"database/sql"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/0RAJA/Road/internal/pkg/times"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func testCreateTag(t *testing.T) Tag {
	id := snowflake.GetID()
	err := TestQueries.CreateTag(context.Background(), CreateTagParams{
		ID:      id,
		TagName: utils.RandomOwner(),
	})
	require.NoError(t, err)
	tag, err := testGetTagById(id)
	require.NoError(t, err)
	return tag
}

func TestQueries_CreateTag(t *testing.T) {
	arg := CreateTagParams{
		ID:      snowflake.GetID(),
		TagName: utils.RandomOwner(),
	}
	st := times.GetNowTime()
	err := TestQueries.CreateTag(context.Background(), arg)
	require.NoError(t, err)
	tag, err := testGetTagById(arg.ID)
	require.NoError(t, err)
	require.Equal(t, tag.ID, arg.ID)
	require.Equal(t, tag.TagName, arg.TagName)
	require.WithinDuration(t, st, tag.CreateTime, time.Second)
}

func testGetTagById(id int64) (Tag, error) {
	return TestQueries.GetTagById(context.Background(), id)
}

func TestQueries_GetTagById(t *testing.T) {
	tag1 := testCreateTag(t)
	tag, err := TestQueries.GetTagById(context.Background(), tag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	require.Equal(t, tag.ID, tag1.ID)
	tag, err = TestQueries.GetTagById(context.Background(), 1)
	require.Error(t, err)
	require.Empty(t, tag)
}

func TestQueries_UpdateTag(t *testing.T) {
	tag1 := testCreateTag(t)
	arg := UpdateTagParams{
		TagName: "hhh",
		ID:      tag1.ID,
	}
	err := TestQueries.UpdateTag(context.Background(), arg)
	require.NoError(t, err)
	tag, err := testGetTagById(arg.ID)
	require.NoError(t, err)
	require.Equal(t, tag.TagName, arg.TagName)
}

func TestQueries_ListTag(t *testing.T) {
	for i := 0; i < 10; i++ {
		testCreateTag(t)
	}
	args, err := TestQueries.ListTag(context.Background(), ListTagParams{
		Offset: 0,
		Limit:  10,
	})
	require.NoError(t, err)
	require.Len(t, args, 10)
}

func testDeleteTagByID(t *testing.T, tagID int64) {
	err := TestQueries.DeleteTagByTagID(context.Background(), tagID)
	require.NoError(t, err)
}

func TestQueries_DeleteTagByTagID(t *testing.T) {
	tag := testCreateTag(t)
	err := TestQueries.DeleteTagByTagID(context.Background(), tag.ID)
	require.NoError(t, err)
	tag1, err := testGetTagById(tag.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, tag1)
}
