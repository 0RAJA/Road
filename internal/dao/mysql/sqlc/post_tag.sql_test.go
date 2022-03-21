package db

import (
	"context"
	"database/sql"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreatePostTag3(t *testing.T, post Post) PostTag {
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := testQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}
func testCreatePostTag2(t *testing.T, tag Tag) PostTag {
	post := testCreatePost(t)
	arg := CreatePost_TagParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := testQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}

func testCreatePostTag(t *testing.T) PostTag {
	post := testCreatePost(t)
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := testQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}

func TestQueries_CreatePost_Tag(t *testing.T) {
	post := testCreatePost(t)
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		ID:     snowflake.GetID(),
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := testQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
}

func TestQueries_DeletePost_TagByID(t *testing.T) {
	postTag := testCreatePostTag(t)
	result, err := testGetPostTagByID(postTag.ID)
	require.NoError(t, err)
	require.Equal(t, result.ID, postTag.ID)
	//先放到垃圾桶里才能真正删除从而解除关系
	err = testDeletePostByID(postTag.PostID) //删除帖子
	require.NoError(t, err)
	result, err = testGetPostTagByID(postTag.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result.ID)
	testModifyPostDeletedByID(t, ModifyPostDeletedByIDParams{ //放进回收站
		Deleted: true,
		ID:      postTag.PostID,
	})
	err = testDeletePostByID(postTag.PostID) //再次删除帖子
	require.NoError(t, err)
	result, err = testGetPostTagByID(postTag.ID)
	require.Error(t, sql.ErrNoRows, err)
	require.Empty(t, result)
	postTag = testCreatePostTag(t)
	testDeleteTagByID(t, postTag.TagID) //删除标签
	result, err = testGetPostTagByID(postTag.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	postTag = testCreatePostTag(t)
	result, err = testGetPostTagByID(postTag.ID)
	require.NoError(t, err)
	require.Equal(t, result.ID, postTag.ID)
	err = testQueries.DeletePost_TagByID(context.TODO(), postTag.ID)
	require.NoError(t, err)
	postTag, err = testGetPostTagByID(postTag.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, postTag)
}

func TestQueries_ListPostByTagID(t *testing.T) {
	tag := testCreateTag(t)
	for i := 0; i < 10; i++ {
		testCreatePostTag2(t, tag)
	}
	posts, err := testQueries.ListPostByTagID(context.Background(), ListPostByTagIDParams{
		TagID:  tag.ID,
		Offset: 0,
		Limit:  100,
	})
	require.NoError(t, err)
	require.Len(t, posts, 10)
}

func TestQueries_ListTagByPostID(t *testing.T) {
	post := testCreatePost(t)
	for i := 0; i < 10; i++ {
		testCreatePostTag3(t, post)
	}
	tags, err := testQueries.ListTagByPostID(context.Background(), ListTagByPostIDParams{
		PostID: post.ID,
		Offset: 0,
		Limit:  100,
	})
	require.NoError(t, err)
	require.Len(t, tags, 10)
}

func testGetPostTagByID(id int64) (PostTag, error) {
	return testQueries.GetPost_TagById(context.Background(), id)
}

func TestQueries_GetPost_TagByID(t *testing.T) {
	postTag := testCreatePostTag(t)
	postTag2, err := testQueries.GetPost_TagById(context.Background(), postTag.ID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag2)
	require.Equal(t, postTag, postTag2)
}
