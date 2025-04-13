package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_ipfs_check/models"
	"github.com/6022protocol/agentic-ai-pinata-bridge/src/pinata_bridge_http_ipfs_check/settings"
)

const (
	CHECK_ENDPOINT = "/check"
)

type IpfsCheckClientInterface interface {
	Check(cid string) ([]models.ExternalCheckResponse, error)
}

type IpfsCheckClient struct {
	ipfsCheckSettings *settings.IpfsCheckSettings
}

func NewIpfsCheckClient(ipfsCheckSettings *settings.IpfsCheckSettings) *IpfsCheckClient {
	return &IpfsCheckClient{
		ipfsCheckSettings: ipfsCheckSettings,
	}
}

func (client *IpfsCheckClient) Check(cid string) ([]models.ExternalCheckResponse, error) {
	url := fmt.Sprintf("%s%s?cid=%s", client.ipfsCheckSettings.BaseUrl, CHECK_ENDPOINT, cid)

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers if needed (e.g., Content-Type)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body into the expected structure
	var result []models.ExternalCheckResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return result, nil
}
