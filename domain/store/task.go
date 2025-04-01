package store

import (
	"github.com/charlesfan/task-go/model"
)

type ITaskStore interface {
	Save(*model.StoreTask) error
	Find() ([]*model.StoreTask, error)
	Set(int64) (*model.StoreTask, error)
	Delete(int64) error
}
