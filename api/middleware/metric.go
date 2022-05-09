package middleware

import (
	"strconv"

	"github.com/eduardohslfreire/animalia-api/pkg/metric"
	"github.com/gin-gonic/gin"
)

// MetricMiddleware ...
func (m *Middleware) MetricMiddleware(service metric.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpMetric := metric.NewHTTP(c.Request.URL.Path, c.Request.Method)
		httpMetric.Started()

		c.Next()

		httpMetric.Finished()
		httpMetric.StatusCode = strconv.Itoa(c.Request.Response.StatusCode)
		service.SaveHTTP(httpMetric)
	}
}
