package settings

import (
	"strings"
	"time"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const PinataSettingsKey = "pinata"

// defaultPinataBaseUrl is the server the generated v3 client is built against.
const defaultPinataBaseUrl = "https://api.pinata.cloud/v3"

const defaultPinataRequestTimeoutSeconds = 30

type PinataSettings struct {
	ApiKey         string
	BaseUrl        string
	RequestTimeout time.Duration
}

type pinataSettingsRaw struct {
	ApiKey                string `koanf:"api_key"`
	BaseUrl               string `koanf:"base_url"`
	RequestTimeoutSeconds int    `koanf:"request_timeout_seconds"`
}

func NewPinataSettings(logger *zap.Logger, k *koanf.Koanf) *PinataSettings {
	raw := pinataSettingsRaw{
		BaseUrl:               defaultPinataBaseUrl,
		RequestTimeoutSeconds: defaultPinataRequestTimeoutSeconds,
	}
	if err := k.Unmarshal(PinataSettingsKey, &raw); err != nil {
		logger.Fatal("failed to unmarshal pinata settings", zap.Error(err))
	}

	if raw.RequestTimeoutSeconds <= 0 {
		logger.Warn(
			"non-positive pinata.request_timeout_seconds, using default",
			zap.Int("default", defaultPinataRequestTimeoutSeconds),
		)
		raw.RequestTimeoutSeconds = defaultPinataRequestTimeoutSeconds
	}

	settings := PinataSettings{
		ApiKey:         strings.TrimSpace(raw.ApiKey),
		BaseUrl:        strings.TrimSpace(raw.BaseUrl),
		RequestTimeout: time.Duration(raw.RequestTimeoutSeconds) * time.Second,
	}

	if settings.ApiKey == "" {
		logger.Fatal("please set pinata.api_key (or PINATA__API_KEY env)")
	}

	if settings.BaseUrl == "" {
		logger.Warn("empty pinata.base_url, using default", zap.String("default", defaultPinataBaseUrl))
		settings.BaseUrl = defaultPinataBaseUrl
	}

	return &settings
}
