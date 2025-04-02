package store

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	"github.com/charlesfan/task-go/entity/config"
)

const (
	RedisStore = "redis"
)

type IStore interface {
	Status() (string, error)
	Set(context.Context, string, any) *Result
	Get(context.Context, string) *Result
	Delete(context.Context, string) *Result
}

type Result struct {
	err  error
	Rows []byte
}

func (r *Result) setErr(err error) {
	if err != nil && err != redis.Nil {
		r.err = err
	}
}

func (r *Result) setRows(d []byte) {
	r.Rows = d
}

func (r *Result) Err() error {
	return r.err
}

func (r *Result) Bind(dest any) error {
	if err := json.Unmarshal(r.Rows, dest); err != nil {
		return err
	}

	return nil
}

type Store struct {
	db IStore
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
