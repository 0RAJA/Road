package logic

import (
	"context"
	"github.com/0RAJA/Road/internal/dao/mysql"
	db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/dao/redis"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"time"
)

func EnduranceViews(ctx context.Context) {
	sum, err := redis.Query.CountVisitedNumsAndSetZero(ctx)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	if err = mysql.Query.CreateView(ctx, sum); err != nil {
		global.Logger.Error(err.Error())
	}
}

func ListViewsByCreateTime(ctx *gin.Context, startTime, endTime time.Time, offset, limit int32) ([]db.View, *errcode.Error) {
	views, err := mysql.Query.ListViewByCreateTime(ctx, db.ListViewByCreateTimeParams{
		CreateTime:   startTime,
		CreateTime_2: endTime,
		Offset:       offset,
		Limit:        limit,
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, errcode.ServerErr
	}
	return views, nil
}

func AddPostViews(ctx *gin.Context, postID int64) {
	if err := redis.Query.AddVisitedPostNum(ctx, postID); err != nil {
		global.Logger.Error(err.Error())
	}
}

func GetGrowViewsByPostID(ctx *gin.Context, postID int64) (int64, *errcode.Error) {
	view, err := redis.Query.GetVisitedPostNum(ctx, postID)
	if err != nil {
		return 0, errcode.ServerErr
	}
	return view, nil
}

func EndurancePostGrowViews(ctx context.Context) {
	results, err := redis.Query.ListAllPostIDByVisitedNumAndSetZero(ctx)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	for postID := range results {
		err := mysql.Query.UpdatePost_Num_Visited(ctx, db.UpdatePost_Num_VisitedParams{
			VisitedNum: results[postID],
			PostID:     postID,
		})
		if err != nil {
			global.Logger.Error(err.Error())
		}
	}
}
