package repository

import (
	"fmt"
	"time"

	"github.com/eduardohslfreire/animalia-api/pkg/cache"
	"github.com/go-redsync/redsync/v4"
)

// RedisRepository ...
type RedisRepository struct {
	Server *cache.Redis
}

// NewRedisRepository ...
func NewRedisRepository(Server *cache.Redis) IRedisRepository {
	return &RedisRepository{Server}
}

// GetValue ...
func (r *RedisRepository) GetValue(key string) (string, bool) {
	actualKey := r.Server.ConnClient.Get(r.Server.ConnClient.Context(), key)
	if actualKey.Err() != nil || len(actualKey.Val()) == 0 {
		r.Server.Logger.LogIt("ERROR", fmt.Sprintf("[GET-REDIS-ERROR] - Failed to get value from key %s", key), nil)
		return "", false
	}

	return actualKey.Val(), true
}

// SetValue ...
func (r *RedisRepository) SetValue(key string, value interface{}, expirationTime int) {
	expiration := time.Duration(expirationTime) * time.Hour

	if err := r.Server.ConnClient.Set(r.Server.ConnClient.Context(), key, value, expiration).Err(); err != nil {
		r.Server.Logger.LogIt("ERROR", fmt.Sprintf("[SET-REDIS-ERROR] - Failed to set value from key %s", key), nil)
	}
}

// DeleteValue ...
func (r *RedisRepository) DeleteValue(key string) {
	if err := r.Server.ConnClient.Del(r.Server.ConnClient.Context(), key).Err(); err != nil {
		r.Server.Logger.LogIt("ERROR", fmt.Sprintf("[DEL-REDIS-ERROR] - Failed to delete value from key %s", key), nil)
	}
}

// Lock ...
func (r *RedisRepository) Lock(key string) (*redsync.Mutex, error) {
	mutex := r.Server.ConnSync.NewMutex(key)
	if err := mutex.Lock(); err != nil {
		r.Server.Logger.LogIt("ERROR", fmt.Sprintf("[LOCK-REDIS-ERROR] - Failed to get lock from key %s", key), nil)
		return nil, err
	}
	return mutex, nil
}

// Unlock ...
func (r *RedisRepository) Unlock(mutex *redsync.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		r.Server.Logger.LogIt("ERROR", fmt.Sprintf("[UNLOCK-REDIS-ERROR] - Failed to get unlock from key %s", mutex.Name()), nil)
		return err
	}
	return nil
}
