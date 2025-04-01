package entity

import (
	"github.com/charlesfan/task-go/model"
)

const (
	TaskIncomplete = 0
	TaskCompleted  = 1
)

type Task struct {
	Id     int64
	Name   string
	Status *int
}

func (t *Task) StoreModel() *model.StoreTask {
	mst := &model.StoreTask{
		Id:   model.NewTaskId(t.Id),
		Name: t.Name,
	}

	if t.Status != nil {
		mst.Status = &model.NullInt{
			Int:   *t.Status,
			Valid: true,
		}
	}

	return mst
}
