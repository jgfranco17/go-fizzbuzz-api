package main

import (
	"context"
	"flag"

	"api/env"
	"api/logging"
	"api/router"
	"api/router/system"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	port    = flag.Int("port", 8080, "Port to listen on")
	devMode = flag.Bool("dev", true, "Run server in debug mode")
)

func init() {
	logrus.SetReportCaller(true)

	if env.IsLocalEnvironment() {
		logrus.SetFormatter(&logrus.TextFormatter{})
		gin.SetMode(gin.DebugMode)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	}
	prometheus.Register(system.HttpLastRequestReceivedTime)
}

func main() {
	flag.Parse()
	ctx := context.WithValue(context.Background(), "section", "Main")
	log := logging.GetLogger(ctx)
	if *devMode {
		log.Infof("Running API server on port %d in dev mode", *port)
	} else {
		log.Infof("Running API production server on port %d", *port)
		gin.SetMode(gin.ReleaseMode)
	}
	service := router.CreateService(*port)
	err := service.Run()
	if err != nil {
		log.Error("Error starting the server:", err)
	}
}
