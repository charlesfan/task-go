package store

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"

	storeDomain "github.com/charlesfan/task-go/domain/store"
	"github.com/charlesfan/task-go/entity/config"
)

const (
	RedisStore = "redis"
)

var (
	once sync.Once
)

type IStore interface {
	Status() (string, error)
	Set(context.Context, string, any) *Result
	Get(context.Context, string) *Result
	Delete(context.Context, string) *Result
}

type onceRepo struct {
	once sync.Once
	repo interface{}
}

type Store struct {
	db IStore

	taskStore onceRepo
}

func (s *Store) DB() IStore {
	return s.db
}

func (s *Store) initRedis(c *config.Redis) {
	opts := &redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	}
	s.db = newRedis(opts)
}

func (s *Store) TaskStore() storeDomain.ITaskStore {
	s.taskStore.once.Do(func() {
		s.taskStore.repo = newTaskStore(s.db)
	})

	return s.taskStore.repo.(storeDomain.ITaskStore)
}

func NewStore(c config.Config) *Store {
	s := &Store{}
	switch c.Store {
	case RedisStore:
		s.initRedis(c.Redis)
		return s
	default:
		return nil
	}
}
