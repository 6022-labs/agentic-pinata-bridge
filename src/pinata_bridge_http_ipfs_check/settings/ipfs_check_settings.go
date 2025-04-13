package settings

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

type IpfsCheckSettings struct {
	BaseUrl string `json:"BaseUrl"`
}

func NewIpfsCheckSettings(logger *zap.Logger) *IpfsCheckSettings {
	ipfsCheckBaseUrl := os.Getenv("IPFS_CHECK_BASE_URL")
	if len(strings.TrimSpace(ipfsCheckBaseUrl)) == 0 {
		logger.Fatal("IPFS_CHECK_BASE_URL is not set")
	}

	return &IpfsCheckSettings{
		BaseUrl: ipfsCheckBaseUrl,
	}
}
