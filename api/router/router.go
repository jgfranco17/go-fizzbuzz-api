package router

import (
	"fmt"
	"os"

	"api/context_settings"
	"api/env"
	"api/logger"
	"api/router/headers"
	system "api/router/system"
	v0 "api/router/v0"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	Router *gin.Engine
	Port   int
}

func (s *Service) Run() error {
	err := s.Router.Run(fmt.Sprintf(":%v", s.Port))
	if err != nil {
		return fmt.Errorf("Failed to start service on port %v: %w", s.Port, err)
	}
	return nil
}

// Add the fields we want to expose in the logger to the request context
func addLoggerFields() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !env.IsLocalEnvironment() {
			requestID := uuid.NewString()
			environment := os.Getenv(env.ENV_KEY_ENVIRONMENT)
			version := os.Getenv(env.ENV_KEY_VERSION)

			// Golang recommends contexts to use custom types instead
			// of strings, but gin defines key as a string.
			c.Set(string(context_settings.RequestId), requestID)
			c.Set(string(context_settings.Environment), environment)
			c.Set(string(context_settings.Version), version)

			originInfo, err := headers.CreateOriginInfoHeader(c)

			if err == nil && originInfo.Origin != "" {
				c.Set(string(context_settings.Origin), fmt.Sprintf("%s@%s", originInfo.Origin, originInfo.Version))
			}
		}
		c.Next()
	}
}

// Log the start and completion of a request
func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.FromContext(c)

		origin := c.Request.Header.Get("Origin")
		log.Infof("Request Started: [%s] %s from %s", c.Request.Method, c.Request.URL, origin)
		c.Next()
		log.Infof("Request Completed: [%s] %s", c.Request.Method, c.Request.URL)
	}
}

// Configure the router adding routes and middlewares
func getRouter() *gin.Engine {
	router := gin.Default()
	router.Use(addLoggerFields())
	router.Use(logRequest())
	router.Use(GetCors())
	router.Use(system.PrometheusMiddleware())
	system.SetSystemRoutes(router)
	v0.SetRoutes(router)

	return router
}

/*
Create a backend service instance.

[IN] port: server port to listen on

[OUT] *Service: new backend service instance
*/
func CreateService(port int) *Service {
	router := getRouter()
	return &Service{
		Router: router,
		Port:   port,
	}
}
