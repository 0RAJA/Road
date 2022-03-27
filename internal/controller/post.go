package controller

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/0RAJA/Road/internal/pkg/conversion"
	"github.com/gin-gonic/gin"
)

func checkPost(title, abstract string) *errcode.Error {
	if len(title) <= 0 || len(title) > global.AllSetting.Rule.TitleLen {
		return errcode.InvalidParamsErr
	}
	if len(abstract) <= 0 || len(abstract) > global.AllSetting.Rule.AbstractLen {
		return errcode.InvalidParamsErr
	}
	return nil
}

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
	params := logic.PostParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := checkPost(params.Title, params.Abstract); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := logic.AddPost(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
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
	params := logic.UpdatePostParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(errs.Errors()...))
		return
	}
	if err := checkPost(params.Title, params.Abstract); err != nil {
		response.ToErrorResponse(err)
		return
	}
	if err := logic.UpdatePost(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// GetPost
// @Summary 获取一个帖子的完整信息
// @Description 获取一个帖子的完整信息
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {object} logic.Post "获取一个帖子的ID,封面，标题，简介，内容,是否公开,是否删除以及,创建时间,修改时间,点赞数和浏览数"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/post/{post_id} [get]
func GetPost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postID := conversion.AtoInt64Must(app.GetPath(ctx, "post_id"))
	if postID == 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	post, err := logic.GetPost(ctx, postID)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(post)
}

// GetPostInfo
// @Summary 获取一个帖子的简介信息
// @Description 获取一个帖子的简介信息
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {object}  logic.PostInfo "返回一个帖子的ID,封面，标题，简介，是否公开,是否删除以及,创建时间和修改时间以及点赞数和访问数和其对应标签的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/info/{post_id} [get]
func GetPostInfo(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postID := conversion.AtoInt64Must(app.GetPath(ctx, "post_id"))
	if postID == 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	postInfo, err := logic.GetPostInfo(ctx, postID)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(postInfo)
}

// ModifyPostDeleted
// @Summary 修改一个帖子删除状态
// @Description 修改一个帖子删除状态
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param deleted body bool true "帖子删除状态"
// @Success 200 {string} string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/delete [put]
func ModifyPostDeleted(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ModifyPostDeletedParam{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	if err := logic.ModifyPostDeleted(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// RealDeletePost
// @Summary 真正删除一个帖子
// @Description 将一个处于删除状态的帖子真正删除
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id path int64 true "帖子ID"
// @Success 200 {string}  string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/{post_id} [delete]
func RealDeletePost(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	postID := conversion.AtoInt64Must(app.GetPath(ctx, "post_id"))
	if postID <= 0 {
		response.ToErrorResponse(errcode.InvalidParamsErr)
		return
	}
	if err := logic.RealDeletePost(ctx, postID); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// ModifyPostPublic
// @Summary 修改一个帖子公开状态
// @Description 修改一个帖子公开状态
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param post_id body int64 true "帖子ID"
// @Param public body bool true "帖子公开状态"
// @Success 200 {string}  string ""
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/public [put]
func ModifyPostPublic(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ModifyPostPublicParam{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	if err := logic.ModifyPostPublic(ctx, params); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// ListPostInfos
// @Summary 列出帖子简介
// @Description 列出帖子简介，默认按创建时间倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param list_by query string true "列出什么类型的帖子 Enums(infos,public,private,deleted,topping,star_num,visited_num) 分别对应默认时间排序，公开的，私密的，删除的，置顶的，点赞数排序，访问数排序"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos [get]
func ListPostInfos(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.ListPostInfosParams{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPage(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.ListPostInfos(ctx, params.ListBy, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// SearchPostInfosByKey
// @Summary 通过关键字搜索帖子的标题和简介
// @Description 通过关键字搜索帖子的标题和简介，默认按置顶先后倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param key query string key "关键字 1<=len<=15"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/search [get]
func SearchPostInfosByKey(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.SearchPostInfosByKeyParam{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPage(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.SearchPostInfosByKey(ctx, params.Key, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// SearchPostInfosByCreateTime
// @Summary 通过时间段来检索帖子
// @Description 通过时间段来检索帖子，默认按创建时间先后倒序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
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
	params := logic.SearchPostInfosByCreateTimeParam{
		Pagination: logic.Pagination{
			Page:     app.GetPage(ctx),
			PageSize: app.GetPage(ctx),
		},
	}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.SearchPostInfosByCreateTime(ctx, params.StartTime, params.EndTime, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
}

// ListPostInfosOrderByGrowingVisited
// @Summary 通过按新增访问数排序的帖子简介信息
// @Description 通过按新增访问数排序的帖子简介信息，按新增访问数由高到低排序
// @Tags 文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param page query int false "页码 default 1"
// @Param page_size query int false "每页数量 default and max 10"
// @Success 200 {object} logic.ListPostInfosReply "返回帖子简介的数组和描述数组大小的信息"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /post/infos/visit/grow [get]
func ListPostInfosOrderByGrowingVisited(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.Pagination{
		Page:     app.GetPage(ctx),
		PageSize: app.GetPage(ctx),
	}
	reply, err := logic.ListPostInfosOrderByGrowingVisited(ctx, app.GetPageOffset(params.Page, params.PageSize), params.PageSize)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponseList(reply, len(reply))
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
