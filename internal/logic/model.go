package logic

import (
	"mime/multipart"
	"time"
)

type Pagination struct {
	Page     int32 `json:"page" binding:"gte=1" form:"page"`           //页数
	PageSize int32 `json:"page_size" binding:"gte=1" form:"page_size"` //一页的大小
}

type Pager struct {
	Page      int `json:"page,omitempty"` //页数
	PageSize  int `json:"page_size"`      //一页的大小
	TotalRows int `json:"total_rows"`     //结果的总数量
}

type AddCommentParams struct {
	PostID      int64  `json:"post_id" binding:"required,gte=1" form:"post_id"`
	Content     string `json:"content" bind:"required,gte=1,lte=100" form:"content"`
	ToCommentID int64  `json:"to_comment_id" bind:"required,gte=0" form:"to_comment_id"`
}

type ModifyCommentParams struct {
	Content   string `json:"content" binding:"required,gte=1,lte=100" form:"content"`
	CommentID int64  `json:"comment_id" binding:"required,gte=1" form:"comment_id"`
}

type ListCommentByPostIDParams struct {
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	Pagination
}

type Comment struct {
	ID            int64     `json:"id"`             //评论ID
	PostID        int64     `json:"post_id"`        //帖子ID
	Username      string    `json:"username"`       //用户名
	Content       string    `json:"content"`        //内容
	ToCommentID   int64     `json:"to_comment_id"`  //回复的评论的ID
	CreateTime    time.Time `json:"create_time"`    //创建时间
	ModifyTime    time.Time `json:"modify_time"`    //修改时间
	AvatarUrl     string    `json:"avatar_url"`     //头像链接
	DepositoryUrl string    `json:"depository_url"` // 仓库连接
}

type ListCommentByPostIDReply struct {
	List  []Comment `json:"list"`
	Pager Pager     `json:"pager"`
}

type UserInfo struct {
	Username      string `json:"username" `      //"用户名"`
	AvatarUrl     string `json:"avatar_url" `    //"头像链接"`
	DepositoryUrl string `json:"depository_url"` //"仓库链接"`
}

type Token struct {
	Token     string    `json:"token"`     // "token"`
	ExpiredAt time.Time `json:"expiredAt"` // "过期时间"`
}

type ReToken struct {
	RefreshToken string    `json:"refreshToken"` //"刷新token"`
	ExpiredAt    time.Time `json:"expiredAt"`    //"过期时间"`
}

type LoginManagerParams struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

type LoginManagerReply struct {
	Manager `json:"manager"` //用户信息
	Token   `json:"token"`   //Token
	ReToken `json:"reToken"` //ReToken
}

type AddManagerRequest struct {
	Username  string `json:"username" binding:"required,gte=3,lte=50"`
	Password  string `json:"password" binding:"required,gte=3,lte=32"`
	AvatarUrl string `json:"avatar_url" binding:"required,url"`
}

type UpdateManagerRequest struct {
	Password  string `json:"password"`
	AvatarUrl string `json:"avatar_url"`
}

type Manager struct {
	Username  string `json:"username"`   //用户名
	AvatarUrl string `json:"avatar_url"` //头像链接
}

type ListManagerReply struct {
	List  []Manager `json:"list,omitempty"`
	Pager Pager     `json:"pager,omitempty"`
}

type Tag struct {
	ID         int64     `json:"id"`          //标签ID
	TagName    string    `json:"tag_name"`    //标签名
	CreateTime time.Time `json:"create_time"` //创建时间
}

type Post struct {
	ID         int64     `json:"id"`          //帖子ID
	Cover      string    `json:"cover"`       //帖子封面ID
	Title      string    `json:"title"`       //标题
	Abstract   string    `json:"abstract"`    //简介
	Content    string    `json:"content"`     //内容
	Public     bool      `json:"public"`      //是否公开
	Deleted    bool      `json:"deleted"`     //是否删除
	CreateTime time.Time `json:"create_time"` //创建时间
	ModifyTime time.Time `json:"modify_time"` //修改时间
	StarNum    int64     `json:"star_num"`    //点赞数
	VisitedNum int64     `json:"visited_num"` //访问数
}

type PostInfo struct {
	ID         int64     `json:"id"`          //帖子ID
	Cover      string    `json:"cover"`       //帖子封面ID
	Title      string    `json:"title"`       //标题
	Abstract   string    `json:"abstract"`    //简介
	Public     bool      `json:"public"`      //是否公开
	Deleted    bool      `json:"deleted"`     //是否删除
	CreateTime time.Time `json:"create_time"` //创建时间
	ModifyTime time.Time `json:"modify_time"` //修改时间
	StarNum    int64     `json:"star_num"`    //点赞数
	VisitedNum int64     `json:"visited_num"` //访问数
}

