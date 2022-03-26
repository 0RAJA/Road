package routing

import (
	_ "github.com/0RAJA/Road/docs"
	"github.com/0RAJA/Road/internal/controller"
	mid "github.com/0RAJA/Road/internal/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouting() *gin.Engine {
	r := gin.New()
	r.Use(mid.Cors(), mid.GinRecovery(true), mid.GinLogger())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	comment := r.Group("/comment")
	{
		loginAuth := comment.Use(mid.AuthMiddleware())
		comment.GET("/", controller.ListComments)
		loginAuth.POST("/", controller.AddComment)
		loginAuth.DELETE("/:post_id", controller.DeleteComment)
		loginAuth.PUT("/", controller.ModifyComment)
	}
	manager := r.Group("/manager", mid.AuthMiddleware(), mid.ManagerAuth())
	{
		manager.GET("/check/:username", controller.CheckManagerName)
		manager.POST("/login", controller.LoginManager)
		manager.POST("/create", controller.AddManager)
		manager.PUT("/", controller.UpdateManager)
		manager.DELETE("/:post_id", controller.DeleteManager)
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
			infos.GET("/star/grow", controller.ListPostInfosOrderByGrowingStar)
			infos.GET("/visit/grow", controller.ListPostInfosOrderByGrowingVisited)
		}
		managerAuth := post.Use(mid.AuthMiddleware(), mid.ManagerAuth())
		managerAuth.PUT("/update", controller.UpdatePost)
		managerAuth.PUT("/delete", controller.ModifyPostDeleted)
		managerAuth.PUT("/public", controller.ModifyPostPublic)
		managerAuth.POST("/create", controller.AddPost)
		managerAuth.DELETE("/:post_id", controller.RealDeletePost)
	}
	postTag := r.Group("/postTag")
	{
		managerAuth := postTag.Use(mid.AuthMiddleware(), mid.ManagerAuth())
		managerAuth.POST("/", controller.AddPostTag)
		managerAuth.DELETE("/", controller.DeletePostTag)
		postTag.GET("/tags", controller.ListTagsByPostID)
		postTag.GET("/infos", controller.ListPostInfosByTagID)
	}
	tag := r.Group("/tag")
	{
		managerAuth := tag.Use(mid.AuthMiddleware(), mid.ManagerAuth())
		managerAuth.POST("/", controller.AddTag)
		managerAuth.DELETE("/:tag_id", controller.DeleteTag)
		managerAuth.PUT("/", controller.UpdateTag)
		tag.GET("/", controller.ListTags)
		tag.GET("/check/:tag_name", controller.CheckTagName)
	}
	token := r.Group("/token", controller.GetToken)
	{
		noLoginAuth := token.Use(mid.NoLogin())
		noLoginAuth.GET("/get", controller.GetToken)
		token.GET("/redirect", controller.TokenRedirect) //用于从github接受消息用
		token.PUT("/refresh", controller.RefreshToken)
	}
	user := r.Group("/user")
	{
		managerAuth := user.Use(mid.AuthMiddleware(), mid.ManagerAuth())
		managerAuth.GET("/:username", controller.GetUserInfo)
		managerAuth.GET("/users", controller.ListUsers)
		managerAuth.GET("/createTime", controller.ListUsersByCreateTime)
	}
	star := r.Group("/star", mid.AuthMiddleware())
	{
		star.GET("/:post_id", controller.GetUserStar)
		star.PUT("/", controller.UserStarPost)
	}
	views := r.Group("/views")
	{
		views.GET("/post/:post_id", controller.GetGrowViewsByPostID)
		managerAuth := views.Use(mid.AuthMiddleware(), mid.ManagerAuth())
		managerAuth.GET("/all", controller.ListViewsByCreateTime)
	}
	r.POST("/upload", mid.AuthMiddleware(), mid.ManagerAuth(), controller.Upload)
	return r
}
