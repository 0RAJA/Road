package routing

import (
	"context"
	_ "github.com/0RAJA/Road/docs"
	"github.com/0RAJA/Road/internal/controller"
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/logic"
	mid "github.com/0RAJA/Road/internal/middleware"
	"github.com/0RAJA/Road/internal/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

var limit = limiter.NewPrefixLimiter().AddBuckets(limiter.BucketRule{
	Key:          "/post/infos",
	FillInterval: time.Second,
	Cap:          100,
	Quantum:      100,
})

func NewRouting() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinRecovery(true), mid.GinLogger(), mid.Translations(), mid.Auth(), mid.Limiter(limit))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/static/", global.AllSetting.Upload.StaticPath) //静态资源位置
	comment := r.Group("/comment")
	{
		loginAuth := comment.Use(mid.Login())
		comment.GET("/", controller.ListComments)
		loginAuth.POST("/", controller.AddComment)
		loginAuth.DELETE("/:comment_id", controller.DeleteComment)
		loginAuth.PUT("/", controller.ModifyComment)
	}
	manager := r.Group("/manager", mid.Login(), mid.ManagerAuth())
	{
		manager.GET("/check/:username", controller.CheckManagerName)
		manager.POST("/login", controller.LoginManager)
		manager.POST("/create", controller.AddManager)
		manager.PUT("/", controller.UpdateManager)
		manager.DELETE("/:username", controller.DeleteManager)
		manager.GET("/", controller.ListManagers)
	}
	post := r.Group("/post")
	{
		post.GET("/post/:post_id", controller.GetPost)
		post.GET("/info/:post_id", controller.GetPostInfo)
		infos := post.Group("/infos")
		{
			infos.GET("/", controller.ListPostInfos)
			infos.GET("/search", controller.SearchPostInfosByKey)
			infos.GET("/time", controller.SearchPostInfosByCreateTime)
			infos.GET("/visit/grow", controller.ListPostInfosOrderByGrowingVisited)
		}
		managerAuth := post.Use(mid.Login(), mid.ManagerAuth())
		managerAuth.PUT("/update", controller.UpdatePost)
		managerAuth.PUT("/delete", controller.ModifyPostDeleted)
		managerAuth.PUT("/public", controller.ModifyPostPublic)
		managerAuth.POST("/create", controller.AddPost)
		managerAuth.DELETE("/:post_id", controller.RealDeletePost)
	}
	postTag := r.Group("/postTag")
	{
		managerAuth := postTag.Use(mid.Login(), mid.ManagerAuth())
		managerAuth.POST("/", controller.AddPostTag)
		managerAuth.DELETE("/", controller.DeletePostTag)
		postTag.GET("/tags", controller.ListTagsByPostID)
		postTag.GET("/infos", controller.ListPostInfosByTagID)
	}
	tag := r.Group("/tag")
	{
		managerAuth := tag.Use(mid.Login(), mid.ManagerAuth())
		managerAuth.POST("/", controller.AddTag)
		managerAuth.DELETE("/:tag_id", controller.DeleteTag)
		managerAuth.PUT("/", controller.UpdateTag)
		tag.GET("/", controller.ListTags)
		tag.GET("/check", controller.CheckTagName)
	}
	token := r.Group("/token")
	{
		token.GET("/get", controller.GetToken)
		token.GET("/redirect", controller.TokenRedirect) //用于从github接受消息用
		token.PUT("/refresh", controller.RefreshToken)
	}
	user := r.Group("/user")
	{
		managerAuth := user.Use(mid.Login(), mid.ManagerAuth())
		managerAuth.GET("/:username", controller.GetUserInfo)
		managerAuth.GET("/users", controller.ListUsers)
		managerAuth.GET("/createTime", controller.ListUsersByCreateTime)
	}
	star := r.Group("/star", mid.Login())
	{
		star.GET("/:post_id", controller.GetUserStar)
		star.PUT("/", controller.UserStarPost)
	}
	views := r.Group("/views")
	{
		views.GET("/post/:post_id", controller.GetGrowViewsByPostID)
		managerAuth := views.Use(mid.Login(), mid.ManagerAuth())
		managerAuth.GET("/all", controller.ListViewsByCreateTime)
	}
	r.POST("/upload", mid.Login(), mid.ManagerAuth(), controller.Upload)
	logic.Automation(context.Background()) //自动化
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("time", mid.TimeFormat) //注册自定义标签
	}
	return r
}
