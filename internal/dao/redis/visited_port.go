package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

/*
帖子新增访问量:
	增加 postid ++
	获取所有的post_id和访问量 并 清零
	获取某个post_id 对应的访问量
*/

func AddVisitedPostNum(ctx context.Context, postID int64) error {
	return rdb.ZIncrBy(ctx, getRedisKey(keyZSetVisitedNum), 1, int64toA(postID)).Err()
}

// GetVisitedPostNum 获取某个post_id 对应的新增访问量
func GetVisitedPostNum(ctx context.Context, postID int64) (int64, error) {
	result, err := rdb.ZScore(ctx, getRedisKey(keyZSetVisitedNum), int64toA(postID)).Result()
	if err != nil {
		return 0, err
	}
	return int64(result), nil
}

// ListPostIDByVisitedNum 对新增访问数排序返回对应id
func ListPostIDByVisitedNum(ctx context.Context, offset, count int64) ([]string, error) {
	return rdb.ZRevRangeByScore(ctx, getRedisKey(keyZSetVisitedNum), &redis.ZRangeBy{
		Min:    "-1",
		Max:    "+inf",
		Offset: offset,
		Count:  count,
	}).Result()
}

// ListAllPostIDByVisitedNumAndSetZero 获取所有的post_id和对应访问量并清零
func ListAllPostIDByVisitedNumAndSetZero(ctx context.Context) (ret map[int64]int64, err error) {
	pipe := rdb.TxPipeline()
	nums := pipe.ZRevRangeByScoreWithScores(ctx, getRedisKey(keyZSetVisitedNum), &redis.ZRangeBy{
		Min: "-1",
		Max: "+inf",
	})
	pipe.Expire(ctx, getRedisKey(keyZSetVisitedNum), 0)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	ret = make(map[int64]int64)
	for _, v := range nums.Val() {
		ret[v.Member.(int64)] = int64(v.Score)
	}
	return ret, nil
}
