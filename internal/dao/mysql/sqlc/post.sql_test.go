package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreatePost(t *testing.T) {
	arg := CreatePostParams{
		ID:       snowflake.GetID(),
		Cover:    utils.RandomString(10),
		Title:    utils.RandomString(30),
		Abstract: utils.RandomString(30),
		Content:  utils.RandomString(40),
		Public:   true,
	}
	st := time.Now()
	err := TestQueries.CreatePost(context.Background(), arg)
	testCreatePostNum(t, arg.ID)
	require.NoError(t, err)
	post, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, post.ID, arg.ID)
	require.Equal(t, post.Abstract, arg.Abstract)
	require.Equal(t, post.Cover, arg.Cover)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Content, arg.Content)
	require.False(t, post.Deleted)
	require.Zero(t, post.VisitedNum, post.StarNum)
	require.WithinDuration(t, post.CreateTime, st, time.Second)
	require.WithinDuration(t, post.ModifyTime, st, time.Second)
}

func testGetPostByID(postID int64) (GetPostByPostIDRow, error) {
	return TestQueries.GetPostByPostID(context.Background(), postID)
}

func testDeletePostByID(postID int64) error {
	return TestQueries.DeletePostByPostID(context.Background(), postID)
}

func TestQueries_DeletePostByPostID(t *testing.T) {
	post := testCreatePost(t)
	err := TestQueries.DeletePostByPostID(context.Background(), post.ID)
	require.NoError(t, err)
	post1, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post1)
}

func testCreatePost2(t *testing.T, arg CreatePostParams) {
	st := time.Now()
	err := TestQueries.CreatePost(context.Background(), arg)
	testCreatePostNum(t, arg.ID)
	require.NoError(t, err)
	post, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, post.ID, arg.ID)
	require.Equal(t, post.Abstract, arg.Abstract)
	require.Equal(t, post.Cover, arg.Cover)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Content, arg.Content)
	require.False(t, post.Deleted)
	require.Zero(t, post.VisitedNum, post.StarNum)
	require.WithinDuration(t, post.CreateTime, st, time.Second)
	require.WithinDuration(t, post.ModifyTime, st, time.Second)
}
func testCreatePost(t *testing.T) GetPostByPostIDRow {
	arg := CreatePostParams{
		ID:       snowflake.GetID(),
		Cover:    utils.RandomString(10),
		Title:    utils.RandomString(30),
		Abstract: utils.RandomString(30),
		Content:  utils.RandomString(40),
		Public:   true,
	}
	st := time.Now()
	err := TestQueries.CreatePost(context.Background(), arg)
	testCreatePostNum(t, arg.ID)
	require.NoError(t, err)
	post, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.Equal(t, post.ID, arg.ID)
	require.Equal(t, post.Abstract, arg.Abstract)
	require.Equal(t, post.Cover, arg.Cover)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Content, arg.Content)
	require.False(t, post.Deleted)
	require.Zero(t, post.VisitedNum, post.StarNum)
	require.WithinDuration(t, post.CreateTime, st, time.Second)
	require.WithinDuration(t, post.ModifyTime, st, time.Second)
	return post
}

func TestQueries_GetPostInfoByPostID(t *testing.T) {
	post1 := testCreatePost(t)
	post, err := TestQueries.GetPostInfoByPostID(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.ID, post1.ID)
	post, err = TestQueries.GetPostInfoByPostID(context.Background(), 1)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, post)
}

func TestQueries_GetPostByPostID(t *testing.T) {
	post1 := testCreatePost(t)
	post, err := TestQueries.GetPostByPostID(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.ID, post1.ID)
	post, err = TestQueries.GetPostByPostID(context.Background(), 1)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, post)
}

func TestQueries_ListPostBySearchKey(t *testing.T) {
	post := testCreatePost(t)
	rows, err := TestQueries.ListPostBySearchKey(context.Background(), ListPostBySearchKeyParams{
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
			Cover:    utils.RandomOwner(),
			Title:    utils.RandomOwner() + key + utils.RandomOwner(),
			Abstract: utils.RandomOwner(),
			Content:  utils.RandomOwner(),
			Public:   true,
		}
		testCreatePost2(t, arg)
	}
	rows, err = TestQueries.ListPostBySearchKey(context.Background(), ListPostBySearchKeyParams{
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
			Cover:    utils.RandomOwner(),
			Title:    utils.RandomOwner(),
			Abstract: utils.RandomOwner(),
			Content:  utils.RandomOwner(),
			Public:   true,
		}
		testCreatePost2(t, arg)
	}
	time.Sleep(2 * time.Second)
	post, err := TestQueries.ListPostByStartTime(context.Background(), ListPostByStartTimeParams{
		CreateTime:   st,
		CreateTime_2: time.Now(),
		Offset:       0,
		Limit:        1000,
	})
	require.NoError(t, err)
	require.Len(t, post, 10)
}

