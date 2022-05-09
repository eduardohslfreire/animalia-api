package repository

import (
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
		return "", false
	}

	return actualKey.Val(), true
}

// SetValue ...
func (r *RedisRepository) SetValue(key string, value interface{}, expirationTime int) {
	expiration := time.Duration(expirationTime) * time.Hour

	if err := r.Server.ConnClient.Set(r.Server.ConnClient.Context(), key, value, expiration).Err(); err != nil {
		// Log
	}
}

// DeleteValue ...
func (r *RedisRepository) DeleteValue(key string) {
	r.Server.ConnClient.Del(r.Server.ConnClient.Context(), key)
}

// Lock ...
func (r *RedisRepository) Lock(key string) (*redsync.Mutex, error) {
	mutex := r.Server.ConnSync.NewMutex(key)

	return mutex, mutex.Lock()
}

// Unlock ...
func (r *RedisRepository) Unlock(mutex *redsync.Mutex) (bool, error) {
	return mutex.Unlock()
}
