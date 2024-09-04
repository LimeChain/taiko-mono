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
)

var (
	// Request urls.
	sidecarsRequestURL       = "/eth/v1/beacon/blob_sidecars/%d"
	proposerDutiesRequestURL = "/eth/v1/validator/duties/proposer/%d"
	genesisRequestURL        = "/eth/v1/beacon/genesis"
	getConfigSpecPath        = "/eth/v1/config/spec"
	getNodeSyncingPath       = "/eth/v1/node/syncing"
)

type ConfigSpec struct {
	SecondsPerSlot string `json:"SECONDS_PER_SLOT"`
	SlotsPerEpoch  string `json:"SLOTS_PER_EPOCH"`
}

type GenesisResponse struct {
	Data struct {
		GenesisTime string `json:"genesis_time"`
	} `json:"data"`
}

type GetNodeSyncingResponse struct {
	Data struct {
		HeadSlot string `json:"head_slot"`
	} `json:"data"`
}

type ProposerDutiesResponse struct {
	Data []*ProposerDuty `json:"data"`
}

type ProposerDuty struct {
	PubKey         string `json:"pubkey"`
	ValidatorIndex string `json:"validator_index"`
	Slot           string `json:"slot"`
}

type BeaconClient struct {
	*beacon.Client

	timeout        time.Duration
	genesisTime    uint64
	secondsPerSlot uint64
	slotsPerEpoch  uint64
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
	var genesisDetail *GenesisResponse
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

func (c *BeaconClient) GetNextProposerDuties(ctx context.Context, headSlot uint64, maxSlots uint64) ([]*ProposerDuty, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, c.timeout)
	nextDuties := make([]*ProposerDuty, 0)
	defer cancel()

	epochBegin := headSlot / c.slotsPerEpoch
	epochEnd := (headSlot + maxSlots - 1) / c.slotsPerEpoch

	log.Info("Requesting proposer duties", "epochBegin", epochBegin, "epochEnd", epochEnd)

	resBytes, err := c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(proposerDutiesRequestURL, epochBegin))
	if err != nil {
		return nil, err
	}

	var duties *ProposerDutiesResponse
	if err = json.Unmarshal(resBytes, &duties); err != nil {
		return nil, err
	}

	nextDuties = append(nextDuties, duties.Data...)

	if epochEnd != epochBegin {
		resBytes, err = c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(proposerDutiesRequestURL, epochEnd))
		if err != nil {
			log.Error("Requesting duties for epochEnd", "epochEnd", epochEnd, "error", err)
		} else {
			var duties *ProposerDutiesResponse
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

func (c *BeaconClient) filterProposerDuties(duties []*ProposerDuty, headSlot uint64, maxSlots uint64) []*ProposerDuty {
	filteredDuties := make([]*ProposerDuty, 0)
	for _, duty := range duties {
		slot, err := strconv.Atoi(duty.Slot)
		log.Trace("Filtering duty", "slot", slot, "headSlot", headSlot, "maxSlots", maxSlots)
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

// GetHeadSlot returns the current head slot.
func (c *BeaconClient) GetHeadSlot(ctx context.Context) (*uint64, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, c.timeout)
	defer cancel()

	resBytes, err := c.Get(ctxWithTimeout, c.BaseURL().Path+getNodeSyncingPath)
	if err != nil {
		return nil, errors.Wrap(err, "error requesting configSpecPath")
	}

	nodeSyncing := &GetNodeSyncingResponse{}
	if err := json.Unmarshal(resBytes, &nodeSyncing); err != nil {
		return nil, err
	}

	headSlot, err := strconv.Atoi(nodeSyncing.Data.HeadSlot)
	if err != nil {
		return nil, err
	}

	currentSlot := uint64(headSlot)

	return &currentSlot, nil
}
