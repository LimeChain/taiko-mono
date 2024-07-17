// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./ISequencerRegistry.sol";

contract TaikoStake {
    uint256 public activationThreshold = 1 ether; // Placeholder value
    mapping(address => uint256) public stakes;
    mapping(address => bool) public activated;

    event Staked(address indexed staker, uint256 amount);
    event Activated(address indexed staker);

    function _stakeSequencer(
        ISequencerRegistry sequencerRegistry,
        bytes calldata pubkey,
        ISequencerRegistry.ValidatorProof calldata validatorProof
    )
        internal
    {
        address staker = msg.sender;
        require(msg.value > 0, "Stake amount must be greater than zero");

        stakes[staker] += msg.value;

        if (stakes[staker] >= activationThreshold && !activated[staker]) {
            _activateSequencer(sequencerRegistry, staker, pubkey, validatorProof);
        }

        emit Staked(staker, msg.value);
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

        activated[staker] = true;
        emit Activated(staker);
    }
}
