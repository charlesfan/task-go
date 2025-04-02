package store

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/model"
)

type taskStore struct {
	rdb IStore
}

func (s *taskStore) Save(f *model.StoreTask) error {
	if f.Id <= 0 {
		return fmt.Errorf("task id: %d is not available", f.Id)
	}

	key := f.Key()
	if key == "" {
		return fmt.Errorf("the key of %+v is null", f)
	}

	ctx := context.Background()
	data, err := json.Marshal(f)
	if err != nil {
		return err
	}

	r := s.rdb.Set(ctx, key, data)

	return r.Err()
}

func (s *taskStore) Find() ([]model.StoreTask, error) {
	key := strings.Join([]string{model.TaskPrefix, "*"}, model.TaskSep)
	ctx := context.Background()
	result := s.rdb.Get(ctx, key)

	var (
		strArr []string
		f      []model.StoreTask
	)

	if err := result.Bind(&strArr); err != nil {
		return nil, err
	}

	for i := 0; i < len(strArr); i++ {
		var v model.StoreTask
		if err := json.Unmarshal([]byte(strArr[i]), &v); err != nil {
			return nil, err
		}
		f = append(f, v)
	}

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
