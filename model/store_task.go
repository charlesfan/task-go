package model

import (
	"strconv"
	"strings"

	"github.com/charlesfan/task-go/utils/log"
)

const (
	TaskPrfix = "task"
	TaskSep   = ":"
)

type StoreTaskId string

func (s StoreTaskId) String() string {
	return string(s)
}

func (s StoreTaskId) Int64() int64 {
	a := strings.Split(s.String(), TaskSep)
	if len(a) < 2 {
		log.Error("fail to int64: ", s)
		return 0
	}

	i, err := strconv.ParseInt(a[1], 10, 64)
	if err != nil {
		log.Error("fail to Parse int64: ", err)
		return 0
	}

	return i
}

func NewStoreTaskId(d int64) StoreTaskId {
	s := strconv.FormatInt(d, 10)
	id := strings.Join([]string{TaskPrfix, s}, TaskSep)
	return StoreTaskId(id)
}

type StoreTask struct {
	Id     StoreTaskId
	Name   string
	Status *NullInt
}
