package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/http"
)

type Validators struct {
	Consensus []string `json:"consensus"`
	Proxy     []string `json:"proxy"`
}

// MevBoostClient represents a client for interacting with the MEV Boost server.
type MevBoostClient struct {
	baseURL string
	client  *http.Client
}

// NewMevBoostClient initializes a new MEV Boost client with the specified base URL.
func NewMevBoostClient(baseURL string, timeout time.Duration) *MevBoostClient {
	return &MevBoostClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// GetValidators fetches the validators from the /v1/validators endpoint.
func (c *MevBoostClient) GetValidatorPubKeyHex() (string, error) {
	url := fmt.Sprintf("%s/v1/validators", c.baseURL)

	resp, err := c.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var validators Validators
	err = json.Unmarshal(body, &validators)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(validators.Consensus) == 0 {
		return "", fmt.Errorf("no consensus validators found")
	}

	return validators.Consensus[0], nil
}
