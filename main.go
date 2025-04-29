package main

import (
	"fmt"
	"net/http"

	"github.com/jgfranco/fizzbuzz-go/api/environment"
	"github.com/jgfranco/fizzbuzz-go/api/router"
)

func main() {
	port := environment.GetServicePort()
	r := router.NewRestService(port)
	fmt.Println("Starting service! Port:", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