func TestQueries_ListPostTopping(t *testing.T) {
	sum, err := TestQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	n := 10
	tops := make([]Top, n)
	for i := 0; i < n; i++ {
		tops[i] = testCreateTop(t)
	}
	tops1, err := TestQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, tops1, len(sum)+n)
	for i := range tops {
		testDeleteTopByTopID(t, tops[i].PostID)
	}
	tops1, err = TestQueries.ListPostTopping(context.Background(), ListPostToppingParams{
		Offset: 0,
		Limit:  1000,
	})
	require.Len(t, tops1, len(sum))
}

func TestQueries_ListPostDeleted(t *testing.T) {
	posts1, err := TestQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]GetPostByPostIDRow, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostDeletedByID(t, ModifyPostDeletedByIDParams{
			Deleted: true,
			ID:      posts[i].ID,
		})
	}
	result, err := TestQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
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
	result, err = TestQueries.ListPostDeleted(context.Background(), ListPostDeletedParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func TestQueries_ListPostPrivate(t *testing.T) {
	posts1, err := TestQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]GetPostByPostIDRow, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: false,
			ID:     posts[i].ID,
		})
	}
	result, err := TestQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
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
	result, err = TestQueries.ListPostPrivate(context.Background(), ListPostPrivateParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func TestQueries_ListPostPublic(t *testing.T) {
	posts1, err := TestQueries.ListPostPublic(context.Background(), ListPostPublicParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	n := 10
	posts := make([]GetPostByPostIDRow, n)
	for i := range posts {
		posts[i] = testCreatePost(t)
		testModifyPostPublicByID(t, ModifyPostPublicByIDParams{
			Public: true,
			ID:     posts[i].ID,
		})
	}
	result, err := TestQueries.ListPostPublic(context.Background(), ListPostPublicParams{
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
	result, err = TestQueries.ListPostPublic(context.Background(), ListPostPublicParams{
		Offset: 0,
		Limit:  1000,
	})
	require.NoError(t, err)
	require.Len(t, result, len(posts1))
}

func testModifyPostDeletedByID(t *testing.T, arg ModifyPostDeletedByIDParams) {
	err := TestQueries.ModifyPostDeletedByID(context.Background(), arg)
	require.NoError(t, err)
	post2, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.True(t, post2.Deleted == arg.Deleted)
}

func TestQueries_ModifyPostDeletedByID(t *testing.T) {
	post := testCreatePost(t)
	require.False(t, post.Deleted)
	err := TestQueries.ModifyPostDeletedByID(context.Background(), ModifyPostDeletedByIDParams{
		Deleted: true,
		ID:      post.ID,
	})
	require.NoError(t, err)
	post2, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.True(t, post2.Deleted)
	err = TestQueries.ModifyPostDeletedByID(context.Background(), ModifyPostDeletedByIDParams{
		Deleted: false,
		ID:      post.ID,
	})
	require.NoError(t, err)
	post3, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.False(t, post3.Deleted)
}

func testModifyPostPublicByID(t *testing.T, arg ModifyPostPublicByIDParams) {
	err := TestQueries.ModifyPostPublicByID(context.Background(), arg)
	require.NoError(t, err)
	post2, err := testGetPostByID(arg.ID)
	require.NoError(t, err)
	require.True(t, post2.Public == arg.Public)
}

func TestQueries_ModifyPostPublicByID(t *testing.T) {
	post := testCreatePost(t)
	require.True(t, post.Public)
	err := TestQueries.ModifyPostPublicByID(context.Background(), ModifyPostPublicByIDParams{
		Public: false,
		ID:     post.ID,
	})
	require.NoError(t, err)
	post2, err := testGetPostByID(post.ID)
	require.NoError(t, err)
	require.False(t, post2.Public)
	err = TestQueries.ModifyPostPublicByID(context.Background(), ModifyPostPublicByIDParams{
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
		Cover:    utils.RandomOwner(),
		Title:    utils.RandomOwner(),
		Abstract: utils.RandomOwner(),
		Content:  utils.RandomOwner(),
		Public:   true,
		ID:       post.ID,
	}
	st := time.Now()
	err := TestQueries.UpdatePostByPostID(context.Background(), arg)
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

func TestListPostOrderByCreatedTime(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		testCreatePost(t)
		time.Sleep(time.Second)
	}
	posts, err := TestQueries.ListPostOrderByCreatedTime(context.Background(), ListPostOrderByCreatedTimeParams{
		Offset: 0,
		Limit:  int32(n),
	})
	require.NoError(t, err)
	require.Len(t, posts, n)
	for i := 0; i < n-1; i++ {
		//log.Println(posts[i].CreateTime, posts[i+1].CreateTime)
		require.True(t, posts[i].CreateTime.After(posts[i+1].CreateTime))
	}
}
