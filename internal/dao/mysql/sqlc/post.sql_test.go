package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/0RAJA/Bank/db/util"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreatePost(t *testing.T) {
	arg := CreatePostParams{
		ID:       snowflake.GetID(),
		Cover:    util.RandomString(10),
		Title:    util.RandomString(30),
		Abstract: util.RandomString(30),
		Content:  util.RandomString(40),
		Public:   true,
	}
	st := time.Now()
	err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	tag, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, tag.ID, arg.ID)
	require.Equal(t, tag.Abstract, arg.Abstract)
	require.Equal(t, tag.Cover, arg.Cover)
	require.Equal(t, tag.Title, arg.Title)
	require.Equal(t, tag.Content, arg.Content)
	require.False(t, tag.Deleted)
	require.Zero(t, tag.VisitedNum, tag.StarNum)
	require.WithinDuration(t, tag.CreateTime, st, time.Second)
	require.WithinDuration(t, tag.ModifyTime, st, time.Second)
}

func testGetPostByID(postID int64) (Post, error) {
	return testQueries.GetPostByPostID(context.Background(), postID)
}

func testDeletePostByID(postID int64) error {
	return testQueries.DeletePostByPostID(context.Background(), postID)
}

func TestQueries_DeletePostByPostID(t *testing.T) {
	post := testCreatePost(t)
	err := testQueries.DeletePostByPostID(context.Background(), post.ID)
	require.NoError(t, err)
	post1, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post1)
}

func testCreatePost2(t *testing.T, arg CreatePostParams) {
	st := time.Now()
	err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	tag, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, tag.ID, arg.ID)
	require.Equal(t, tag.Abstract, arg.Abstract)
	require.Equal(t, tag.Cover, arg.Cover)
	require.Equal(t, tag.Title, arg.Title)
	require.Equal(t, tag.Content, arg.Content)
	require.False(t, tag.Deleted)
	require.Zero(t, tag.VisitedNum, tag.StarNum)
	require.WithinDuration(t, tag.CreateTime, st, time.Second)
	require.WithinDuration(t, tag.ModifyTime, st, time.Second)
}
func testCreatePost(t *testing.T) Post {
	arg := CreatePostParams{
		ID:       snowflake.GetID(),
		Cover:    util.RandomString(10),
		Title:    util.RandomString(30),
		Abstract: util.RandomString(30),
		Content:  util.RandomString(40),
		Public:   true,
	}
	st := time.Now()
	err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	tag, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, tag.ID, arg.ID)
	require.Equal(t, tag.Abstract, arg.Abstract)
	require.Equal(t, tag.Cover, arg.Cover)
	require.Equal(t, tag.Title, arg.Title)
	require.Equal(t, tag.Content, arg.Content)
	require.False(t, tag.Deleted)
	require.Zero(t, tag.VisitedNum, tag.StarNum)
	require.WithinDuration(t, tag.CreateTime, st, time.Second)
	require.WithinDuration(t, tag.ModifyTime, st, time.Second)
	return tag
}
func TestQueries_GetPostByPostID(t *testing.T) {
	post1 := testCreatePost(t)
	post, err := testQueries.GetPostByPostID(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.ID, post1.ID)
	post, err = testQueries.GetPostByPostID(context.Background(), 1)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, post)
}

func TestQueries_ListPostBySearchKey(t *testing.T) {
	post := testCreatePost(t)
	rows, err := testQueries.ListPostBySearchKey(context.Background(), ListPostBySearchKeyParams{
		Title:    fmt.Sprintf("%%%s%%", post.Title[1:len(post.Title)-1]),
		Abstract: fmt.Sprintf("%%%s%%", post.Abstract[1:len(post.Title)-1]),
		Offset:   0,
		Limit:    1000,
	})
	require.NoError(t, err)
	require.Len(t, rows, 1)
	require.Equal(t, rows[0].ID, post.ID)
	key := "abcdefghijk"
	for i := 0; i < 10; i++ {
		arg := CreatePostParams{
			ID:       snowflake.GetID(),
			Cover:    util.RandomOwner(),
			Title:    util.RandomOwner() + key + util.RandomOwner(),
			Abstract: util.RandomOwner(),
			Content:  util.RandomOwner(),
			Public:   true,
		}
		testCreatePost2(t, arg)
	}
	rows, err = testQueries.ListPostBySearchKey(context.Background(), ListPostBySearchKeyParams{
		Title:    fmt.Sprintf("%%%s%%", key),
		Abstract: key,
		Offset:   0,
		Limit:    1000,
	})
	require.NoError(t, err)
	require.True(t, len(rows) >= 10)
}

