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

type TaskId string

func (s TaskId) String() string {
	return string(s)
}

func (s TaskId) Int64() int64 {
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

func NewTaskId(d int64) string {
	s := strconv.FormatInt(d, 10)
	return strings.Join([]string{TaskPrfix, s}, TaskSep)
}

type StoreTask struct {
	Id     string
	Name   string
	Status *NullInt
}
