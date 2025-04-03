package tasksrv_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	srvDomain "github.com/charlesfan/task-go/domain/service"
	storeDomain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity"
	"github.com/charlesfan/task-go/entity/config"
	"github.com/charlesfan/task-go/entity/errcode"
	"github.com/charlesfan/task-go/model"
	"github.com/charlesfan/task-go/repository/store"
	"github.com/charlesfan/task-go/service"
	"github.com/charlesfan/task-go/test"
)

var (
	taskIncomplete = entity.TaskIncomplete
	taskCompleted  = entity.TaskCompleted
)

type TaskServiceTestCaseSuite struct {
	store   storeDomain.ITaskStore
	service srvDomain.ITaskService
}

func setupTaskServiceTestCaseSuite(t *testing.T) (TaskServiceTestCaseSuite, func(t *testing.T)) {
	config.Init()
	c := config.New()
	c.Redis.Addr = "127.0.0.1:6379"
	store.Init(c)
	repo := store.New()
	srv := service.New()

	s := TaskServiceTestCaseSuite{
		store:   repo.TaskStore(),
		service: srv.TaskSrv(),
	}

	return s, func(t *testing.T) {
		defer repo.DB().FlushDB(context.Background())
	}
}

func TestTaskService_Save(t *testing.T) {
	s, teardownTestCase := setupTaskServiceTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		givenFile    *entity.Task
		wantErr      error
		setupSubTest test.SetupSubTest
	}{
		{
			name: "success",
			givenFile: &entity.Task{
				Id:     int64(575880729439768651),
				Name:   "Tasksrv-testing",
				Status: &taskIncomplete,
			},
			wantErr:      nil,
			setupSubTest: test.EmptySubTest(),
		},
		{
			name: "zero id",
			givenFile: &entity.Task{
				Id:     int64(0),
				Name:   "Tasksrv-testing",
				Status: &taskIncomplete,
			},
			wantErr:      errcode.New(errcode.ErrorCodeBadRequest),
			setupSubTest: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupSubTest(t)
			defer teardownSubTest(t)

			err := s.service.Save(tc.givenFile)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestTaskService_Find(t *testing.T) {
	s, teardownTestCase := setupTaskServiceTestCaseSuite(t)
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

			datas, err := s.service.Find()
			t.Log(datas)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantLen, len(datas))
		})
	}
}
