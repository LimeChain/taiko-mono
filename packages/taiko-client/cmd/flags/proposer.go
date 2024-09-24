package flags

import (
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/version"
)

// Required flags used by proposer.
var (
	L1ProposerPrivKey = &cli.StringFlag{
		Name:     "l1.proposerPrivKey",
		Usage:    "Private key of the L1 proposer, who will send TaikoL1.proposeBlock transactions",
		Required: true,
		Category: proposerCategory,
		EnvVars:  []string{"L1_PROPOSER_PRIV_KEY"},
	}
	L1MevBoostEndpoint = &cli.StringFlag{
		Name:     "l1.mevBoost",
		Usage:    "HTTP endpoint of a L1 mev-boost service",
		Required: true,
		Category: proposerCategory,
		EnvVars:  []string{"L1_MEV_BOOST"},
	}
	ProverEndpoints = &cli.StringFlag{
		Name:     "proverEndpoints",
		Usage:    "Comma-delineated list of prover endpoints proposer should query when attempting to propose a block",
		Required: true,
		Category: proposerCategory,
		EnvVars:  []string{"PROVER_ENDPOINTS"},
	}
	PreconfDelay = &cli.DurationFlag{
		Name:     "epoch.interval",
		Usage:    "Time delay to send L2 preconf transaction",
		Category: proposerCategory,
		Value:    8,
		EnvVars:  []string{"PRECONF_DELAY"},
	}
	L2SuggestedFeeRecipient = &cli.StringFlag{
		Name:     "l2.suggestedFeeRecipient",
		Usage:    "Address of the proposed block's suggested L2 fee recipient",
		Required: true,
		Category: proposerCategory,
		EnvVars:  []string{"L2_SUGGESTED_FEE_RECIPIENT"},
	}
	SequencerRegistryAddress = &cli.StringFlag{
		Name:     "sequencerRegistry",
		Usage:    "SequencerRegistry contract `address`",
		Required: true,
		Category: commonCategory,
		EnvVars:  []string{"SEQUENCER_REGISTRY"},
	}
)

