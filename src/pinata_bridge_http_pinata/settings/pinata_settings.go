package settings

import (
	"strings"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const PinataSettingsKey = "pinata"

type PinataSettings struct {
	ApiKey  string `koanf:"api_key"`
	BaseUrl string `koanf:"base_url"`
}

func NewPinataSettings(logger *zap.Logger, k *koanf.Koanf) *PinataSettings {
	settings := PinataSettings{}
	if err := k.Unmarshal(PinataSettingsKey, &settings); err != nil {
		logger.Fatal("failed to unmarshal pinata settings", zap.Error(err))
	}

	settings.ApiKey = strings.TrimSpace(settings.ApiKey)
	if settings.ApiKey == "" {
		logger.Fatal("please set pinata.api_key (or PINATA__API_KEY env)")
	}

	settings.BaseUrl = strings.TrimSpace(settings.BaseUrl)
	if settings.BaseUrl == "" {
		logger.Fatal("please set pinata.base_url (or PINATA__BASE_URL env)")
	}

	return &settings
}