type PostParams struct {
	Cover    string `json:"cover" binding:"required,gte=1"`
	Title    string `json:"title" binding:"required,gte=1,lte=50"`
	Abstract string `json:"abstract" binding:"required,gte=1,lte=100"`
	Content  string `json:"content" binding:"required,gte=1"`
	Public   bool   `json:"public" binding:""`
}

type UpdatePostParams struct {
	PostParams
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
}
type ListPostInfosParams struct {
	ListBy string `form:"list_by" binding:"required"`
	Pagination
}
type ListPostInfosReply struct {
	List  []PostInfo `json:"list,omitempty"`
	Pager Pager      `json:"pager"`
}

type ModifyPostDeletedParam struct {
	PostID  int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	Deleted bool  `json:"deleted" binding:"" form:"deleted"`
}

type ModifyPostPublicParam struct {
	PostID int64 `json:"post_id" binding:"required,gte=1"`
	Public bool  `json:"public" binding:""`
}

type SearchPostInfosByKeyParam struct {
	Pagination
	Key string `json:"key" binding:"required,gte=1,lte=15,alphanumunicode" form:"key"`
}

type SearchPostInfosByCreateTimeParam struct {
	Pagination
	StartTime time.Time `json:"start_time" binding:"required,time" form:"start_time"`
	EndTime   time.Time `json:"end_time" binding:"required,time" form:"end_time"`
}

type PostTagParams struct {
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	TagID  int64 `json:"tag_id" binding:"required,gte=1" form:"tag_id"`
}

type DeletePostTagParams struct {
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	TagID  int64 `json:"tag_id" binding:"required,gte=1" form:"tag_id"`
}

type ListTagsByPostIDParams struct {
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	Pagination
}

type ListTagsReply struct {
	List  []Tag `json:"list"`
	Pager Pager `json:"pager"`
}

type ListPostInfosByTagIDParams struct {
	TagID int64 `json:"tag_id" bind:"required,gte=1" form:"tag_id"`
	Pagination
}

type ListPostInfosByTagIDReply struct {
	List  []PostInfo `json:"list"`
	Pager Pager      `json:"pager"`
}

type TokenRedirectParams struct {
	Code string `form:"code" bind:"required,gte=1" form:"code"`
}

type AddTagParams struct {
	TagName string `json:"tag_name" binding:"required,gte=1,lte=10,alphanumunicode" form:"tag_name"`
}

type UpdateTagParams struct {
	TagID   int64  `json:"tag_id" binding:"required,gte=1" form:"tag_id"`
	TagName string `json:"tag_name" binding:"required,gte=1,lte=10,alphanumunicode" form:"tag_name"`
}

type UploadParams struct {
	File     *multipart.FileHeader `json:"file,omitempty" binding:"required" form:"file"`
	FileType string                `json:"file_type,omitempty" binding:"required" form:"file_type"`
}

type GetTokenReply struct {
	User    UserInfo `json:"user"`     //用户信息
	Token   Token    `json:"token"`    //Token
	ReToken ReToken  `json:"re_token"` //ReToken
}

type RefreshTokenReplyParams struct {
	Token   string `json:"token,omitempty" binding:"required" form:"token"`
	ReToken string `json:"re_token,omitempty" binding:"required" form:"re_token"`
}

type RefreshTokenReply struct {
	Token Token `json:"token"` // Token
}

type User struct {
	Username      string    `json:"username"`       //用户名
	AvatarUrl     string    `json:"avatar_url"`     //头像链接
	DepositoryUrl string    `json:"depository_url"` //仓库连接
	Address       string    `json:"address"`        //IP地址
	CreateTime    time.Time `json:"create_time"`    //创建时间
	ModifyTime    time.Time `json:"modify_time"`    //修改时间
}
type ListUsersByCreateTimeParams struct {
	StartTime time.Time `json:"start_time" binding:"required,time" form:"start_time"`
	EndTime   time.Time `json:"end_time" binding:"required,time" form:"end_time"`
	Pagination
}
type UserStarPostParams struct {
	PostID int64 `json:"post_id" binding:"required,gte=1" form:"post_id"`
	State  bool  `json:"state" binding:"" form:"state"`
}
type ListUsersReply struct {
	Users []User `json:"users"`
	Pager Pager  `json:"pager"`
}

type View struct {
	ID         int64     `json:"id"`          //ID
	ViewsNum   int64     `json:"views_num"`   //访问数
	CreateTime time.Time `json:"create_time"` //截止时间
}

type ListViewsByCreateTimeReply struct {
	List  []View `json:"list"`
	Pager Pager  `json:"pager"`
}
