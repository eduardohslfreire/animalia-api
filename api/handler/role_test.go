package handler_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohslfreire/animalia-api/api/handler"
	"github.com/eduardohslfreire/animalia-api/api/middleware"
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
)

func TestFind(t *testing.T) {

	t.Run("success - find role by id", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		mockRole := &entity.Role{ID: 1, Name: "Civil", Single: false}
		mockeRoleUsecase.On("FindByID", mocks.AnythingOfType("uint")).Return(mockRole, nil)

		code, response := makeRequest(r, http.MethodGet, "/api/v1/roles/1")

		assert.Equal(t, `{"id":1,"name":"Civil"}`, response)
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("error - find role by id", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		mockeRoleUsecase.On("FindByID", mocks.AnythingOfType("uint")).Return(nil, fmt.Errorf("error"))

		code, response := makeRequest(r, http.MethodGet, "/api/v1/roles/1")

		assert.Equal(t, `{"status_code":500,"message":"A system or application error occurred."}`, response)
		assert.Equal(t, http.StatusInternalServerError, code)
	})

	t.Run("error - find role by invalid id", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		mockeRoleUsecase.On("FindByID", mocks.AnythingOfType("uint")).Return(nil, fmt.Errorf("error"))

		code, response := makeRequest(r, http.MethodGet, "/api/v1/roles/1A")

		assert.Equal(t, `{"status_code":400,"message":"Invalid ID"}`, response)
		assert.Equal(t, http.StatusBadRequest, code)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("success - find all roles", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		mockRoles := &entity.Roles{entity.Role{ID: 1, Name: "Civil", Single: false}}
		mockeRoleUsecase.On("FindAll").Return(mockRoles, nil)

		code, response := makeRequest(r, http.MethodGet, "/api/v1/roles")

		assert.Equal(t, `[{"id":1,"name":"Civil"}]`, response)
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("success - no roles", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		//mockRoles := make(entity.Roles, 0)
		mockeRoleUsecase.On("FindAll").Return(nil, nil)

		code, _ := makeRequest(r, http.MethodGet, "/api/v1/roles")

		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("error - find all roles", func(t *testing.T) {
		r, v1 := initRouter()

		mockeRoleUsecase := new(mock.RoleUsecase)
		handler.NewRoleHandler(v1, mockeRoleUsecase)

		mockeRoleUsecase.On("FindAll").Return(nil, fmt.Errorf("error"))

		code, response := makeRequest(r, http.MethodGet, "/api/v1/roles")

		assert.Equal(t, `{"status_code":500,"message":"A system or application error occurred."}`, response)
		assert.Equal(t, http.StatusInternalServerError, code)
	})
}

func initRouter() (*gin.Engine, *gin.RouterGroup) {
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.Use(middleware.InitMiddleware().ErrorMiddleware())

	return r, v1
}

func makeRequest(r *gin.Engine, method string, url string) (int, string) {
	req, _ := http.NewRequest(method, url, nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	body, _ := ioutil.ReadAll(rr.Body)
	return rr.Code, string(body)
}
