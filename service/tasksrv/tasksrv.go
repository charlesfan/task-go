package tasksrv

import (
	srvDomain "github.com/charlesfan/task-go/domain/service"
	repoDomain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity"
)

type taskService struct {
	repo repoDomain.ITaskStore
}

func (s *taskService) Save(f *entity.Task) error {
	return nil
}

func (s *taskService) Find() ([]entity.Task, error) {
	var f []entity.Task

	return f, nil
}

func (s *taskService) Set(f *entity.Task) (*entity.Task, error) {
	return f, nil
}

func (s *taskService) Delete(key int64) error {
	return nil
}

func NewTaskService(r repoDomain.ITaskStore) srvDomain.ITaskService {
	return &taskService{
		repo: r,
	}
}
