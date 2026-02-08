package router

import (
	"exchange_backend/controllers"
	"exchange_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//路由分组（Group）。它创建了一个路由前缀，所有在这个组内的路由都将以 /api/auth 开头
	auth := r.Group("/api/auth")
	{
		//注册了一个 POST 请求的路由，当用户向这个地址发送 POST 请求时，执行后面的匿名函数
		auth.POST("/login", controllers.Login)
		//同上
		auth.POST("/register", controllers.Register)

	}
	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMidddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticlesByID)

		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("/articles/:id/like", controllers.GetArticleLikes)
	}
	return r
}
