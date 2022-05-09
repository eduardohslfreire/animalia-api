package cache

import (
	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/eduardohslfreire/animalia-api/pkg/cache"
)

// InitCache ...
func InitCache() (*cache.Redis, error) {
	cache := new(cache.Redis)
	cache.Address = env.RedisHost
	cache.Password = env.RedisPassword

	return cache.StartConnection()
}
