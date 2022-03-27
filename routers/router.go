package routers

import (
	v1 "gin_vue_blog/api/v1"
	"gin_vue_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routerV1 := r.Group("api/v1/")
	{
		// User模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("user/list", v1.GetUserList)
		routerV1.PUT("user/edit/:id", v1.EditUser)
		routerV1.DELETE("user/delete/:id", v1.DeleteUser)
		// 分类模块的路由接口

		// 文章模块的路由接口

	}

	r.Run(utils.HttpPort)
}
