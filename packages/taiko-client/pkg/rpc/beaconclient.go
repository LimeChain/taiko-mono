package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/api/client"
	"github.com/prysmaticlabs/prysm/v4/api/client/beacon"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/blob"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/config"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/types"
)

var (
	// Request urls.
	sidecarsRequestURL       = "/eth/v1/beacon/blob_sidecars/%d"
	proposerDutiesRequestURL = "/eth/v1/validator/duties/proposer/%d"
	genesisRequestURL        = "/eth/v1/beacon/genesis"
	getConfigSpecPath        = "/eth/v1/config/spec"
)

type BeaconClient struct {
	*beacon.Client

	timeout        time.Duration
	genesisTime    uint64
	secondsPerSlot uint64
	slotsPerEpoch  uint64
}

type IBeaconClient interface {
	GetBlobs(ctx context.Context, time uint64) ([]*blob.Sidecar, error)
	GetNextProposerDuties(ctx context.Context, headSlot uint64, maxSlots uint64) ([]*types.ProposerDuty, error)
	GetL1HeadSlot() uint64
	GetTimestampBySlot(slot uint64) uint64
	GetGenesisTimestamp() uint64
	GetSecondsPerSlot() uint64
	GetSlotsPerEpoch() uint64
}

// NewBeaconClient returns a new beacon client.
func NewBeaconClient(endpoint string, timeout time.Duration) (*BeaconClient, error) {
	cli, err := beacon.NewClient(strings.TrimSuffix(endpoint, "/"), client.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Get the genesis time.
	var genesisDetail *types.GenesisResponse
	resBytes, err := cli.Get(ctx, cli.BaseURL().Path+genesisRequestURL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBytes, &genesisDetail); err != nil {
		return nil, err
	}

	genesisTime, err := strconv.Atoi(genesisDetail.Data.GenesisTime)
	if err != nil {
		return nil, err
	}

	log.Info("L1 genesis time", "time", genesisTime)

	// Get the seconds per slot.
	spec, err := getConfigSpec(ctx, cli)
	if err != nil {
		return nil, err
	}

	secondsPerSlot, err := strconv.Atoi(spec.Data.(map[string]interface{})["SECONDS_PER_SLOT"].(string))
	if err != nil {
		return nil, err
	}

	slotsPerEpoch, err := strconv.Atoi(spec.Data.(map[string]interface{})["SLOTS_PER_EPOCH"].(string))
	if err != nil {
		return nil, err
	}

	log.Info("L1 seconds per slot", "seconds", secondsPerSlot)
	log.Info("L1 slots per epoch", "slots", slotsPerEpoch)

	return &BeaconClient{cli, timeout, uint64(genesisTime), uint64(secondsPerSlot), uint64(slotsPerEpoch)}, nil
}

// GetBlobs returns the sidecars for a given slot.
func (c *BeaconClient) GetBlobs(ctx context.Context, time uint64) ([]*blob.Sidecar, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, c.timeout)
	defer cancel()

	slot, err := c.timeToSlot(time)
	if err != nil {
		return nil, err
	}

	resBytes, err := c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(sidecarsRequestURL, slot))
	if err != nil {
		return nil, err
	}

	var sidecars *blob.SidecarsResponse
	if err = json.Unmarshal(resBytes, &sidecars); err != nil {
		return nil, err
	}

	return sidecars.Data, nil
}

// timeToSlot returns the slots of the given timestamp.
func (c *BeaconClient) timeToSlot(timestamp uint64) (uint64, error) {
	if timestamp < c.genesisTime {
		return 0, fmt.Errorf("provided timestamp (%v) precedes genesis time (%v)", timestamp, c.genesisTime)
	}
	return (timestamp - c.genesisTime) / c.secondsPerSlot, nil
}

// getConfigSpec retrieve the current configs of the network used by the beacon node.
func getConfigSpec(ctx context.Context, c *beacon.Client) (*config.GetSpecResponse, error) {
	body, err := c.Get(ctx, c.BaseURL().Path+getConfigSpecPath)
	if err != nil {
		return nil, errors.Wrap(err, "error requesting configSpecPath")
	}
	fsr := &config.GetSpecResponse{}
	err = json.Unmarshal(body, fsr)
	if err != nil {
		return nil, err
	}
	return fsr, nil
}

func (c *BeaconClient) GetNextProposerDuties(
	ctx context.Context,
	headSlot uint64,
	maxSlots uint64) ([]*types.ProposerDuty, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, c.timeout)
	nextDuties := make([]*types.ProposerDuty, 0)
	defer cancel()

	epochBegin := headSlot / c.slotsPerEpoch
	epochEnd := (headSlot + maxSlots - 1) / c.slotsPerEpoch

	log.Info("Requesting proposer duties", "epochBegin", epochBegin, "epochEnd", epochEnd)

	resBytes, err := c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(proposerDutiesRequestURL, epochBegin))
	if err != nil {
		return nil, err
	}

	var duties *types.ProposerDutiesResponse
	if err = json.Unmarshal(resBytes, &duties); err != nil {
		return nil, err
	}

	nextDuties = append(nextDuties, duties.Data...)

	if epochEnd != epochBegin {
		resBytes, err = c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(proposerDutiesRequestURL, epochEnd))
		if err != nil {
			log.Error("Requesting duties for epochEnd", "epochEnd", epochEnd, "error", err)
		} else {
			var duties *types.ProposerDutiesResponse
			if err = json.Unmarshal(resBytes, &duties); err != nil {
				return nil, err
			}
			nextDuties = append(nextDuties, duties.Data...)
		}
	}

	nextDuties = c.filterProposerDuties(nextDuties, headSlot, maxSlots)

	log.Info("Received proposer duties", "duties", len(nextDuties))

	return nextDuties, nil
}

func (c *BeaconClient) filterProposerDuties(
	duties []*types.ProposerDuty,
	headSlot uint64,
	maxSlots uint64,
) []*types.ProposerDuty {
	filteredDuties := make([]*types.ProposerDuty, 0)
	for _, duty := range duties {
		slot, err := strconv.Atoi(duty.Slot)
		if err != nil {
			log.Error("Failed to convert slot to integer", "slot", duty.Slot, "error", err)
			continue
		}
		if uint64(slot) >= headSlot && uint64(slot) < (headSlot+maxSlots) {
			filteredDuties = append(filteredDuties, duty)
		}
	}
	return filteredDuties
}

func (c *BeaconClient) GetL1HeadSlot() uint64 {
	now := time.Now().Unix()
	elapsedTime := uint64(now) - c.genesisTime
	l1HeadSlot := elapsedTime / c.secondsPerSlot
	return l1HeadSlot
}

func (c *BeaconClient) GetTimestampBySlot(slot uint64) uint64 {
	return c.genesisTime + slot*c.secondsPerSlot
}

func (c *BeaconClient) GetGenesisTimestamp() uint64 {
	return c.genesisTime
}

func (c *BeaconClient) GetSecondsPerSlot() uint64 {
	return c.secondsPerSlot
}

func (c *BeaconClient) GetSlotsPerEpoch() uint64 {
	return c.slotsPerEpoch
}
