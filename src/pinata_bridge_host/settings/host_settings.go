package settings

import (
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type HostSettings struct {
	UseApi       bool
	UseListeners bool
	ApiPort      int
	AppName      string
}

func NewHostSettings(logger *zap.Logger) *HostSettings {
	useApiStr := os.Getenv("USE_API")
	if len(strings.TrimSpace(useApiStr)) == 0 {
		logger.Warn("environment variable USE_API is not set, defaulting to true")
		useApiStr = "true"
	}

	useApi := strings.ToLower(useApiStr) == "true" || useApiStr == "1"

	useListenersStr := os.Getenv("USE_LISTENERS")
	if len(strings.TrimSpace(useListenersStr)) == 0 {
		logger.Warn("environment variable USE_LISTENERS is not set, defaulting to true")
		useListenersStr = "true"
	}

	useListeners := strings.ToLower(useListenersStr) == "true" || useListenersStr == "1"

	portStr := os.Getenv("API_PORT")
	if len(strings.TrimSpace(portStr)) == 0 {
		logger.Warn("environment variable API_PORT is not set, defaulting to 3000")
		portStr = "3000"
	}

	apiPort, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Warn("invalid API_PORT environment variable value, defaulting to 3000")
		apiPort = 3000
	}

	appName := os.Getenv("APP_NAME")
	if len(strings.TrimSpace(appName)) == 0 {
		logger.Fatal("please set your APP_NAME value in your environment")
	}

	return &HostSettings{
		ApiPort: apiPort,
		AppName: appName,

		UseApi:       useApi,
		UseListeners: useListeners,
	}
}
