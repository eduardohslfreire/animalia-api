package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

// Redis ...
type Redis struct {
	Cache
	ConnClient *redis.Client
	ConnSync   *redsync.Redsync
}

// StartConnection ...
func (r *Redis) StartConnection() (*Redis, error) {
	r.ConnClient = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       0,
	})

	if _, err := r.ConnClient.Ping(r.ConnClient.Context()).Result(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	pool := goredis.NewPool(r.ConnClient)

	r.ConnSync = redsync.New(pool)

	return r, nil
}
