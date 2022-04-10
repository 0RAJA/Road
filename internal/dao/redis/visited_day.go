package redis

import (
	"context"
)

// AddVisitedUserName 增加每日访问人数
func (q *Queries) AddVisitedUserName(ctx context.Context, username string) error {
	key := getRedisKey(KeyHyperLongLongVisitedNum)
	return q.rdb.Do(ctx, "pfadd", key, username).Err()
}

// CountVisitedNumsAndSetZero 获取并清空访问人数
func (q *Queries) CountVisitedNumsAndSetZero(ctx context.Context) (int64, error) {
	pipe := q.rdb.TxPipeline()
	key := getRedisKey(KeyHyperLongLongVisitedNum)
	v := pipe.Do(ctx, "pfcount", key)
	pipe.Expire(ctx, key, 0)
	_, err := pipe.Exec(ctx)
	return v.Val().(int64), err
}
