package store

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

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
