package settings

import (
	"strings"

	"github.com/knadh/koanf/v2"
	"go.uber.org/zap"
)

const IpfsCheckSettingsKey = "ipfs_check"

type IpfsCheckSettings struct {
	BaseUrl string `koanf:"base_url"`
}

func NewIpfsCheckSettings(logger *zap.Logger, k *koanf.Koanf) *IpfsCheckSettings {
	settings := IpfsCheckSettings{}
	if err := k.Unmarshal(IpfsCheckSettingsKey, &settings); err != nil {
		logger.Fatal("failed to unmarshal ipfs_check settings", zap.Error(err))
	}

	settings.BaseUrl = strings.TrimSpace(settings.BaseUrl)
	if settings.BaseUrl == "" {
		logger.Fatal("please set ipfs_check.base_url (or IPFS_CHECK__BASE_URL env)")
	}

	return &settings
}
