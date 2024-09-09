package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
)

type Validators struct {
	Consensus []string `json:"consensus"`
	Proxy     []string `json:"proxy"`
}

type Constraints struct {
	Slot uint64
	Top  []string
	Rest []string
}

// MevBoostClient represents a client for interacting with the MEV Boost server.
type MevBoostClient struct {
	baseURL string
	client  *http.Client
}

// NewMevBoostClient initializes a new MEV Boost client with the specified base URL.
func NewMevBoostClient(baseURL string, timeout time.Duration) (*MevBoostClient, error) {
	client := &MevBoostClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}

	// Check if the client can connect to the MEV Boost server.
	if _, err := client.GetValidatorPubKeyHex(); err != nil {
		return nil, fmt.Errorf("failed to connect to MEV Boost server: %w", err)
	}

	return client, nil
}

func (c *MevBoostClient) SetConstraints(slot uint64, txs types.Transactions) error {
	url := fmt.Sprintf("%s/v1/constraints", c.baseURL)

	// TODO: Implement constraints
	resp, err := c.client.Post(url, "application/json", nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
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
