package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/eduardohslfreire/animalia-api/api/middleware"
	v1 "github.com/eduardohslfreire/animalia-api/api/v1"
	"github.com/eduardohslfreire/animalia-api/api/validation"
	"github.com/eduardohslfreire/animalia-api/config/cache"
	"github.com/eduardohslfreire/animalia-api/config/db"
	"github.com/eduardohslfreire/animalia-api/config/env"

	"github.com/eduardohslfreire/animalia-api/infrastructure/repository"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	Route  *gin.Engine
	Logger *logger.GenericLogger
}

// Initialize ...
func (s *Server) Initialize() error {
	// Inicializando componentes do Log
	s.Logger = logger.NewGenericLogger()

	database, err := db.InitDatabase()
	if err != nil {
		return err
	}

	cache, err := cache.InitCache()
	if err != nil {
		return err
	}

	// Inicializando os repositórios
	redisRepository := repository.NewRedisRepository(cache)

	citizenRepository := repository.NewCitizenRepository(database)
	roleRepository := repository.NewRoleRepository(database)

	// Inicializando os serviços
	citizenUsecase := usecase.NewCitizenUsecase(citizenRepository, roleRepository, redisRepository)
	roleUsecase := usecase.NewRoleUsecase(roleRepository)

	// Inicializando rota
	s.Route = gin.New()

	// Middleware
	m := middleware.InitMiddleware()
	s.Route.Use(gin.Recovery())
	s.Route.Use(m.CORSMiddleware())

	validation.RegisterCustomValidations()

	// prometheusService, err := metric.NewPrometheusService()
	// if err != nil {
	// 	return err
	// }
	// s.Route.Use(m.MetricMiddleware(prometheusService))

	//s.Route.Static("/docs", "swaggerui")
	routerGroup := s.Route.Group("/api/v1")
	routerGroup.Use(m.ErrorMiddleware())

	// Inicializando a rota das APIs
	v1.InitRouting(routerGroup, citizenUsecase, roleUsecase)

	return nil
}

// StartServer ...
func (s *Server) StartServer() {
	if err := s.Initialize(); err != nil {
		s.Logger.LogIt("ERROR", fmt.Sprintf("[SERVER-START-ERROR] - Erro ao iniciar o servidor. Motivo: %s", err.Error()), nil)
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", env.AppPort),
		ReadTimeout:  time.Duration(env.HTTPReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.HTTPWriteTimeout) * time.Second,
		Handler:      s.Route,
	}

	s.Logger.LogIt("INFO", fmt.Sprintf("[SERVER-START] - Iniciando servidor na porta %s", env.AppPort), nil)
	if err := httpServer.ListenAndServe(); err != nil {
		s.Logger.LogIt("ERROR", fmt.Sprintf("[SERVER-START-ERROR] - Erro ao iniciar o servidor. Motivo: %s", err.Error()), nil)
		os.Exit(1)
	}
}
