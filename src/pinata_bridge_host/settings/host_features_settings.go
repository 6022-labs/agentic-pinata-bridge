package settings

import (
	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const HostFeaturesSettingsKey = "host"

// HostFeaturesSettings switches this host's two roles; the generic server settings live in common.
type HostFeaturesSettings struct {
	UseApi       bool `koanf:"use_api"`
	UseListeners bool `koanf:"use_listeners"`
}

func NewHostFeaturesSettings(logger *zap.Logger, k *koanf.Koanf) *HostFeaturesSettings {
	settings := HostFeaturesSettings{
		UseApi:       true,
		UseListeners: true,
	}
	if err := k.Unmarshal(HostFeaturesSettingsKey, &settings); err != nil {
		logger.Fatal("failed to unmarshal host settings", zap.Error(err))
	}

	return &settings
}
