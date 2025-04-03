package entity

import (
	"github.com/charlesfan/task-go/model"
)

const (
	TaskIncomplete = 0
	TaskCompleted  = 1
)

type Task struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status *int   `json:"status"`
}

func (t *Task) StoreModel() *model.StoreTask {
	mst := &model.StoreTask{
		Id:   t.Id,
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

func (t *Task) FromStoreModel(m *model.StoreTask) {
	t.Id = m.Id
	t.Name = m.Name

	if m.Status != nil && m.Status.Valid {
		status := m.Status.Int
		t.Status = &status
	}
}
