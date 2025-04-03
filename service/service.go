package service

import (
	"sync"

	srvDomain "github.com/charlesfan/task-go/domain/service"
	"github.com/charlesfan/task-go/repository/store"
	"github.com/charlesfan/task-go/service/tasksrv"
)

var (
	defaultSrv *serviceFactory
	once       sync.Once
)

type onceService struct {
	once sync.Once
	srv  interface{}
}

type serviceFactory struct {
	taskSrv onceService
}

func (s *serviceFactory) TaskSrv() srvDomain.ITaskService {
	s.taskSrv.once.Do(func() {
		db := store.New().TaskStore()
		s.taskSrv.srv = tasksrv.NewTaskService(db)
	})
	return s.taskSrv.srv.(srvDomain.ITaskService)
}

func New() *serviceFactory {
	once.Do(func() {
		defaultSrv = &serviceFactory{}
	})
	return defaultSrv
}
