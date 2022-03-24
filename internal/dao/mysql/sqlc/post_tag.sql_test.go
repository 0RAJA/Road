package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreatePostTag3(t *testing.T, post CreatePostParams) PostTag {
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := TestQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.PostID, arg.TagID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}
func testCreatePostTag2(t *testing.T, tag Tag) PostTag {
	post := testCreatePost(t)
	arg := CreatePost_TagParams{
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := TestQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.PostID, arg.TagID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}

func testCreatePostTag(t *testing.T) PostTag {
	post := testCreatePost(t)
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := TestQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
	postTag, err := testGetPostTagByID(arg.PostID, arg.TagID)
	require.NoError(t, err)
	require.NotEmpty(t, postTag)
	return postTag
}

func TestQueries_CreatePost_Tag(t *testing.T) {
	post := testCreatePost(t)
	tag := testCreateTag(t)
	arg := CreatePost_TagParams{
		PostID: post.ID,
		TagID:  tag.ID,
	}
	err := TestQueries.CreatePost_Tag(context.Background(), arg)
	require.NoError(t, err)
}

func TestQueries_DeletePost_TagByID(t *testing.T) {
	postTag := testCreatePostTag(t)
	result, err := testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.NoError(t, err)
	require.Equal(t, result.ID, postTag.ID)
	//先放到垃圾桶里才能真正删除从而解除关系
	err = testDeletePostByID(postTag.PostID) //删除帖子
	require.NoError(t, err)
	result, err = testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.NoError(t, err)
	require.NotEmpty(t, result.ID)
	testModifyPostDeletedByID(t, ModifyPostDeletedByIDParams{ //放进回收站
		Deleted: true,
		ID:      postTag.PostID,
	})
	err = testDeletePostByID(postTag.PostID) //再次删除帖子
	require.NoError(t, err)
	result, err = testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.Error(t, sql.ErrNoRows, err)
	require.Empty(t, result)
	postTag = testCreatePostTag(t)
	testDeleteTagByID(t, postTag.TagID) //删除标签
	result, err = testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	postTag = testCreatePostTag(t)
	result, err = testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.NoError(t, err)
	require.Equal(t, result.ID, postTag.ID)
	err = TestQueries.DeletePost_Tag(context.TODO(), DeletePost_TagParams{
		PostID: postTag.PostID,
		TagID:  postTag.TagID,
	})
	require.NoError(t, err)
	postTag, err = testGetPostTagByID(postTag.PostID, postTag.TagID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, postTag)
}

func TestQueries_ListPostByTagID(t *testing.T) {
	tag := testCreateTag(t)
	for i := 0; i < 10; i++ {
		testCreatePostTag2(t, tag)
	}
	posts, err := TestQueries.ListPostByTagID(context.Background(), ListPostByTagIDParams{
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
		testCreatePostTag3(t, CreatePostParams{
			ID:       post.ID,
			Cover:    "",
			Title:    "",
			Abstract: "",
			Content:  "",
			Public:   false,
		})
	}
	tags, err := TestQueries.ListTagByPostID(context.Background(), ListTagByPostIDParams{
		PostID: post.ID,
		Offset: 0,
		Limit:  100,
	})
	require.NoError(t, err)
	require.Len(t, tags, 10)
}

func testGetPostTagByID(postID, tagID int64) (PostTag, error) {
	return TestQueries.GetPost_Tag(context.Background(), GetPost_TagParams{
		PostID: postID,
		TagID:  tagID,
	})
}

func TestQueries_GetPost_TagByID(t *testing.T) {
	postTag := testCreatePostTag(t)
	postTag2, err := TestQueries.GetPost_Tag(context.Background(), GetPost_TagParams{
		PostID: postTag.PostID,
		TagID:  postTag.TagID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, postTag2)
	require.Equal(t, postTag, postTag2)
}
