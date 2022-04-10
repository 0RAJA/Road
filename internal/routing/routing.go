package routing

import (
	"context"
	"github.com/0RAJA/Road/internal/controller"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/logic"
	mid "github.com/0RAJA/Road/internal/middleware"
	"github.com/0RAJA/Road/internal/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"time"
)

var swagHandler gin.HandlerFunc

var limit = limiter.NewPrefixLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/post/infos",
	FillInterval: time.Second,
	Cap:          100,
	Quantum:      100,
})

func NewRouting() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors())
	r.Use(mid.GinRecovery(true), mid.GinLogger(), mid.Translations(), mid.Auth(), mid.Limiter(limit))
	//优化
	if swagHandler != nil {
		r.GET("swagger/*doc", swagHandler)
	}
	r.Static("static/", global.AllSetting.Upload.StaticPath) //静态资源位置
	comment := r.Group("comment")
	{
		comment.GET("list", controller.ListComments)
		loginAuth := r.Group("comment").Use(mid.Login())
		loginAuth.POST("add", controller.AddComment)
		loginAuth.DELETE(":comment_id", controller.DeleteComment)
		loginAuth.PUT("update", controller.ModifyComment)
	}
	manager := r.Group("manager")
	{
		manager.POST("login", controller.LoginManager)
		managerAuth := r.Group("manager").Use(mid.Login(), mid.ManagerAuth())
		managerAuth.GET("check/:username", controller.CheckManagerName)
		managerAuth.POST("create", controller.AddManager)
		managerAuth.PUT("update", controller.UpdateManager)
		managerAuth.DELETE(":username", controller.DeleteManager)
		managerAuth.GET("list", controller.ListManagers)
	}
	post := r.Group("post")
	{
		post.GET("post/:post_id", controller.GetPost)
		post.GET("info/:post_id", controller.GetPostInfo)
		infos := post.Group("infos")
		{
			infos.GET("list", controller.ListPostInfos)
			infos.GET("search", controller.SearchPostInfosByKey)
			infos.GET("time", controller.SearchPostInfosByCreateTime)
			infos.GET("visit", controller.ListPostInfosOrderByGrowingVisited)
		}
		managerAuth := r.Group("post").Use(mid.Login(), mid.ManagerAuth())
		managerAuth.PUT("update", controller.UpdatePost)
		managerAuth.PUT("delete", controller.ModifyPostDeleted)
		managerAuth.PUT("public", controller.ModifyPostPublic)
		managerAuth.POST("create", controller.AddPost)
		managerAuth.DELETE(":post_id", controller.RealDeletePost)
	}
	postTag := r.Group("postTag")
	{
		postTag.GET("tags", controller.ListTagsByPostID)
		postTag.GET("infos", controller.ListPostInfosByTagID)
		managerAuth := r.Group("postTag").Use(mid.Login(), mid.ManagerAuth())
		managerAuth.POST("add", controller.AddPostTag)
		managerAuth.DELETE("delete", controller.DeletePostTag)
	}
	tag := r.Group("tag")
	{
		tag.GET("list", controller.ListTags)
		tag.GET("check", controller.CheckTagName)
		managerAuth := r.Group("tag").Use(mid.Login(), mid.ManagerAuth())
		managerAuth.POST("add", controller.AddTag)
		managerAuth.DELETE(":tag_id", controller.DeleteTag)
		managerAuth.PUT("update", controller.UpdateTag)
	}
	token := r.Group("token")
	{
		token.GET("get", controller.GetToken)
		token.GET("redirect", controller.TokenRedirect) //用于从github接受消息用
		token.PUT("refresh", controller.RefreshToken)
	}
	user := r.Group("user", mid.Login(), mid.ManagerAuth())
	{
		user.GET(":username", controller.GetUserInfo)
		user.GET("users", controller.ListUsers)
		user.GET("createTime", controller.ListUsersByCreateTime)
	}
	star := r.Group("star", mid.Login())
	{
		star.GET(":post_id", controller.GetUserStar)
		star.PUT("update", controller.UserStarPost)
	}
	views := r.Group("views")
	{
		views.GET("post/:post_id", controller.GetGrowViewsByPostID)
		managerAuth := r.Group("views").Use(mid.Login(), mid.ManagerAuth())
		managerAuth.GET("all", controller.ListViewsByCreateTime)
	}
	r.POST("upload", mid.Login(), mid.ManagerAuth(), controller.Upload)
	logic.Automation(context.Background()) //自动化
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("time", mid.TimeFormat) //注册自定义标签
	}
	return r
}
