package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/eduardohslfreire/animalia-api/api/handler"
	"github.com/eduardohslfreire/animalia-api/api/middleware"
	"github.com/eduardohslfreire/animalia-api/api/validation"
	"github.com/eduardohslfreire/animalia-api/config/cache"
	"github.com/eduardohslfreire/animalia-api/config/db"
	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/eduardohslfreire/animalia-api/infrastructure/repository"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/eduardohslfreire/animalia-api/pkg/metric"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server ...
type Server struct {
	Route  *gin.Engine
	Logger *logger.GenericLogger
}

// Initialize ...
func (s *Server) Initialize() error {
	// Initializing log components
	s.Logger = logger.NewGenericLogger()

	// Initializing database
	database, err := db.InitDatabase()
	if err != nil {
		return err
	}

	// Initializing cache server
	cache, err := cache.InitCache()
	if err != nil {
		return err
	}

	// Initializing the repositories
	redisRepository := repository.NewRedisRepository(cache)

	citizenRepository := repository.NewCitizenRepository(database)
	roleRepository := repository.NewRoleRepository(database)

	// Initializing the usecases
	citizenUsecase := usecase.NewCitizenUsecase(citizenRepository, roleRepository, redisRepository)
	roleUsecase := usecase.NewRoleUsecase(roleRepository)

	// Initializing the routes
	s.Route = gin.New()

	// Initializing metric collector
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Adding middlewares
	m := middleware.InitMiddleware()
	s.Route.Use(gin.Recovery())
	s.Route.Use(m.CORSMiddleware())
	s.Route.Use(m.MetricMiddleware(metricService))

	s.Route.Any("/metrics", gin.WrapH(promhttp.Handler()))

	s.Route.Use(m.ErrorMiddleware())

	// Adding custom messages to field validations
	validation.RegisterCustomValidations()

	// Initializing the APIs route
	v1 := s.Route.Group("/api/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler.NewCitizenHandler(v1, citizenUsecase)
	handler.NewCitizenRoleHandler(v1, citizenUsecase)
	handler.NewRoleHandler(v1, roleUsecase)
	handler.NewRoleCitizenHandler(v1, roleUsecase)

	return nil
}

// StartServer ...
func (s *Server) StartServer() {
	if err := s.Initialize(); err != nil {
		s.Logger.LogIt("ERROR", fmt.Sprintf("[SERVER-START-ERROR] - Failed to start the server. %s", err.Error()), nil)
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", env.AppPort),
		ReadTimeout:  time.Duration(env.HTTPReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.HTTPWriteTimeout) * time.Second,
		Handler:      s.Route,
	}

	s.Logger.LogIt("INFO", fmt.Sprintf("[SERVER-START] - Starting server on port - %s", env.AppPort), nil)
	if err := httpServer.ListenAndServe(); err != nil {
		s.Logger.LogIt("ERROR", fmt.Sprintf("[SERVER-LISTEN-ERROR] - Failed at server to listen on port %s. %s", env.AppPort, err.Error()), nil)
		os.Exit(1)
	}
}
