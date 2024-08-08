// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "./ISequencerRegistry.sol";

contract TaikoProposer {
    /// @notice Validates if the given address is an active proposer
    /// @param proposer The address of the proposer
    /// @return True if the proposer is active, false otherwise
    function validateProposer(
        ISequencerRegistry sequencerRegistry,
        address proposer
    )
        public
        view
        returns (bool)
    {
        return sequencerRegistry.isEligibleSigner(proposer);
    }

    /// @notice Selects a fallback proposer based on the parent block hash
    /// @param parentBlockHash The hash of the parent block
    /// @return The address of the fallback proposer
    function fallbackProposer(
        ISequencerRegistry sequencerRegistry,
        uint256 parentBlockHash
    )
        public
        view
        returns (address)
    {
        uint256 fallbackIndex =
            uint256(parentBlockHash) % sequencerRegistry.eligibleCountAt(block.number);
        (address signer,,) = sequencerRegistry.sequencerByIndex(fallbackIndex);
        return signer;
    }
}
