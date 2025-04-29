package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPortFromEnv(t *testing.T) {
	t.Setenv(envServicePort, "12345")
	expectedPort := 12345
	port := GetServicePort()
	assert.Equalf(t, expectedPort, port, "Wanted port %d but got %d", expectedPort, port)
}

func TestGetPortDefault_ValidPort(t *testing.T) {
	t.Setenv(envServicePort, "")
	port := GetServicePort()
	assert.Equalf(t, defaultPort, port, "Wanted port %d but got %d", defaultPort, port)
}

func TestGetPort_DefaultIfInvalid(t *testing.T) {
	t.Setenv(envServicePort, "abc")
	port := GetServicePort()
	assert.Equalf(t, defaultPort, port, "Wanted port %d but got %d", defaultPort, port)
}
