package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jgfranco/fizzbuzz-go/api/calculation"
)

func GetHomeRoute() http.HandlerFunc {
	type simpleResponse struct {
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		greeting := simpleResponse{
			Message: "Welcome to Fizzbuzz API (Go edition)",
		}
		err := json.NewEncoder(w).Encode(greeting)
		if err != nil {
			http.Error(w, "Failed to serve landing page", http.StatusInternalServerError)
			return
		}
	}
}

func GetHealthCheckRoute() http.HandlerFunc {
	type healthCheck struct {
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		greeting := healthCheck{
			Status: "healthy",
		}
		err := json.NewEncoder(w).Encode(greeting)
		if err != nil {
			http.Error(w, "Failed to serve health-check", http.StatusInternalServerError)
			return
		}
	}
}

func GetFizzbuzzRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		numberParam := r.URL.Query().Get("number")
		if numberParam == "" {
			http.Error(w, "No number parameter provided", http.StatusBadRequest)
			return
		}
		number, err := strconv.Atoi(numberParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid number parameter provided: got %v", numberParam), http.StatusBadRequest)
			return
		}
		sequence, err := calculation.GenerateFizzbuzzSequence(number)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate sequence: %s", err.Error()), http.StatusBadRequest)
			return
		}
		if err := json.NewEncoder(w).Encode(sequence); err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	}
}
