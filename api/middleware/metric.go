package middleware

import (
	"regexp"
	"strconv"

	"github.com/eduardohslfreire/animalia-api/pkg/metric"
	"github.com/gin-gonic/gin"
)

// MetricMiddleware ...
func (m *Middleware) MetricMiddleware(service metric.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpMetric := metric.NewHTTP(toGenericPath(c.Request.URL.Path), c.Request.Method)
		httpMetric.Started()

		c.Next()

		httpMetric.Finished()
		httpMetric.StatusCode = strconv.Itoa(c.Writer.Status())
		service.SaveHTTP(httpMetric)
	}
}

func toGenericPath(path string) string {
	rMiddlePath := regexp.MustCompile(`(\/)(\d+)(\/)`)
	rEndPath := regexp.MustCompile(`(\/)(\d+)$`)

	return rEndPath.ReplaceAllString(rMiddlePath.ReplaceAllString(path, "${1}id${3}"), "${1}id")
}
