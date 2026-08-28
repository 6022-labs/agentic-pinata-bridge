package settings

import (
	"strings"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const PinataSettingsKey = "pinata"

// defaultPinataBaseUrl is the server the generated v3 client is built against.
const defaultPinataBaseUrl = "https://api.pinata.cloud/v3"

type PinataSettings struct {
	ApiKey  string `koanf:"api_key"`
	BaseUrl string `koanf:"base_url"`
}

func NewPinataSettings(logger *zap.Logger, k *koanf.Koanf) *PinataSettings {
	settings := PinataSettings{BaseUrl: defaultPinataBaseUrl}
	if err := k.Unmarshal(PinataSettingsKey, &settings); err != nil {
		logger.Fatal("failed to unmarshal pinata settings", zap.Error(err))
	}

	settings.ApiKey = strings.TrimSpace(settings.ApiKey)
	if settings.ApiKey == "" {
		logger.Fatal("please set pinata.api_key (or PINATA__API_KEY env)")
	}

	settings.BaseUrl = strings.TrimSpace(settings.BaseUrl)
	if settings.BaseUrl == "" {
		logger.Warn("empty pinata.base_url, using default", zap.String("default", defaultPinataBaseUrl))
		settings.BaseUrl = defaultPinataBaseUrl
	}

	return &settings
}
