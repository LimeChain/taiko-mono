// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./LibUtils.sol";

/// @title LibStaking
/// @notice A library for handling staking of L1 proposers that would like
/// to opt-in as sequencers.
library LibStaking {
    event Staked(bytes32 indexed pubkeyHash, bytes pubkey, uint256 amount);

    error L1_STAKING_ALREADY_STAKED();
    error L1_STAKING_INVALID_THRESHOLD();

    function stakeSequencer(
        TaikoData.State storage _state,
        TaikoData.Config memory _config,
        ISequencerRegistry _sequencerRegistry,
        bytes calldata _pubkey,
        ISequencerRegistry.ValidatorProof calldata _validatorProof
    )
        internal
    {
        if (msg.value < _config.activationThreshold) {
            revert L1_STAKING_INVALID_THRESHOLD();
        }

        bytes32 pubkeyHash = keccak256(_pubkey);
        if (_state.stakes[pubkeyHash] >= _config.activationThreshold) {
            revert L1_STAKING_ALREADY_STAKED();
        }

        _sequencerRegistry.activate(_pubkey, _validatorProof);
        _state.stakes[pubkeyHash] = msg.value;

        emit Staked(pubkeyHash, _pubkey, msg.value);
    }
}
