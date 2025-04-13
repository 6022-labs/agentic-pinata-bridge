package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/models"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"go.uber.org/zap"
)

type PinataClientInterface interface {
	PinByHash(hash string) (*models.ExternalPinByHashResponse, error)
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

func (p *PinataClient) PinByHash(hash string) (*models.ExternalPinByHashResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s%s", p.pinataSettings.BaseUrl, PinataPinCidEndpoint)

	// Create the JSON payload
	payload := map[string]string{
		"hashToPin": hash,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		p.logger.Error("Failed to marshal payload", zap.Error(err))
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		p.logger.Error("Failed to create HTTP request", zap.Error(err))
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+p.pinataSettings.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		p.logger.Error("Failed to send request to Pinata", zap.Error(err))
		return nil, fmt.Errorf("failed to send request to Pinata: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		responseBodyBytes, _ := io.ReadAll(resp.Body)
		p.logger.Error("Received non-200 response during decode",
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(responseBodyBytes)),
		)

		return nil, fmt.Errorf("received non-200 response during decode: %d", resp.StatusCode)
	}

	// Parse the response body
	var responseBody models.ExternalPinByHashResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		p.logger.Error("Failed to decode response body", zap.Error(err))
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &responseBody, nil
}
