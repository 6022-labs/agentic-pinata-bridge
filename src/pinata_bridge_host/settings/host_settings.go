package settings

import (
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const HostSettingsKey = "host"

type HostSettings struct {
	UseApi       bool `koanf:"use_api"`
	UseListeners bool `koanf:"use_listeners"`
	ApiPort      int  `koanf:"api_port"`
}

func NewHostSettings(logger *zap.Logger, k *koanf.Koanf) *HostSettings {
	settings := HostSettings{
		UseApi:       true,
		UseListeners: true,
		ApiPort:      3000,
	}
	if err := k.Unmarshal(HostSettingsKey, &settings); err != nil {
		logger.Fatal("failed to unmarshal host settings", zap.Error(err))
	}

	if settings.ApiPort <= 0 {
		logger.Warn("invalid host.api_port value, defaulting to 3000")
		settings.ApiPort = 3000
	}

	return &settings
}
