package clients

import (
	"context"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/settings"
	"go.uber.org/zap"
)

const (
	CHECK_ENDPOINT = "/check"
)

type IpfsCheckClientInterface interface {
	Check(ctx context.Context, cid string) ([]models.ExternalCheckResponse, error)
}

type IpfsCheckClient struct {
	logger            *zap.Logger
	httpClient        *http.Client
	ipfsCheckSettings *settings.IpfsCheckSettings
}

func NewIpfsCheckClient(
	logger *zap.Logger,
	httpClient *http.Client,
	ipfsCheckSettings *settings.IpfsCheckSettings,
) *IpfsCheckClient {
	return &IpfsCheckClient{
		logger:            logger,
		httpClient:        httpClient,
		ipfsCheckSettings: ipfsCheckSettings,
	}
}

func (client *IpfsCheckClient) Check(ctx context.Context, cid string) ([]models.ExternalCheckResponse, error) {
	client.logger.Info("Starting Check method", zap.String("cid", cid))

	url := fmt.Sprintf("%s%s?cid=%s", client.ipfsCheckSettings.BaseUrl, CHECK_ENDPOINT, cid)
	client.logger.Debug("Constructed URL", zap.String("url", url))

	// Create a new POST request
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(nil))
	if err != nil {
		client.logger.Error("Failed to create request", zap.Error(err))
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers if needed (e.g., Content-Type)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		client.logger.Error("Request failed", zap.Error(err))
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		client.logger.Error("Failed to read response body", zap.Error(err))
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		client.logger.Error(
			"Unexpected status code",
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(body)),
		)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body into the expected structure
	var result []models.ExternalCheckResponse
	if err := json.Unmarshal(body, &result); err != nil {
		client.logger.Error("Failed to unmarshal response", zap.Error(err))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	client.logger.Info("Check method completed successfully", zap.Int("response_count", len(result)))
	return result, nil
}
