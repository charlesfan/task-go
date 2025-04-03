package sercvice

import (
	"github.com/charlesfan/task-go/entity"
)

type ITaskService interface {
	Save(*entity.Task) (*entity.Task, error)
	Find() ([]entity.Task, error)
	Set(*entity.Task) (*entity.Task, error)
	Delete(int64) error
}
