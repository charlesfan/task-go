package task

import (
	"github.com/gin-gonic/gin"
)

func ConfigRouterGroup(group *gin.RouterGroup) {
	c := NewTaskController()

	employeeGroup := group.Group("/tasks")
	{
		employeeGroup.POST("", c.Save)
	}
}
