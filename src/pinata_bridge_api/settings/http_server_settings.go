package settings

import (
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type HttpServerSettings struct {
	Port    int
	AppName string
}

func NewHttpServerSettings(logger *zap.Logger) *HttpServerSettings {
	portStr := os.Getenv("PORT")
	if len(strings.TrimSpace(portStr)) == 0 {
		logger.Warn("environment variable PORT is not set, defaulting to 3000")
		portStr = "3000"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Warn("invalid PORT environment variable value, defaulting to 3000")
		port = 3000
	}

	appName := os.Getenv("APP_NAME")
	if len(strings.TrimSpace(appName)) == 0 {
		logger.Fatal("please set your APP_NAME value in your environment")
	}

	return &HttpServerSettings{
		Port:    port,
		AppName: appName,
	}
}