// Optional flags used by proposer.
var (
	// Tier fee related.
	OptimisticTierFee = &cli.Float64Flag{
		Name:     "tierFee.optimistic",
		Usage:    "Initial tier fee (in GWei) paid to prover to generate an optimistic proofs",
		Category: proposerCategory,
		EnvVars:  []string{"TIER_FEE_OPTIMISTIC"},
	}
	SgxTierFee = &cli.Float64Flag{
		Name:     "tierFee.sgx",
		Usage:    "Initial tier fee (in GWei) paid to prover to generate a SGX proofs",
		Category: proposerCategory,
		EnvVars:  []string{"TIER_FEE_SGX"},
	}
	TierFeePriceBump = &cli.Uint64Flag{
		Name:     "tierFee.priceBump",
		Usage:    "Price bump percentage when no prover wants to accept the block at initial fee",
		Value:    10,
		Category: proposerCategory,
		EnvVars:  []string{"TIER_FEE_PRICE_BUMP"},
	}
	MaxTierFeePriceBumps = &cli.Uint64Flag{
		Name:     "tierFee.maxPriceBumps",
		Usage:    "If nobody accepts block at initial tier fee, how many iterations to increase tier fee before giving up",
		Category: proposerCategory,
		Value:    3,
		EnvVars:  []string{"TIER_FEE_MAX_PRICE_BUMPS"},
	}
	MinGasUsed = &cli.Uint64Flag{
		Name:     "epoch.minGasUsed",
		Usage:    "Minimum gas used for a transactions list to propose",
		Category: proposerCategory,
		Value:    0,
		EnvVars:  []string{"EPOCH_MIN_GAS_USED"},
	}
	MinTxListBytes = &cli.Uint64Flag{
		Name:     "epoch.minTxListBytes",
		Usage:    "Minimum bytes for a transactions list to propose",
		Category: proposerCategory,
		Value:    0,
		EnvVars:  []string{"EPOCH_MIN_TX_LIST_BYTES"},
	}
	MaxProposerDutiesSlots = &cli.Uint64Flag{
		Name:     "epoch.maxProposerDutiesSlots",
		Usage:    "The count of the proposer duties slots the proposer will track if it can propose a block",
		Category: proposerCategory,
		Value:    32,
		EnvVars:  []string{"EPOCH_MAX_PROPOSER_DUTIES_SLOTS"},
	}
	ProposerDutiesUpdateFreq = &cli.Uint64Flag{
		Name:     "epoch.proposerDutiesUpdateFreq",
		Usage:    "The frequency of updating proposer duties slots",
		Category: proposerCategory,
		Value:    24,
		EnvVars:  []string{"EPOCH_PROPOSER_DUTIES_UPDATE_FREQ"},
	}
	// Proposing metadata related.
	ExtraData = &cli.StringFlag{
		Name:     "extraData",
		Usage:    "Block extra data set by the proposer (default = client version)",
		Value:    version.CommitVersion(),
		Category: proposerCategory,
		EnvVars:  []string{"EXTRA_DATA"},
	}
	// Transactions pool related.
	TxPoolLocals = &cli.StringSliceFlag{
		Name:     "txPool.locals",
		Usage:    "Comma separated accounts to treat as locals (priority inclusion)",
		Category: proposerCategory,
		EnvVars:  []string{"TX_POOL_LOCALS"},
	}
	TxPoolLocalsOnly = &cli.BoolFlag{
		Name:     "txPool.localsOnly",
		Usage:    "If set to true, proposer will only propose transactions of local accounts",
		Value:    false,
		Category: proposerCategory,
		EnvVars:  []string{"TX_POOL_LOCALS_ONLY"},
	}
	MaxProposedTxListsPerEpoch = &cli.Uint64Flag{
		Name:     "txPool.maxTxListsPerEpoch",
		Usage:    "Maximum number of transaction lists which will be proposed inside one proposing epoch",
		Value:    1,
		Category: proposerCategory,
		EnvVars:  []string{"TX_POOL_MAX_TX_LISTS_PER_EPOCH"},
	}
	ProposeBlockIncludeParentMetaHash = &cli.BoolFlag{
		Name:     "includeParentMetaHash",
		Usage:    "Include parent meta hash when proposing block",
		Value:    false,
		Category: proposerCategory,
		EnvVars:  []string{"INCLUDE_PARENT_META_HASH"},
	}
	// Transaction related.
	BlobAllowed = &cli.BoolFlag{
		Name:    "l1.blobAllowed",
		Usage:   "Send EIP-4844 blob transactions when proposing blocks",
		Value:   false,
		EnvVars: []string{"L1_BLOB_ALLOWED"},
	}
	L1BlockBuilderTip = &cli.Uint64Flag{
		Name:     "l1.blockBuilderTip",
		Usage:    "Amount you wish to tip the L1 block builder",
		Value:    0,
		Category: proposerCategory,
		EnvVars:  []string{"L1_BLOCK_BUILDER_TIP"},
	}
)

// ProposerFlags All proposer flags.
var ProposerFlags = MergeFlags(CommonFlags, []cli.Flag{
	L1BeaconEndpoint,
	L2HTTPEndpoint,
	L2AuthEndpoint,
	JWTSecret,
	TaikoTokenAddress,
	L1ProposerPrivKey,
	L1MevBoostEndpoint,
	L2SuggestedFeeRecipient,
	PreconfDelay,
	TxPoolLocals,
	TxPoolLocalsOnly,
	ExtraData,
	MinGasUsed,
	MinTxListBytes,
	MaxProposedTxListsPerEpoch,
	MaxProposerDutiesSlots,
	ProposerDutiesUpdateFreq,
	ProverEndpoints,
	OptimisticTierFee,
	SgxTierFee,
	TierFeePriceBump,
	MaxTierFeePriceBumps,
	ProposeBlockIncludeParentMetaHash,
	BlobAllowed,
	L1BlockBuilderTip,
	SequencerRegistryAddress,
}, TxmgrFlags)
