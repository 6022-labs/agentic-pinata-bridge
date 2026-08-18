package clients

import (
	"context"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"go.uber.org/zap"
)

type PinataClientInterface interface {
	PinByHash(ctx context.Context, request *models.ExternalPinByHashRequest) (*models.ExternalPinByHashResponse, error)
	QueryFileByCid(ctx context.Context, cid string) (*models.ExternalQueryFilesResponse, error)
}

const (
	PinByCidEndpoint   = "/pinning/pinByHash"
	QueryFilesEndpoint = "/v3/files/public"
)

type PinataClient struct {
	logger         *zap.Logger
	httpClient     *http.Client
	pinataSettings *settings.PinataSettings
}

func NewPinataClient(
	logger *zap.Logger,
	httpClient *http.Client,
	pinataSettings *settings.PinataSettings) *PinataClient {
	return &PinataClient{
		logger:         logger,
		httpClient:     httpClient,
		pinataSettings: pinataSettings,
	}
}

func (p *PinataClient) PinByHash(ctx context.Context, request *models.ExternalPinByHashRequest) (*models.ExternalPinByHashResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s%s", p.pinataSettings.BaseUrl, PinByCidEndpoint)

	bodyBytes, err := json.Marshal(request)
	if err != nil {
		p.logger.Error("Failed to marshal body", zap.Error(err))
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		p.logger.Error("Failed to create HTTP request", zap.Error(err))
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+p.pinataSettings.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := p.httpClient.Do(req)
	if err != nil {
		p.logger.Error("Failed to send request to Pinata", zap.Error(err))
		return nil, fmt.Errorf("failed to send request to Pinata: %w", err)
	}
	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		p.logger.Error("Failed to read response body", zap.Error(err))
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		p.logger.Error("Received non-200 response during decode",
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(responseBodyBytes)),
		)

		return nil, fmt.Errorf("received non-200 response during decode: %d", resp.StatusCode)
	}

	// Parse the response body
	var result models.ExternalPinByHashResponse
	if err := json.Unmarshal(responseBodyBytes, &result); err != nil {
		p.logger.Error("Failed to decode response body", zap.Error(err))
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}

func (p *PinataClient) QueryFileByCid(ctx context.Context, cid string) (*models.ExternalQueryFilesResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s%s?cid=%s", p.pinataSettings.BaseUrl, QueryFilesEndpoint, cid)

	// Create the HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		p.logger.Error("Failed to create HTTP request", zap.Error(err))
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+p.pinataSettings.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := p.httpClient.Do(req)
	if err != nil {
		p.logger.Error("Failed to send request to Pinata", zap.Error(err))
		return nil, fmt.Errorf("failed to send request to Pinata: %w", err)
	}
	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		p.logger.Error("Failed to read response body", zap.Error(err))
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		p.logger.Error("Received non-200 response during decode",
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(responseBodyBytes)),
		)

		return nil, fmt.Errorf("received non-200 response during decode: %d", resp.StatusCode)
	}

	// Parse the response body
	var result models.ExternalQueryFilesResponse
	if err := json.Unmarshal(responseBodyBytes, &result); err != nil {
		p.logger.Error("Failed to decode response body", zap.Error(err))
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}
