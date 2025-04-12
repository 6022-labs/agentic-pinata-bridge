package clients

import (
	"fmt"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"go.uber.org/zap"
)

type PinataClientInterface interface {
	PinCidToPinata(cid string) (*string, error)
}

const (
	PinataPinCidEndpoint = "/pinning/pinByHash"
)

type PinataClient struct {
	logger         *zap.Logger
	pinataSettings *settings.PinataSettings
}

func NewPinataClient(logger *zap.Logger, pinataSettings *settings.PinataSettings) *PinataClient {
	return &PinataClient{
		logger:         logger,
		pinataSettings: pinataSettings,
	}
}

func (p *PinataClient) PinCidToPinata(cid string) (*string, error) {
	return nil, fmt.Errorf("pinning to Pinata is not implemented yet")
}
