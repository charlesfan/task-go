package app

import (
	"github.com/gin-gonic/gin"

	"github.com/charlesfan/task-go/controller/route/task"
)

type RequestHeader struct {
	// TODO authorization
	//Authorization string `json:"authorization"`
}

type Router struct {
	addr   string
	router *gin.Engine
}

func NewRouter(addr string) *Router {
	return &Router{
		addr:   addr,
		router: gin.Default(),
	}
}

func (r *Router) Config() {
	r.router.MaxMultipartMemory = 8 << 20 // 8 MiB

	v := r.router.Group("/sandbox")
	v.Any("", func(c *gin.Context) {
		c.String(200, "power by Charles")
	})
	task.ConfigRouterGroup(v)
}

func (r *Router) Run() {
	r.router.Run(r.addr)
}
