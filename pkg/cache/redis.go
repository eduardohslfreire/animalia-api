package cache

import (
	"fmt"

	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

// Redis ...
type Redis struct {
	Cache
	ConnClient *redis.Client
	ConnSync   *redsync.Redsync
	Logger     *logger.GenericLogger
}

// StartConnection ...
func (r *Redis) StartConnection() (*Redis, error) {
	r.Logger = logger.NewGenericLogger()

	r.ConnClient = redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       0,
	})

	if _, err := r.ConnClient.Ping(r.ConnClient.Context()).Result(); err != nil {
		r.Logger.LogIt("ERROR", fmt.Sprintf("PING-REDIS-ERROR] - Error establishing communication with redis. %s", err.Error()), nil)
		return nil, err
	}

	pool := goredis.NewPool(r.ConnClient)

	r.ConnSync = redsync.New(pool)

	r.Logger.LogIt("INFO", "[REDIS-INIT] - Connection to redis started", nil)

	return r, nil
}
