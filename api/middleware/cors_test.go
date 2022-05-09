package middleware_test

import (
	"net/http"
	"testing"

	"github.com/eduardohslfreire/animalia-api/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouterCors() *gin.Engine {
	m := middleware.InitMiddleware()
	router := gin.New()
	router.Use(m.CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(http.StatusOK, "patch")
	})
	router.OPTIONS("/", func(c *gin.Context) {
		c.String(http.StatusOK, "options")
	})
	return router
}

func TestCors(t *testing.T) {
	router := setupRouterCors()
	w := performRequest(router, "GET", "https://gist.github.com")
	assert.Equal(t, 200, w.Code)
	w = performRequest(router, "GET", "https://github.com")
	assert.Equal(t, 200, w.Code)
	w = performRequest(router, "OPTIONS", "https://github.com")
	assert.Equal(t, 204, w.Code)
}
