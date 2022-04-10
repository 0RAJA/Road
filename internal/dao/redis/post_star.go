package redis

import (
	"context"
	"errors"
	"strings"
)

/*
	hmap
	用户点赞关系 key post_id:username true
	设置
	获取所有内容并清空
*/

func stringToBool(b string) bool {
	return b == "true"
}
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
func composePostIDAndUserName(postID int64, userName string) string {
	return int64toA(postID) + ":" + userName
}

// SetPostStar 设置postID和userName的点赞情况
func (q *Queries) SetPostStar(ctx context.Context, postID int64, username string, b bool) error {
	return q.rdb.HSet(ctx, getRedisKey(KeyHMapPostStar), []string{composePostIDAndUserName(postID, username), boolToString(b)}).Err()
}

// GetPostStarByPostIDAndUserName 通过postID和userName获取对应帖子点赞情况
func (q *Queries) GetPostStarByPostIDAndUserName(ctx context.Context, postID int64, username string) (bool, error) {
	result, err := q.rdb.HGet(ctx, getRedisKey(KeyHMapPostStar), composePostIDAndUserName(postID, username)).Result()
	if err != nil {
		return false, err
	}
	return stringToBool(result), nil
}

// ListAllPostStarAndSetZero 列出所有帖子和对应postID的点赞关系，然后清空
func (q *Queries) ListAllPostStarAndSetZero(ctx context.Context) (ret map[int64]map[string]bool, err error) {
	pip := q.rdb.TxPipeline()
	results, _ := pip.HGetAll(ctx, getRedisKey(KeyHMapPostStar)).Result()
	pip.Expire(ctx, getRedisKey(KeyHMapPostStar), 0)
	_, err = pip.Exec(ctx)
	if err != nil {
		return nil, err
	}
	ret = make(map[int64]map[string]bool)
	for postIDUserName, b := range results {
		s := strings.SplitN(postIDUserName, ":", 1)
		postID, userName := atoInt64Must(s[0]), s[1]
		if postID == 0 || userName == "" {
			return nil, errors.New("parsingFailed:" + postIDUserName + ":" + b)
		}
		if ret[postID] == nil {
			ret[postID] = make(map[string]bool)
		}
		ret[postID][userName] = stringToBool(b)
	}
	return ret, nil
}
