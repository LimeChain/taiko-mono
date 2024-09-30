package types

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
