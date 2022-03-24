package redis

import (
	"context"
)

// AddVisitedUserName 增加每日访问人数
func AddVisitedUserName(ctx context.Context, username string) error {
	key := getRedisKey(KeyHyperLongLongVisitedNum)
	return rdb.Do(ctx, "pfadd", key, username).Err()
}

// CountVisitedNumsAndSetZero 获取并清空每日访问人数
func CountVisitedNumsAndSetZero(ctx context.Context) (int64, error) {
	pipe := rdb.TxPipeline()
	key := getRedisKey(KeyHyperLongLongVisitedNum)
	v := pipe.Do(ctx, "pfcount", key)
	pipe.Expire(ctx, key, 0)
	_, err := pipe.Exec(ctx)
	return v.Val().(int64), err
}
