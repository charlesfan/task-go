package store_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlesfan/task-go/entity"
	"github.com/charlesfan/task-go/entity/config"
	"github.com/charlesfan/task-go/model"
	"github.com/charlesfan/task-go/repository/store"
)

var (
	ctx = context.Background()
	rdb store.IStore
)

func Init(t *testing.T) {
	config.Init()
	c := config.New()
	c.Redis.Addr = "127.0.0.1:6379"
	t.Log(c.Redis)
	ss := store.NewStore(c)
	rdb = ss.DB()
}

func Test_Set(t *testing.T) {
	Init(t)

	id := int64(575880729439768609)
	name := "Task-01"
	status := entity.TaskIncomplete

	tt := &entity.Task{
		Id:     id,
		Name:   name,
		Status: &status,
	}

	st := tt.StoreModel()
	data, err := json.Marshal(st)
	assert.Nil(t, err)
	err = rdb.Set(ctx, st.Key(), data).Err()
	assert.Nil(t, err)
}

func Test_Get(t *testing.T) {
	Init(t)

	key := "task:575880729439768609"
	r := rdb.Get(ctx, key)
	assert.Nil(t, r.Err())

	var val model.StoreTask

	r.Bind(&val)
	t.Log(string(r.Rows))
	t.Log(val)
}

func Test_PipeGet(t *testing.T) {
	Init(t)

	key := "task:*"
	r := rdb.Get(ctx, key)
	assert.Nil(t, r.Err())

	var (
		strArr []string
		val    []model.StoreTask
	)

	err := r.Bind(&strArr)
	assert.Nil(t, err)
	for i := 0; i < len(strArr); i++ {
		var v model.StoreTask
		err := json.Unmarshal([]byte(strArr[i]), &v)
		assert.Nil(t, err)
		val = append(val, v)
	}
	t.Log(val)
}

func Test_Del(t *testing.T) {
	Init(t)

	key := "task:575880729439768609"
	err := rdb.Delete(ctx, key).Err()
	assert.Nil(t, err)
	r := rdb.Get(ctx, key)
	assert.Nil(t, r.Err())
	assert.Equal(t, []byte(nil), r.Rows)
}
