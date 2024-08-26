// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./ISequencerRegistry.sol";
import "./TaikoData.sol";

contract TaikoStake {
    event Staked(address indexed staker, uint256 amount);
    event Activated(address indexed staker);

    function _stakeSequencer(
        TaikoData.State storage _state,
        TaikoData.Config memory _config,
        ISequencerRegistry sequencerRegistry,
        bytes calldata pubkey,
        ISequencerRegistry.ValidatorProof calldata validatorProof
    )
        internal
    {
        address staker = msg.sender;
        require(msg.value > 0, "Stake amount must not be zero");
        require(!sequencerRegistry.activated(staker), "Already activated");

        _state.stakes[staker] += msg.value;

        if (_state.stakes[staker] >= _config.activationThreshold) {
            _activateSequencer(sequencerRegistry, staker, pubkey, validatorProof);
        }

        emit Staked(staker, _state.stakes[staker]);
    }

    function _activateSequencer(
        ISequencerRegistry sequencerRegistry,
        address staker,
        bytes calldata pubkey,
        ISequencerRegistry.ValidatorProof calldata validatorProof
    )
        private
    {
        sequencerRegistry.activate(pubkey, validatorProof);

        emit Activated(staker);
    }
}
