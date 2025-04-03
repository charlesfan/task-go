package task

import (
	"github.com/gin-gonic/gin"
)

func ConfigRouterGroup(group *gin.RouterGroup) {
	c := NewTaskController()

	taskGroup := group.Group("/tasks")
	{
		taskGroup.POST("", c.Save)
		taskGroup.GET("", c.Find)
		taskGroup.PUT(":id", c.Set)
		taskGroup.DELETE(":id", c.Del)
	}
}
