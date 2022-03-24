package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// AddPost
// @Summary 新增帖子
// @Description 新增一个帖子的封面，标题，简介，内容以及确定其是否公开
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param cover body string true "封面链接"
// @Param title body string true "标题 1<=len<=50"
// @Param abstract body string true "简介 1<=len<=100"
// @Param content body string true "内容"
// @Param public body bool true "是否公开 Enums[true,false]"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/create [post]
func AddPost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// UpdatePost
// @Summary 更新帖子
// @Description 更新一个帖子的封面，标题，简介，内容以及确定其是否公开
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param cover body string true "封面链接"
// @Param title body string true "标题"
// @Param abstract body string true "简介"
// @Param content body string true "内容"
// @Param public body bool true "是否公开"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/update [put]
func UpdatePost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// GetPost
// @Summary 获取一个帖子的完整信息
// @Description 获取一个帖子的完整信息
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param post_id path int64 true "帖子ID"
// @Success 200 {object} logic.PostWithTags "获取一个帖子的ID,封面，标题，简介，内容,是否公开,是否删除以及,创建时间,修改时间,点赞数和浏览数"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/post/{post_id} [get]
func GetPost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	post := logic.Post{
		ID:         utils.RandomInt(1, 100),
		Cover:      "",
		Title:      "",
		Abstract:   "",
		Content:    "",
		Public:     false,
		Deleted:    false,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
		StarNum:    0,
		VisitedNum: 0,
	}
	response.ToResponse(post)
}

// GetPostInfo
// @Summary 获取一个帖子的简介信息
// @Description 获取一个帖子的简介信息
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param post_id path int64 true "帖子ID"
// @Success 200 {object}  logic.PostInfoWithTags "返回一个帖子的ID,封面，标题，简介，是否公开,是否删除以及,创建时间和修改时间以及点赞数和访问数和其对应标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/info/{post_id} [get]
func GetPostInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postInfo := logic.PostInfoWithTags{
		PostInfo: logic.PostInfo{
			ID:         utils.RandomInt(1, 100),
			Cover:      "",
			Title:      "",
			Abstract:   "",
			Public:     false,
			Deleted:    false,
			CreateTime: time.Now(),
			ModifyTime: time.Now(),
			StarNum:    0,
			VisitedNum: 0,
		},
	}
	response.ToResponse(postInfo)
}

// ModifyPostDeleted
// @Summary 修改一个帖子删除状态
// @Description 修改一个帖子删除状态
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param post_id body int64 true "帖子ID"
// @Param deleted body bool true "帖子删除状态"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/delete [put]
func ModifyPostDeleted(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// RealDeletePost
// @Summary 真正删除一个帖子
// @Description 将一个处于删除状态的帖子真正删除
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param post_id path int64 true "帖子ID"
// @Success 200 {string}  string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/{post_id} [delete]
func RealDeletePost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ModifyPostPublic
// @Summary 修改一个帖子公开状态
// @Description 修改一个帖子公开状态
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param post_id body int64 true "帖子ID"
// @Param public body bool true "帖子公开状态"
// @Success 200 {string}  string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/public [put]
func ModifyPostPublic(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.ToResponse(nil)
}

// ListPostInfos
// @Summary 列出帖子简介
// @Description 列出帖子简介，默认按创建时间倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos [get]
func ListPostInfos(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosPublic
// @Summary 列出公开的帖子的简介信息
// @Description 列出公开的帖子简介信息，默认按创建时间倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/public [get]
func ListPostInfosPublic(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosPrivate
// @Summary 列出私密的帖子的简介信息
// @Description 列出私密的帖子简介信息，默认按创建时间倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/private [get]
func ListPostInfosPrivate(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosDeleted
// @Summary 列出标记为删除的帖子的简介信息
// @Description 列出标记为删除的帖子的简介信息，默认按创建时间倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/deleted [get]
func ListPostInfosDeleted(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosTopping
// @Summary 列出置顶的帖子的简介信息
// @Description 列出置顶的帖子的简介信息，默认按置顶先后倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/topping [get]
func ListPostInfosTopping(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// SearchPostInfosByKey
// @Summary 通过关键字搜索帖子的标题和简介
// @Description 通过关键字搜索帖子的标题和简介，默认按置顶先后倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param key query string key "关键字 1<=len<=15"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/search [get]
func SearchPostInfosByKey(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// SearchPostInfosByCreateTime
// @Summary 通过时间段来检索帖子
// @Description 通过时间段来检索帖子，默认按创建时间先后倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param start_time query time.Time true "起始时间 (2002-03-26)"
// @Param end_time query time.Time true "结束时间 (2002-03-26)"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/time [get]
func SearchPostInfosByCreateTime(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosOrderByStarNum
// @Summary 通过点赞数排序获取帖子简介信息
// @Description 通过点赞数排序获取帖子简介信息，按点赞数由高到低排序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/star [get]
func ListPostInfosOrderByStarNum(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosOrderByVisitedNum
// @Summary 通过访问数排序获取帖子简介信息
// @Description 通过访问数排序获取帖子简介信息，按访问数由高到低排序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/visit [get]
func ListPostInfosOrderByVisitedNum(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosOrderByGrowingStar
// @Summary 通过按新增点赞数排序的帖子简介信息
// @Description 通过按新增点赞数排序的帖子简介信息，按新增点赞数由高到低排序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/star/grow [get]
func ListPostInfosOrderByGrowingStar(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

// ListPostInfosOrderByGrowingVisited
// @Summary 通过按新增访问数排序的帖子简介信息
// @Description 通过按新增访问数排序的帖子简介信息，按新增访问数由高到低排序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/visit/grow [get]
func ListPostInfosOrderByGrowingVisited(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	pageSize := app.GetPageSize(ctx)
	infos := make([]logic.PostInfoWithTags, pageSize)
	for i := range infos {
		infos[i] = logic.PostInfoWithTags{
			PostInfo: logic.PostInfo{
				ID:         utils.RandomInt(1, 100),
				Cover:      "",
				Title:      "",
				Abstract:   "",
				Public:     false,
				Deleted:    false,
				CreateTime: time.Now(),
				ModifyTime: time.Now(),
				StarNum:    0,
				VisitedNum: 0,
			},
			Tags: nil,
		}
	}
	response.ToResponseList(infos, len(infos))
}

/*
帖子
	创建
	更新
	获取单个完整内容
	获取单个简介内容
	设置是否移入回收站
	从回收站删除
	设置是否公开
	给帖子点赞
	获取点赞信息
	列出帖子简介
	列出公开的
	列出私密的
	列出回收站的
	列出置顶的
	通过关键字查询标题和简介
	查询指定创建时间内的帖子
	按帖子最近点赞量排序
	按帖子最近访问量排序
*/
