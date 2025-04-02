package store

import (
	"github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/model"
)

type taskStore struct {
	rdb IStore
}

func (s *taskStore) Save(f *model.StoreTask) error {
	return nil
}

func (s *taskStore) Find() ([]model.StoreTask, error) {
	var f []model.StoreTask
	return f, nil
}

func (s *taskStore) Set(f *model.StoreTask) (*model.StoreTask, error) {
	return f, nil
}

func (s *taskStore) Delete(key int64) error {
	return nil
}

func newTaskStore(db IStore) store.ITaskStore {
	return &taskStore{
		rdb: db,
	}
}
