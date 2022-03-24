package redis

import (
	"context"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/conversion"
)

func SetPostInfo(ctx context.Context, postID int64, postInfo string) error {
	return rdb.Set(ctx, getRedisKey(keyPostInfo)+conversion.Int64toA(postID), postInfo, global.AllSetting.Redis.PostInfoTimeout).Err()
}

func GetPostInfo(ctx context.Context, postID int64) (string, error) {
	return rdb.Get(ctx, getRedisKey(keyPostInfo)+conversion.Int64toA(postID)).Result()
}

func SetPost(ctx context.Context, postID int64, post string) error {
	return rdb.Set(ctx, getRedisKey(keyPost)+conversion.Int64toA(postID), post, global.AllSetting.Redis.PostTimeout).Err()
}

func GetPost(ctx context.Context, postID int64) (string, error) {
	return rdb.Get(ctx, getRedisKey(keyPost)+conversion.Int64toA(postID)).Result()
}
