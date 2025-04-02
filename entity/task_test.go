package entity_test

import (
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
		Id:   id,
		Name: name,
		Status: &model.NullInt{
			Int:   status,
			Valid: true,
		},
	}

	assert.Equal(t, tt.StoreModel(), want)
}

func Test_FromStoreModel(t *testing.T) {
	id := int64(575880729439768609)
	name := "Task-01"
	status := entity.TaskIncomplete

	tt := &entity.Task{
		Id:     id,
		Name:   name,
		Status: &status,
	}

	mt := &model.StoreTask{
		Id:   id,
		Name: name,
		Status: &model.NullInt{
			Int:   status,
			Valid: true,
		},
	}

	want := &entity.Task{}
	want.FromStoreModel(mt)

	assert.Equal(t, tt, want)
}
