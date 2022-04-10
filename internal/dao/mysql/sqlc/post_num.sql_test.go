package db

import (
	"context"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func testCreatePostNum(t *testing.T, postID int64) {
	err := TestQueries.CreatePost_Num(context.Background(), postID)
	require.NoError(t, err)
}
func TestQueries_CreatePost_Num(t *testing.T) {

}

func TestQueries_UpdatePost_Num_Star(t *testing.T) {
	post := testCreatePost(t)
	err := TestQueries.UpdatePost_Num_Star(context.Background(), UpdatePost_Num_StarParams{
		StarNum: utils.RandomInt(1, 100),
		PostID:  post.ID,
	})
	require.NoError(t, err)
}

func TestQueries_UpdatePost_Num_Visited(t *testing.T) {
	post := testCreatePost(t)
	err := TestQueries.UpdatePost_Num_Visited(context.Background(), UpdatePost_Num_VisitedParams{
		VisitedNum: utils.RandomInt(1, 100),
		PostID:     post.ID,
	})
	require.NoError(t, err)
}
