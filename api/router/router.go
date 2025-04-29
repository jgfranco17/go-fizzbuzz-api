package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type ServiceEndpoint struct {
	Method         string
	Path           string
	HandleFunction http.HandlerFunc
}

func (e *ServiceEndpoint) AddToRouter(r *chi.Mux) {
	r.MethodFunc(e.Method, e.Path, e.HandleFunction)
}

func NewRestService(port int) *chi.Mux {
	endpoints := []ServiceEndpoint{
		{
			Method:         "GET",
			Path:           "/",
			HandleFunction: GetHomeRoute(),
		},
		{
			Method:         "GET",
			Path:           "/fizzbuzz",
			HandleFunction: GetFizzbuzzRoute(),
		},
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	for _, endpoint := range endpoints {
		endpoint.AddToRouter(router)
	}
	return router
}
