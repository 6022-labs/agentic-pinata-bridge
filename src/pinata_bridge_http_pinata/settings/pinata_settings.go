package settings

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

type PinataSettings struct {
	ApiKey  string
	BaseUrl string
}

func NewPinataSettings(logger *zap.Logger) *PinataSettings {
	pinataApiKey := os.Getenv("PINATA_API_KEY")
	if len(strings.TrimSpace(pinataApiKey)) == 0 {
		logger.Fatal("please set your PINATA_API_KEY value in your environment")
	}

	pinataBaseUrl := os.Getenv("PINATA_BASE_URL")
	if len(strings.TrimSpace(pinataBaseUrl)) == 0 {
		logger.Fatal("please set your PINATA_BASE_URL value in your environment")
	}

	return &PinataSettings{
		ApiKey:  pinataApiKey,
		BaseUrl: pinataBaseUrl,
	}
}
