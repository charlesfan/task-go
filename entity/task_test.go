package entity_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlesfan/task-go/entity"
	"github.com/charlesfan/task-go/model"
)

func Test_TaskStoreModel(t *testing.T) {
	id := int64(575880729439768609)
	name := "Task-01"
	status := entity.TaskIncomplete

	tt := &entity.Task{
		Id:     id,
		Name:   name,
		Status: &status,
	}

	want := &model.StoreTask{
		Id: strings.Join([]string{
			model.TaskPrfix,
			strconv.FormatInt(id, 10),
		}, model.TaskSep),
		Name: name,
		Status: &model.NullInt{
			Int:   status,
			Valid: true,
		},
	}

	assert.Equal(t, tt.StoreModel(), want)
}
