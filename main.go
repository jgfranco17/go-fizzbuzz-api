package main

import (
	"fmt"
	"net/http"

	"github.com/jgfranco/fizzbuzz-go/api/environment"
	"github.com/jgfranco/fizzbuzz-go/api/logging"
	"github.com/jgfranco/fizzbuzz-go/api/router"
)

func main() {
	logger := logging.NewLogger()
	port := environment.GetServicePort()
	r := router.NewRestService(port)
	logger.Infof("Starting service!")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		logger.Fatalf("Error while running service: %v", err)
	}
}