func TestQueries_ListPostByStartTime(t *testing.T) {
	st := time.Now()
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		arg := CreatePostParams{
			ID:       snowflake.GetID(),
			Cover:    util.RandomOwner(),
			Title:    util.RandomOwner(),
			Abstract: util.RandomOwner(),
			Content:  util.RandomOwner(),
			Public:   true,
		}
		testCreatePost2(t, arg)
	}
	time.Sleep(2 * time.Second)
	post, err := testQueries.ListPostByStartTime(context.Background(), ListPostByStartTimeParams{
		CreateTime:   st,
		CreateTime_2: time.Now(),
		Offset:       0,
		Limit:        1000,
	})
	require.NoError(t, err)
	require.Len(t, post, 10)
}

func TestQueries_ListPostTopping(t *testing.T) {
	sum, err := testQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	n := 10
	tops := make([]Top, n)
	for i := 0; i < n; i++ {
		tops[i] = testCreateTop(t)
	}
	tops1, err := testQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, tops1, len(sum)+n)
	for i := range tops {
		testDeleteTopByTopID(t, tops[i].ID)
	}
	tops1, err = testQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	require.Len(t, tops1, len(sum))
}

func TestQueries_ListPostDeleted(t *testing.T) {
	posts1, err := testQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]Post, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostDeletedByID(t, ModifyPostDeletedByIDParams{
			Deleted: true,
			ID:      posts[i].ID,
		})
	}
	result, err := testQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1)+n)
	for i := range posts {
		testModifyPostDeletedByID(t, ModifyPostDeletedByIDParams{
			Deleted: false,
			ID:      posts[i].ID,
		})
	}
	result, err = testQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func TestQueries_ListPostPrivate(t *testing.T) {
	posts1, err := testQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]Post, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: false,
			ID:     posts[i].ID,
		})
	}
	result, err := testQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1)+n)
	for i := range posts {
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: true,
			ID:     posts[i].ID,
		})
	}
	result, err = testQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func TestQueries_ListPostPublic(t *testing.T) {
	posts1, err := testQueries.ListPostPublic(context.Background(), ListPostPublicParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]Post, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: true,
			ID:     posts[i].ID,
		})
	}
	result, err := testQueries.ListPostPublic(context.Background(), ListPostPublicParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1)+n)
	for i := range posts {
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: false,
			ID:     posts[i].ID,
		})
	}
	result, err = testQueries.ListPostPublic(context.Background(), ListPostPublicParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func testModifyPostDeletedByID(t *testing.T, arg ModifyPostDeletedByIDParams) {
	err := testQueries.ModifyPostDeletedByID(context.Background(), arg)
	require.NoError(t, err)
	post2, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.True(t, post2.Deleted == arg.Deleted)
}

func TestQueries_ModifyPostDeletedByID(t *testing.T) {
	post := testCreatePost(t)
	require.False(t, post.Deleted)
	err := testQueries.ModifyPostDeletedByID(context.Background(), ModifyPostDeletedByIDParams{
		Deleted: true,
		ID:      post.ID,
	})
	require.NoError(t, err)
	post2, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.True(t, post2.Deleted)
	err = testQueries.ModifyPostDeletedByID(context.Background(), ModifyPostDeletedByIDParams{
		Deleted: false,
		ID:      post.ID,
	})
	require.NoError(t, err)
	post3, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.False(t, post3.Deleted)
}

func testModifyPostPublicByID(t *testing.T, arg ModifyPostPublicByIDParams) {
	err := testQueries.ModifyPostPublicByID(context.Background(), arg)
	require.NoError(t, err)
	post2, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.True(t, post2.Public == arg.Public)
}

func TestQueries_ModifyPostPublicByID(t *testing.T) {
	post := testCreatePost(t)
	require.True(t, post.Public)
	err := testQueries.ModifyPostPublicByID(context.Background(), ModifyPostPublicByIDParams{
		Public: false,
		ID:     post.ID,
	})
	require.NoError(t, err)
	post2, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.False(t, post2.Public)
	err = testQueries.ModifyPostPublicByID(context.Background(), ModifyPostPublicByIDParams{
		Public: true,
		ID:     post.ID,
	})
	require.NoError(t, err)
	post3, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.True(t, post3.Public)
}

func TestQueries_UpdatePostByPostID(t *testing.T) {
	post := testCreatePost(t)
	arg := UpdatePostByPostIDParams{
		Cover:    util.RandomOwner(),
		Title:    util.RandomOwner(),
		Abstract: util.RandomOwner(),
		Content:  util.RandomOwner(),
		Public:   true,
		ID:       post.ID,
	}
	st := time.Now()
	err := testQueries.UpdatePostByPostID(context.Background(), arg)
	require.NoError(t, err)
	post2, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)
	require.Equal(t, arg.Cover, post2.Cover)
	require.Equal(t, arg.Title, post2.Title)
	require.Equal(t, arg.Abstract, post2.Abstract)
	require.Equal(t, arg.Content, post2.Content)
	require.WithinDuration(t, post2.ModifyTime, st, time.Second)
}
