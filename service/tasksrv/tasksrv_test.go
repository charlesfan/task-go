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
