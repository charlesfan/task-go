package task_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	routeTask "github.com/charlesfan/task-go/controller/route/task"
	storeDomain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity/config"
	"github.com/charlesfan/task-go/repository/store"
	"github.com/charlesfan/task-go/test"
)

type TaskTestCaseSuite struct {
	c     *gin.Engine
	store storeDomain.ITaskStore
}

func setupTaskTestCaseSuite(t *testing.T) (TaskTestCaseSuite, func(t *testing.T)) {
	config.Init()
	cc := config.New()
	cc.Redis.Addr = "127.0.0.1:6379"
	store.Init(cc)
	repo := store.New()

	s := TaskTestCaseSuite{
		store: repo.TaskStore(),
		c:     gin.New(),
	}

	s.c.Use(gin.Recovery())
	routeTask.ConfigRouterGroup(s.c.Group("/test"))

	return s, func(t *testing.T) {
		defer repo.DB().FlushDB(context.Background())
	}
}

func TestTaskSaveHandler(t *testing.T) {
	s, teardownTestCase := setupTaskTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name         string
		route        string
		method       string
		body         string
		responseCode int
		msg          string
		setupSubTest test.SetupSubTest
	}{
		{
			name:         "success",
			route:        "/test/tasks",
			method:       "POST",
			body:         `{"name": "task-01", "status": 0}`,
			responseCode: http.StatusOK,
			msg:          "success",
			setupSubTest: test.EmptySubTest(),
		},
		{
			name:         "incorrect status",
			route:        "/test/tasks",
			method:       "POST",
			body:         `{"name": "task-01", "status": 2}`,
			responseCode: http.StatusBadRequest,
			msg:          "bad request",
			setupSubTest: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		teardownSubTest := tc.setupSubTest(t)
		defer teardownSubTest(t)

		req := httptest.NewRequest(tc.method, tc.route, strings.NewReader(tc.body))
		req.Header.Set("Content-Type", gin.MIMEJSON)
		rec := httptest.NewRecorder()
		s.c.ServeHTTP(rec, req)

		m := make(map[string]interface{})
		_ = json.Unmarshal(rec.Body.Bytes(), &m)
		assert.Equal(t, tc.responseCode, rec.Code)
		assert.Equal(t, tc.msg, m["msg"])
	}
}
