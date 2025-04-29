package environment

import (
	"os"
	"strconv"
)

var (
	defaultPort    int    = 8080
	envServicePort string = "FIZZBUZZ_SERVICE_PORT"
)

func GetServicePort() int {
	portInEnv := os.Getenv(envServicePort)
	if portInEnv == "" {
		return defaultPort
	}
	port, err := strconv.Atoi(portInEnv)
	if err != nil {
		return defaultPort
	}
	return port
}
