package settings

import (
	"strings"
	"time"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const IpfsCheckSettingsKey = "ipfs_check"

// defaultIpfsCheckRequestTimeoutSeconds outlives ipfs-check's own 60-second limit on a check.
const defaultIpfsCheckRequestTimeoutSeconds = 90

type IpfsCheckSettings struct {
	BaseUrl        string
	RequestTimeout time.Duration
}

type ipfsCheckSettingsRaw struct {
	BaseUrl               string `koanf:"base_url"`
	RequestTimeoutSeconds int    `koanf:"request_timeout_seconds"`
}

func NewIpfsCheckSettings(logger *zap.Logger, k *koanf.Koanf) *IpfsCheckSettings {
	raw := ipfsCheckSettingsRaw{RequestTimeoutSeconds: defaultIpfsCheckRequestTimeoutSeconds}
	if err := k.Unmarshal(IpfsCheckSettingsKey, &raw); err != nil {
		logger.Fatal("failed to unmarshal ipfs_check settings", zap.Error(err))
	}

	if raw.RequestTimeoutSeconds <= 0 {
		logger.Warn(
			"non-positive ipfs_check.request_timeout_seconds, using default",
			zap.Int("default", defaultIpfsCheckRequestTimeoutSeconds),
		)
		raw.RequestTimeoutSeconds = defaultIpfsCheckRequestTimeoutSeconds
	}

	settings := IpfsCheckSettings{
		BaseUrl:        strings.TrimSpace(raw.BaseUrl),
		RequestTimeout: time.Duration(raw.RequestTimeoutSeconds) * time.Second,
	}

	if settings.BaseUrl == "" {
		logger.Fatal("please set ipfs_check.base_url (or IPFS_CHECK__BASE_URL env)")
	}

	return &settings
}
