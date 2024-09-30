package rpc

import (
	"bytes"
	"encoding/hex"
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
	Top        []string `json:"top"`
	Rest       []string `json:"rest"`
	SlotNumber uint64   `json:"slot_number"`
}

type IMevBoostClient interface {
	SetConstraints(slot uint64, tx *types.Transaction) error
	GetValidatorPubKeyHex() (string, error)
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

func (c *MevBoostClient) SetConstraints(slot uint64, tx *types.Transaction) error {
	url := fmt.Sprintf("%s/v1/constraints", c.baseURL)

	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return err
	}

	if err != nil {
		return fmt.Errorf("failed to get transaction signature: %w", err)
	}

	hexStr := hex.EncodeToString(txBytes)

	constraints := Constraints{
		Top:        []string{hexStr},
		Rest:       []string{},
		SlotNumber: slot,
	}

	body, err := json.Marshal(constraints)

	if err != nil {
		return fmt.Errorf("failed to marshal constraints: %w", err)
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

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
