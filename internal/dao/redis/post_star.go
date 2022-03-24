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
type BoolString string

const (
	True  BoolString = "true"
	False BoolString = "false"
)

func boolStringToBool(b string) bool {
	return BoolString(b) == True
}

func composePostIDAndUserName(postID int64, userName string) string {
	return int64toA(postID) + ":" + userName
}

// SetPostStar 设置postID和userName的点赞情况
func SetPostStar(ctx context.Context, postID int64, username string, b BoolString) error {
	return rdb.HSet(ctx, getRedisKey(KeyHMapPostStar), []string{composePostIDAndUserName(postID, username), string(b)}).Err()
}

// GetPostStarByPostIDAndUserName 通过postID和userName获取对应帖子点赞情况
func GetPostStarByPostIDAndUserName(ctx context.Context, postID int64, username string) (bool, error) {
	result, err := rdb.HGet(ctx, getRedisKey(KeyHMapPostStar), composePostIDAndUserName(postID, username)).Result()
	if err != nil {
		return false, err
	}
	return boolStringToBool(result), nil
}

// ListAllPostStarAndSetZero 列出所有帖子和对应postID的点赞关系，然后清空
func ListAllPostStarAndSetZero(ctx context.Context) (ret map[int64]map[string]bool, err error) {
	pip := rdb.TxPipeline()
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
		ret[postID][userName] = boolStringToBool(b)
	}
	return ret, nil
}
