package router

import (
	"example.com/blog/controllers"
	"example.com/blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() //初始化引擎实例

	auth := r.Group("/api/auth")                //不需要鉴权
	auth.Use(middleware.LogRequestMiddleware()) //日志
	auth.Use(middleware.ErrHandler())           //异常中间件
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
		auth.POST("/getPostlist", controllers.GetPostList)
		auth.POST("/getPostById", controllers.GetPostById)
		auth.POST("/getcommentsbypostID", controllers.GetCommentsByPostID)
	}

	api := r.Group("/api")
	api.Use(middleware.LogRequestMiddleware()) //日志
	api.Use(middleware.AuthMiddleWare())       //需要鉴权
	auth.Use(middleware.ErrHandler())          //异常中间件
	{
		api.POST("/createpost", controllers.CreatePost)
		api.POST("/updatePost", controllers.UpdatePost)
		api.POST("/deletePost", controllers.DeletePost)
		api.POST("/createcomment", controllers.CreateComment)

	}

	return r

}
