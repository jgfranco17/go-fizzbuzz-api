package v0

import (
	"api/handlers"
	error_handling "api/router/error_handling"

	"github.com/gin-gonic/gin"
)

// Adds v0 routes to the router.
func SetRoutes(route *gin.Engine) {
	v0 := route.Group("/v0")
	{
		v0.GET("fizzbuzz/:number", error_handling.WithErrorHandling(handlers.FizzbuzzHandler()))
	}
}
