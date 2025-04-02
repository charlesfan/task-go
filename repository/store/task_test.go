package store_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity/config"
	"github.com/charlesfan/task-go/model"
	"github.com/charlesfan/task-go/repository/store"
	"github.com/charlesfan/task-go/test"
)

type TaskStoreTestCaseSuite struct {
	store domain.ITaskStore
}

func setupTaskStoreTestCase(t *testing.T) (TaskStoreTestCaseSuite, func(t *testing.T)) {
	config.Init()
	c := config.New()
	c.Redis.Addr = "127.0.0.1:6379"
	repo := store.NewStore(c)
	s := TaskStoreTestCaseSuite{
		store: repo.TaskStore(),
	}

	return s, func(t *testing.T) {
		defer repo.DB().FlushDB(context.Background())
	}
}

func TestTaskStore_Save(t *testing.T) {
	s, teardownTestCase := setupTaskStoreTestCase(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		givenFile    *model.StoreTask
		wantErr      error
		setupSubTest test.SetupSubTest
	}{
		{
			name: "success",
			givenFile: &model.StoreTask{
				Id:   int64(575880729439768611),
				Name: "Task-testing",
				Status: &model.NullInt{
					Int:   0,
					Valid: true,
				},
			},
			wantErr:      nil,
			setupSubTest: test.EmptySubTest(),
		},
		{
			name: "zero id",
			givenFile: &model.StoreTask{
				Id:   int64(0),
				Name: "Task-zero",
				Status: &model.NullInt{
					Int:   0,
					Valid: true,
				},
			},
			wantErr:      errors.New("task id: 0 is not available"),
			setupSubTest: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupSubTest(t)
			defer teardownSubTest(t)

			err := s.store.Save(tc.givenFile)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestTaskStore_Find(t *testing.T) {
	s, teardownTestCase := setupTaskStoreTestCase(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		wantLen      int
		wantErr      error
		setupSubTest test.SetupSubTest
	}{
		{
			name:    "success",
			wantLen: 2,
			wantErr: nil,
			setupSubTest: func(t *testing.T) func(t *testing.T) {
				f1 := &model.StoreTask{
					Id:   int64(575880729439768611),
					Name: "Task-testing-01",
					Status: &model.NullInt{
						Int:   0,
						Valid: true,
					},
				}

				f2 := &model.StoreTask{
					Id:   int64(575880729439768623),
					Name: "Task-testing-02",
					Status: &model.NullInt{
						Int:   0,
						Valid: true,
					},
				}

				err := s.store.Save(f1)
				assert.Nil(t, err)

				err = s.store.Save(f2)
				assert.Nil(t, err)

				return func(t *testing.T) {
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupSubTest(t)
			defer teardownSubTest(t)

			datas, err := s.store.Find()
			t.Log(datas)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantLen, len(datas))
		})
	}
}

func TestTaskStore_Set(t *testing.T) {
	s, teardownTestCase := setupTaskStoreTestCase(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		givenFile    *model.StoreTask
		wantErr      error
		setupSubTest test.SetupSubTest
	}{
		{
			name: "success",
			givenFile: &model.StoreTask{
				Id:   int64(575880729439768611),
				Name: "Task-testing-update",
				Status: &model.NullInt{
					Int:   0,
					Valid: true,
				},
			},
			wantErr: nil,
			setupSubTest: func(t *testing.T) func(t *testing.T) {
				f := &model.StoreTask{
					Id:   int64(575880729439768611),
					Name: "Task-testing",
					Status: &model.NullInt{
						Int:   1,
						Valid: true,
					},
				}

				err := s.store.Save(f)
				assert.Nil(t, err)

				return func(t *testing.T) {
				}
			},
		},
		{
			name: "not found",
			givenFile: &model.StoreTask{
				Id:   int64(575880729439768622),
				Name: "Task-testing-update",
				Status: &model.NullInt{
					Int:   0,
					Valid: true,
				},
			},
			wantErr: errors.New("task: 575880729439768622 not found"),
			setupSubTest: func(t *testing.T) func(t *testing.T) {
				f := &model.StoreTask{
					Id:   int64(575880729439768611),
					Name: "Task-testing",
					Status: &model.NullInt{
						Int:   1,
						Valid: true,
					},
				}

				err := s.store.Save(f)
				assert.Nil(t, err)

				return func(t *testing.T) {
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupSubTest(t)
			defer teardownSubTest(t)

			datas, err := s.store.Set(tc.givenFile)
			t.Log(datas)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestTaskStore_Delete(t *testing.T) {
	s, teardownTestCase := setupTaskStoreTestCase(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		givenId      int64
		wantLen      int
		wantErr      error
		setupSubTest test.SetupSubTest
	}{
		{
			name:    "success",
			givenId: int64(575880729439768611),
			wantLen: 1,
			wantErr: nil,
			setupSubTest: func(t *testing.T) func(t *testing.T) {
				f1 := &model.StoreTask{
					Id:   int64(575880729439768611),
					Name: "Task-testing-01",
					Status: &model.NullInt{
						Int:   0,
						Valid: true,
					},
				}

				f2 := &model.StoreTask{
					Id:   int64(575880729439768623),
					Name: "Task-testing-02",
					Status: &model.NullInt{
						Int:   0,
						Valid: true,
					},
				}

				err := s.store.Save(f1)
				assert.Nil(t, err)

				err = s.store.Save(f2)
				assert.Nil(t, err)

				return func(t *testing.T) {
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupSubTest(t)
			defer teardownSubTest(t)

			err := s.store.Delete(tc.givenId)
			datas, err := s.store.Find()
			t.Log(datas)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantLen, len(datas))
		})
	}
}
