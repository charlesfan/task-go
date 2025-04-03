package task

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/charlesfan/task-go/controller/resp"
	srvDomain "github.com/charlesfan/task-go/domain/service"
	"github.com/charlesfan/task-go/entity"
	"github.com/charlesfan/task-go/entity/errcode"
	"github.com/charlesfan/task-go/service"
	"github.com/charlesfan/task-go/utils/log"
)

type TaskRequest struct {
	Id     int64  `json:"id" form:"id" query:"id"`
	Name   string `json:"name" form:"name" query:"name"`
	Status *int   `json:"status" form:"status" query:"status"`
}

func (r *TaskRequest) entityTask() *entity.Task {
	return &entity.Task{
		Id:     r.Id,
		Name:   r.Name,
		Status: r.Status,
	}
}

type TaskController struct {
	taskSrv srvDomain.ITaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskSrv: service.New().TaskSrv(),
	}
}

func (a *TaskController) Save(c *gin.Context) {
	var r TaskRequest

	if err := c.Bind(&r); err != nil {
		log.Error("task binding error: ", err)
		resp.WriteResponse(c, errcode.New(errcode.ErrorCodeBadRequest), nil)

		return
	}

	re, err := a.taskSrv.Save(r.entityTask())
	resp.WriteResponse(c, err, re)
}

func (a *TaskController) Set(c *gin.Context) {
	var r TaskRequest

	if err := c.Bind(&r); err != nil {
		log.Error("task binding error: ", err)
		resp.WriteResponse(c, errcode.New(errcode.ErrorCodeBadRequest), nil)

		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		resp.WriteResponse(c, errcode.New(errcode.ErrorCodeBadRequest), nil)

		return
	}

	r.Id = id
	re, err := a.taskSrv.Save(r.entityTask())
	resp.WriteResponse(c, err, re)
}

func (a *TaskController) Find(c *gin.Context) {
	datas, err := a.taskSrv.Find()
	resp.WriteResponse(c, err, datas)
}

func (a *TaskController) Del(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		resp.WriteResponse(c, errcode.New(errcode.ErrorCodeBadRequest), nil)

		return
	}

	resp.WriteResponse(c, a.taskSrv.Delete(id), nil)
}
