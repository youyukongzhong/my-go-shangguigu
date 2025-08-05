package router

import (
	"github.com/gin-gonic/gin"
	"learnGo/controllers"
	"net/http"
)

//router/：路由定义，请求转发给谁
//Router（路由员）：接待顾客，说“你要吃什么？跟我来”。它决定把请求交给哪个服务员（Controller）。

func Router() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", controllers.GetUserInfo)
		userGroup.POST("/list", controllers.GetList)
		userGroup.PUT("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "add user")
		})
		userGroup.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "delete user")
		})
	}

	return r
}
