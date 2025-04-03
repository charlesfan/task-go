package tasksrv

import (
	srvDomain "github.com/charlesfan/task-go/domain/service"
	repoDomain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity"
	"github.com/charlesfan/task-go/entity/errcode"
	"github.com/charlesfan/task-go/utils/log"
)

type taskService struct {
	repo repoDomain.ITaskStore
}

func (s *taskService) setTask(f *entity.Task) error {
	if f.Id <= 0 {
		log.Errorf("task id is not available: %d", f.Id)
		return errcode.New(errcode.ErrorCodeBadRequest)
	}

	if f.Status != nil {
		status := *f.Status
		switch status {
		case entity.TaskIncomplete, entity.TaskCompleted:
		default:
			log.Errorf("task status is not available: %d", status)
			return errcode.New(errcode.ErrorCodeBadRequest)
		}
	} else {
		status := entity.TaskIncomplete
		f.Status = &status
	}

	return nil
}

func (s *taskService) Save(f *entity.Task) error {
	if err := s.setTask(f); err != nil {
		return err
	}

	if err := s.repo.Save(f.StoreModel()); err != nil {
		log.Error(err)
		return errcode.New(errcode.ErrorCodeTaskErr)
	}

	return nil
}

func (s *taskService) Find() ([]entity.Task, error) {
	rows, err := s.repo.Find()
	if err != nil {
		log.Error(err)
		return nil, errcode.New(errcode.ErrorCodeTaskErr)
	}

	f := make([]entity.Task, len(rows))
	for k, v := range rows {
		t := entity.Task{}
		t.FromStoreModel(&v)
		f[k] = t
	}

	return f, nil
}

func (s *taskService) Set(f *entity.Task) (*entity.Task, error) {
	if err := s.setTask(f); err != nil {
		return nil, err
	}

	row, err := s.repo.Set(f.StoreModel())
	if err != nil {
		log.Error(err)
		return nil, errcode.New(errcode.ErrorCodeTaskErr)
	}

	t := &entity.Task{}
	t.FromStoreModel(row)
	return t, nil
}

func (s *taskService) Delete(key int64) error {
	if err := s.repo.Delete(key); err != nil {
		log.Error(err)
		return errcode.New(errcode.ErrorCodeTaskErr)
	}
	return nil
}

func NewTaskService(r repoDomain.ITaskStore) srvDomain.ITaskService {
	return &taskService{
		repo: r,
	}
}
