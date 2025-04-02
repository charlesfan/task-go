package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-redis/redis/v8"

	"github.com/charlesfan/task-go/model"
	"github.com/charlesfan/task-go/utils/log"
)

type redisEngine struct {
	client *redis.Client
}

func newRedis(opts *redis.Options) *redisEngine {
	c := redis.NewClient(opts)
	return &redisEngine{
		client: c,
	}
}

func (e *redisEngine) withPrefix(input string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9]+:\*$`)
	return re.MatchString(input)
}

func (e *redisEngine) getKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	var keys []string
	cursor := uint64(0)
	str := strings.Join([]string{prefix, "*"}, model.DefaultSep)

	for {
		partialKeys, nextCursor, err := e.client.Scan(ctx, cursor, str, 100).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to scan keys: %w", err)
		}

		keys = append(keys, partialKeys...)
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return keys, nil
}

func (e *redisEngine) getValuesInPipeline(ctx context.Context, keys []string) ([]string, error) {
	pipe := e.client.Pipeline()
	cmds := make([]*redis.StringCmd, len(keys))

	for i, key := range keys {
		cmds[i] = pipe.Get(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	result := make([]string, len(keys))
	for i, cmd := range cmds {
		if cmd.Err() == nil {
			result[i] = cmd.Val()
		} else if cmd.Err() == redis.Nil {
			result[i] = ""
		} else {
			log.Errorf("Error fetching value for key %s: %v", keys[i], cmd.Err())
		}
	}

	return result, nil
}

func (e *redisEngine) Status() (string, error) {
	ctx := context.Background()
	return e.client.Ping(ctx).Result()
}

func (e *redisEngine) Client() *redis.Client {
	return e.client
}

func (e *redisEngine) Set(ctx context.Context, key string, value any) (r *Result) {
	r = &Result{}
	if e.client == nil {
		r.setErr(errors.New("redis client is null"))
		return
	}
	r.setErr(e.client.Set(ctx, key, value, 0).Err())
	return
}

func (e *redisEngine) Get(ctx context.Context, key string) (r *Result) {
	var (
		rows []byte
		err  error
	)
	r = &Result{}

	if e.withPrefix(key) {
		a := strings.Split(key, model.DefaultSep)
		if len(a) <= 0 {
			r.setErr(errors.New("fetch prefix error: " + key))
			return
		}
		p := a[0]
		keys, err := e.getKeysByPrefix(ctx, p)
		if err != nil {
			r.setErr(err)
			return
		}

		vals, err := e.getValuesInPipeline(ctx, keys)
		if err != nil {
			r.setErr(err)
			return
		}

		rows, err = json.Marshal(vals)
	} else {
		rows, err = e.client.Get(ctx, key).Bytes()
	}

	r.setErr(err)
	r.setRows(rows)
	return
}

func (e *redisEngine) Delete(ctx context.Context, key string) (r *Result) {
	r = &Result{}
	err := e.client.Del(ctx, key).Err()
	if err != nil {
		r.setErr(err)
	}
	return
}

func (e *redisEngine) FlushDB(ctx context.Context) error {
	return e.client.FlushDB(ctx).Err()
}
