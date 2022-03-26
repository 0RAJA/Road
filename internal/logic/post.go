package logic

import (
	"encoding/json"
	"errors"
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/dao/redis"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"time"
)

func AddPost(ctx *gin.Context, request PostRequest) error {
	arg := db.CreatePostParams{
		ID:       snowflake.GetID(),
		Cover:    request.Cover,
		Title:    request.Title,
		Abstract: request.Abstract,
		Content:  request.Content,
		Public:   request.Public,
	}
	err := mysql.Query.CreatePost(ctx, arg)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	err = mysql.Query.CreatePost_Num(ctx, arg.ID)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	return nil
}

func UpdatePost(ctx *gin.Context, params UpdatePostParams) error {
	_, err := mysql.Query.GetPostByPostID(ctx, params.PostID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	err = mysql.Query.UpdatePostByPostID(ctx, db.UpdatePostByPostIDParams{
		Cover:    params.Cover,
		Title:    params.Title,
		Abstract: params.Abstract,
		Content:  params.Content,
		Public:   params.Public,
		ID:       params.PostID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	//删除之间存留的缓存
	if err = redis.Query.DeleteCache(ctx, params.PostID); err != nil {
		global.Logger.Error(err.Error())
	}
	return nil
}

func GetPost(ctx *gin.Context, postID int64) (Post, error) {
	isRoot := getRoot(ctx)
	post, err := getPost(ctx, postID)
	if err != nil {
		return Post{}, err
	}
	if !post.Public && !isRoot {
		return Post{}, errcode.InsufficientPermissionsErr
	}
	return post, nil
}

//同一时间内只会有一次请求
func getPost(ctx *gin.Context, postID int64) (Post, error) {
	result, err := doOnce.Do(getPostKey(postID), func() (interface{}, error) {
		var (
			ret Post
			err error
		)
		//尝试获取缓存
		cache, err := redis.Query.GetPost(ctx, postID)
		if err == nil {
			err := json.Unmarshal([]byte(cache), &ret)
			if err == nil {
				return ret, nil
			}
			global.Logger.Error(err.Error())
		} else if !redis.IsNil(err) {
			global.Logger.Error(err.Error())
		}
		post, err := mysql.Query.GetPostByPostID(ctx, postID)
		if err != nil {
			if mysql.IsNil(err) {
				return Post{}, errcode.ErrPostNotFind
			}
			global.Logger.Error(err.Error())
			return Post{}, errcode.ServerErr
		}
		ret = Post(post)
		//设置缓存
		cache1, err := json.Marshal(ret)
		if err != nil {
			global.Logger.Error(err.Error())
			return ret, nil
		}
		err = redis.Query.SetPost(ctx, postID, string(cache1))
		if err != nil {
			global.Logger.Error(err.Error())
		}
		return ret, nil
	})
	if err != nil {
		return Post{}, err
	}
	return result.(Post), nil
}

//同一时间内只会有一次请求
func getPostInfo(ctx *gin.Context, postID int64) (PostInfo, error) {
	result, err := doOnce.Do(getPostInfoKey(postID), func() (interface{}, error) {
		var (
			ret PostInfo
			err error
		)
		//尝试获取缓存
		cache, err := redis.Query.GetPostInfo(ctx, postID)
		if err == nil {
			err := json.Unmarshal([]byte(cache), &ret)
			if err == nil {
				return ret, nil
			}
			global.Logger.Error(err.Error())
		} else if !redis.IsNil(err) {
			global.Logger.Error(err.Error())
		}
		//获取数据库
		postInfo, err := mysql.Query.GetPostInfoByPostID(ctx, postID)
		if err != nil {
			if mysql.IsNil(err) {
				return PostInfo{}, errcode.ErrPostNotFind
			}
			global.Logger.Error(err.Error())
			return PostInfo{}, errcode.ServerErr
		}
		ret = PostInfo(postInfo)
		//加入缓存
		cache1, err := json.Marshal(ret)
		if err != nil {
			global.Logger.Error(err.Error())
			return ret, nil
		}
		err = redis.Query.SetPostInfo(ctx, postID, string(cache1))
		if err != nil {
			global.Logger.Error(err.Error())
		}
		return ret, nil
	})
	if err != nil {
		return PostInfo{}, err
	}
	return result.(PostInfo), nil
}

func GetPostInfo(ctx *gin.Context, postID int64) (PostInfo, error) {
	isRoot := getRoot(ctx)
	postInfo, err := getPostInfo(ctx, postID)
	if err != nil {
		return postInfo, err
	}
	if !postInfo.Public && !isRoot {
		return PostInfo{}, errcode.InsufficientPermissionsErr
	}
	return postInfo, nil
}

func ModifyPostDeleted(ctx *gin.Context, params ModifyPostDeletedParam) error {
	postInfo, err := mysql.Query.GetPostInfoByPostID(ctx, params.PostID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if postInfo.Deleted == params.Deleted {
		return errcode.ErrStateRepeat
	}
	err = mysql.Query.ModifyPostDeletedByID(ctx, db.ModifyPostDeletedByIDParams{
		Deleted: params.Deleted,
		ID:      params.PostID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	//删除之间存留的缓存
	if err = redis.Query.DeleteCache(ctx, params.PostID); err != nil {
		global.Logger.Error(err.Error())
	}
	return nil
}

func RealDeletePost(ctx *gin.Context, postID int64) error {
	postInfo, err := mysql.Query.GetPostInfoByPostID(ctx, postID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if !postInfo.Deleted {
		return errcode.ErrDeletedState
	}
	err = mysql.Query.DeletePostByPostID(ctx, postID)
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	//删除之间存留的缓存
	if err = redis.Query.DeleteCache(ctx, postID); err != nil {
		global.Logger.Error(err.Error())
	}
	return nil
}

func ModifyPostPublic(ctx *gin.Context, params ModifyPostPublicParam) error {
	postInfo, err := mysql.Query.GetPostInfoByPostID(ctx, params.PostID)
	if err != nil {
		if mysql.IsNil(err) {
			return errcode.ErrPostNotFind
		}
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	if postInfo.Public == params.Public {
		return errcode.ErrStateRepeat
	}
	err = mysql.Query.ModifyPostPublicByID(ctx, db.ModifyPostPublicByIDParams{
		Public: params.Public,
		ID:     params.PostID,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return errcode.ServerErr
	}
	//删除之间存留的缓存
	if err = redis.Query.DeleteCache(ctx, params.PostID); err != nil {
		global.Logger.Error(err.Error())
	}
	return nil
}

// ListPostInfos options: Enums(infos,public,private,deleted,topping,star_num,visited_num)
func ListPostInfos(ctx *gin.Context, options string, offset, limit int32) ([]PostInfo, error) {
	var (
		postInfos []PostInfo
		err       error
	)
	switch options {
	case "infos":
		postInfos, err = listPostInfos(ctx, offset, limit)
	case "public":
		postInfos, err = listPostInfosByPublic(ctx, offset, limit)
	case "private":
		postInfos, err = listPostInfosByPrivate(ctx, offset, limit)
	case "deleted":
		postInfos, err = listPostInfosByDeleted(ctx, offset, limit)
	case "topping":
		postInfos, err = listPostInfosByTopping(ctx, offset, limit)
	case "star_num":
		postInfos, err = listPostInfosByStarNum(ctx, offset, limit)
	case "visited_num":
		postInfos, err = listPostInfosByVisitedNum(ctx, offset, limit)
	default:
		return nil, errcode.ErrListPostInfosOptions
	}
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return postInfos, nil
}

func listPostInfos(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostOrderByCreatedTime(ctx, db.ListPostOrderByCreatedTimeParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func listPostInfosByPublic(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostPublic(ctx, db.ListPostPublicParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func listPostInfosByPrivate(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostPrivate(ctx, db.ListPostPrivateParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}
func listPostInfosByDeleted(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostDeleted(ctx, db.ListPostDeletedParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}
func listPostInfosByTopping(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostTopping(ctx, db.ListPostToppingParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func listPostInfosByStarNum(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostOrderByStarNum(ctx, db.ListPostOrderByStarNumParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func listPostInfosByVisitedNum(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostOrderByVisitedNum(ctx, db.ListPostOrderByVisitedNumParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func SearchPostInfosByKey(ctx *gin.Context, key string, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostBySearchKey(ctx, db.ListPostBySearchKeyParams{
		Title:    key,
		Abstract: key,
		Offset:   offset,
		Limit:    limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func SearchPostInfosByCreateTime(ctx *gin.Context, startTime, endTime time.Time, offset, limit int32) ([]PostInfo, error) {
	isRoot := getRoot(ctx)
	results, err := mysql.Query.ListPostByStartTime(ctx, db.ListPostByStartTimeParams{
		CreateTime:   startTime,
		CreateTime_2: endTime,
		Offset:       offset,
		Limit:        limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	ret := make([]PostInfo, 0, len(results))
	for i := range results {
		if results[i].Public || isRoot {
			ret = append(ret, PostInfo(results[i]))
		}
	}
	return ret, nil
}

func ListPostInfosOrderByGrowingVisited(ctx *gin.Context, offset, limit int32) ([]PostInfo, error) {
	ids, err := redis.Query.ListPostIDByVisitedNum(ctx, offset, limit)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	postInfos := make([]PostInfo, 0, len(ids))
	for i := range postInfos {
		info, err := GetPostInfo(ctx, ids[i])
		if err != nil {
			if mysql.IsNil(err) || errors.Is(err, errcode.InsufficientPermissionsErr) {
				continue
			}
			global.Logger.Error(err.Error())
			return nil, errcode.ServerErr
		}
		postInfos = append(postInfos, info)
	}
	return postInfos, nil
}
