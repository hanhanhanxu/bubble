package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//静态文件、html 路由
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	group := r.Group("v1")
	{
		group.POST("/todo", controller.CreateTodoHandler)
		group.GET("/todo", controller.GetTodosHandler)
		group.PUT("/todo/:id", controller.UpdateTodoHandler)
		group.DELETE("/todo/:id", controller.DeleteTodoHandler)
	}
	return r
}
