package middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var header map[string][]string

func performRequest(r *gin.Engine, method, origin string) *httptest.ResponseRecorder {
	return performRequestWithHeaders(r, method, origin, http.Header{})
}

func performRequestWithHeaders(r *gin.Engine, method, origin string, header http.Header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, "/", nil)

	req.Host = header.Get("Host")
	header.Del("Host")
	if len(origin) > 0 {
		header.Set("Origin", origin)
	}
	req.Header = header
	req.Header.Set("Authorization", "Bearer aaa")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
