package router

import (
	"github.com/gin-gonic/gin"
	"todo/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有数据
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
