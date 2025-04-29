package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

// List of keys needed by the core functionality
const (
	RequestId   string = "requestId"
	Version     string = "version"
	Environment string = "environment"
	Origin      string = "origin"
)

// Returns an instance of the logger, adding the fields found in the context.
func GetLoggerFromContext(ctx context.Context) *logrus.Entry {
	entry := logrus.WithFields(logrus.Fields{})
	if ctx == nil {
		return entry
	}
	fields := []string{
		RequestId,
		Version,
		Environment,
	}

	// Add the fields to the logger
	for _, field := range fields {
		value := ctx.Value(field)
		if value != nil {
			entry = entry.WithField(string(field), value)
		}
	}

	return entry
}

// NewLogger configures and registers a new logger instance.
func NewLogger(isLocal bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	if isLocal {
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	return logger
}
