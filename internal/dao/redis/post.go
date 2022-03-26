package redis

import (
	"context"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/conversion"
)

func (q *Queries) SetPostInfo(ctx context.Context, postID int64, postInfo string) error {
	return q.rdb.Set(ctx, getRedisKey(keyPostInfo)+conversion.Int64toA(postID), postInfo, global.AllSetting.Redis.PostInfoTimeout).Err()
}

func (q *Queries) GetPostInfo(ctx context.Context, postID int64) (string, error) {
	key := getRedisKey(keyPostInfo) + conversion.Int64toA(postID)
	return q.rdb.Get(ctx, key).Result()
}

func (q *Queries) DeletePostInfo(ctx context.Context, postID int64) error {
	return q.rdb.Expire(ctx, getRedisKey(keyPostInfo)+conversion.Int64toA(postID), 0).Err()
}

func (q *Queries) SetPost(ctx context.Context, postID int64, post string) error {
	return q.rdb.Set(ctx, getRedisKey(keyPost)+conversion.Int64toA(postID), post, global.AllSetting.Redis.PostTimeout).Err()
}

func (q *Queries) GetPost(ctx context.Context, postID int64) (string, error) {
	key := getRedisKey(keyPost) + conversion.Int64toA(postID)
	return q.rdb.Get(ctx, key).Result()
}

func (q *Queries) DeletePost(ctx context.Context, postID int64) error {
	return q.rdb.Expire(ctx, getRedisKey(keyPost)+conversion.Int64toA(postID), 0).Err()
}

func (q *Queries) DeleteCache(ctx context.Context, postID int64) (err error) {
	//删除之间存留的缓存
	if err = q.DeletePost(ctx, postID); err != nil {
		return err
	}
	if err = q.DeletePostInfo(ctx, postID); err != nil {
		return err
	}
	return nil
}
