package routers

import (
	v1 "gin_vue_blog/api/v1"
	"gin_vue_blog/middleware"
	"gin_vue_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.UpdateUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.UpdateCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.UpdateArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		auth.POST("upload", v1.Upload)
	}
	routerV1 := r.Group("/api/v1")
	{
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.GET("categories", v1.GetCategories)
		routerV1.GET("articles", v1.GetArticles)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.GET("article/info/:id", v1.GetArtInfo)
		routerV1.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
