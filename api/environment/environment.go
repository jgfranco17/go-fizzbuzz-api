package environment

import (
	"os"
	"strconv"

	"github.com/jgfranco/fizzbuzz-go/api/logging"
)

var (
	defaultPort    int    = 8000
	envServicePort string = "FIZZBUZZ_SERVICE_PORT"
)

func GetServicePort() int {
	logger := logging.NewLogger()
	portInEnv := os.Getenv(envServicePort)
	if portInEnv == "" {
		logger.Warnf("Env %s was not set, using default port %d", envServicePort, defaultPort)
		return defaultPort
	}
	port, err := strconv.Atoi(portInEnv)
	if err != nil {
		logger.Warnf("Env %s had invalid value '%s', using default port %d", envServicePort, portInEnv, defaultPort)
		return defaultPort
	}
	logger.Infof("Using port %d", port)
	return port
}
