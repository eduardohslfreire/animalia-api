package middleware

import "github.com/eduardohslfreire/animalia-api/pkg/logger"

// Middleware represent the data-struct for middleware
type Middleware struct {
	Logger logger.GenericLogger
}

// InitMiddleware intialize the middleware
func InitMiddleware() *Middleware {
	return &Middleware{Logger: logger.NewLogger()}
}
