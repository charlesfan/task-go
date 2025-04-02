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

type StoreKeyTask string

func (s StoreKeyTask) String() string {
	return string(s)
}

func (s StoreKeyTask) Int64() int64 {
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

func newStoreKeyTask(d int64) StoreKeyTask {
	s := strconv.FormatInt(d, 10)
	id := strings.Join([]string{TaskPrfix, s}, TaskSep)
	return StoreKeyTask(id)
}

type StoreTask struct {
	Id     int64
	Name   string
	Status *NullInt
}

func (t *StoreTask) Key() string {
	if t.Id <= 0 {
		log.Errorf("id is null: %+v", t)
		return ""
	}

	return newStoreKeyTask(t.Id).String()
}
