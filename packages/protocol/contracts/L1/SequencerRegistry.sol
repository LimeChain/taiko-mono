// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/EssentialContract.sol";
import "./ISequencerRegistry.sol";

/// @title SequencerRegistry
contract SequencerRegistry is EssentialContract, ISequencerRegistry {
    /// @dev Emitted when the status of a sequencer is updated.
    /// @param sequencer The address of the sequencer whose state has updated.
    /// @param enabled If the sequencer is now enabled or not.
    event SequencerUpdated(address indexed sequencer, bool enabled);

    bytes[] public allSequencers;
    /// @notice Whitelisted sequencers
    mapping(address => bool) public whitelisted;
    /// @notice BLS public key to Sequencer mapping
    mapping(bytes => Sequencer) public seqByPubkey;
    uint256 public activationThreshold;
    uint256 public withdrawalChallengePeriod;

    uint8 public constant PROTOCOL_VERSION = 1;

    uint256[50] private __gap;

    /// @notice Initializes the contract with the provided address manager.
    /// @param _owner The address of the owner.
    function init(
        address _owner,
        uint256 _activationThreshold,
        uint256 _withdrawalChallengePeriod
    )
        external
        initializer
    {
        __Essential_init(_owner);
        activationThreshold = _activationThreshold;
        withdrawalChallengePeriod = _withdrawalChallengePeriod;
    }

    /// @notice Registers a sequencer with its metadata.
    /// @param signer The address of the sequencer
    /// @param metadata Metadata associated with the sequencer
    /// @param authHash The authorisation hash
    /// @param signature The signature over the authHash performed by the validator key
    /// @param // validatorProof The data needed to validate the existence of the validator
    function register(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature,
        ValidatorProof calldata /*validatorProof*/
    )
        external
        override
    {
        require(signer != address(0), "invalid address");
        // Mock signature verification and SSZ multiproof verification
        require(_verifySignature(authHash, signature), "invalid signature");

        bytes memory pubkey = _recoverPubKey(authHash, signature);
        Sequencer storage seq = seqByPubkey[pubkey];

        require(seq.signer == address(0), "already registered");

        seq.pubkey = pubkey;
        seq.metadata = metadata;
        seq.signer = signer;
        seq.activationBlock = 0; // Not activated yet
        seq.deactivationBlock = 0;

        allSequencers.push(pubkey);
        whitelisted[signer] = false;

        emit SequencerUpdated(signer, false);
    }

    /// @notice Changes the registration details of a sequencer.
    /// @param signer The new address of the sequencer
    /// @param metadata The new metadata associated with the sequencer
    /// @param authHash The authorisation hash
    /// @param signature The signature over the authHash performed by the validator key
    function changeRegistration(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature
    )
        external
        override
    {
        require(signer != address(0), "invalid address");
        require(_verifySignature(authHash, signature), "invalid signature");

        bytes memory pubkey = _recoverPubKey(authHash, signature);
        Sequencer storage seq = seqByPubkey[pubkey];

        require(seq.signer != address(0), "not registered");
        require(signer == seq.signer, "unathorized");

        seq.signer = signer;
        seq.metadata = metadata;

        emit SequencerUpdated(signer, true);
    }

    /// @notice Activates a sequencer.
    /// @param pubkey The validator's BLS12-381 public key
    /// @param // validatorProof The data needed to validate the existence of the validator
    function activate(
        bytes calldata pubkey,
        ValidatorProof calldata /*validatorProof*/
    )
        external
        override
        onlyOwner
    {
        Sequencer storage seq = seqByPubkey[pubkey];
        require(seq.signer != address(0), "not registered");
        require(seq.activationBlock == 0, "already activated");
        // Mock SSZ proof validation (Here we should validate the validatorProof)

        seq.activationBlock = block.number;
        whitelisted[seq.signer] = true;

        emit SequencerUpdated(seq.signer, true);
    }

    /// @notice Deactivates a sequencer.
    /// @param authHash The authorisation hash
    /// @param signature The signature over the authHash performed by the validator key
    function deactivate(bytes32 authHash, bytes calldata signature) external override {
        require(_verifySignature(authHash, signature), "invalid signature");

        bytes memory pubkey = _recoverPubKey(authHash, signature);

        Sequencer storage seq = seqByPubkey[pubkey];
        require(seq.signer != address(0), "not registered");
        require(seq.deactivationBlock == 0, "already deactivated");

        seq.deactivationBlock = block.number;
        whitelisted[seq.signer] = false;

        emit SequencerUpdated(seq.signer, false);
    }

    /// @notice Forcefully deactivates a sequencer.
    /// @param pubkey The validator's BLS12-381 public key
    /// @param // validatorProof The data needed to validate the existence and state of the
    /// validator
    function forceDeactivate(
        bytes calldata pubkey,
        ValidatorProof calldata /*validatorProof*/
    )
        external
        override
    {
        Sequencer storage seq = seqByPubkey[pubkey];
        require(seq.signer != address(0), "not registered");
        // Mock SSZ proof validation (Here we should validate the validatorProof)

        seq.deactivationBlock = block.number;
        whitelisted[seq.signer] = false;

        emit SequencerUpdated(seq.signer, false);
    }

    /// @notice Checks if a sequencer is eligible.
    /// @param pubkey The validator's BLS12-381 public key
    function isEligible(bytes calldata pubkey) external view override returns (bool) {
        Sequencer storage seq = seqByPubkey[pubkey];
        return seq.activationBlock != 0 && block.number >= seq.activationBlock;
    }

    /// @notice Returns the status of a sequencer.
    /// @param pubkey The validator's BLS12-381 public key
    function statusOf(bytes calldata pubkey) external view override returns (Sequencer memory) {
        return seqByPubkey[pubkey];
    }

    /// @inheritdoc ISequencerRegistry
    function isEligibleSigner(address proposer) public view override returns (bool) {
        return whitelisted[proposer];
    }

    /// @inheritdoc ISequencerRegistry
    function eligibleCountAt(uint256 blockNumber) external view override returns (uint256) {
        uint256 count = 0;
        for (uint256 i = 0; i < allSequencers.length; i++) {
            bytes memory pubkey = allSequencers[i];
            Sequencer storage seq = seqByPubkey[pubkey];
            if (isEligibleSigner(seq.signer) && seq.activationBlock <= blockNumber) {
                count++;
            }
        }
        return count;
    }

    /// @inheritdoc ISequencerRegistry
    function sequencerByIndex(uint256 _index)
        external
        view
        override
        returns (address, bytes memory, bytes memory)
    {
        if (_index >= allSequencers.length) {
            revert("index out of bounds");
        }

        bytes memory pubkey = allSequencers[_index];
        Sequencer storage seq = seqByPubkey[pubkey];
        bytes memory metadata = seq.metadata;

        return (seq.signer, metadata, pubkey);
    }

    // TODO: The implementations must recover the 48 bytes BLS12 pub key - pubkey - of the caller
    // through the signature recovery process over the signature and authHash.
    // Note that BLS signature recovery is not possible until EIP-2573 is included (currently
    // scheduled for Pectra).
    function _recoverPubKey(
        bytes32, /*authHash*/
        bytes calldata /*signature*/
    )
        private
        pure
        returns (bytes memory)
    {
        bytes memory b = new bytes(48);
        address signer = address(0);
        assembly {
            mstore(add(b, 32), shl(96, signer))
        }
        return b;
    }

    function _verifySignature(
        bytes32, /*authHash*/
        bytes calldata /*signature*/
    )
        private
        pure
        returns (bool)
    {
        // Mock signature verification
        return true;
    }

    function activationTimeout() external pure override returns (uint8) {
        return 1; // Placeholder
    }

    function deactivationPeriod() external pure override returns (uint8) {
        return 1; // Placeholder
    }

    function protocolVersion() external pure override returns (uint8) {
        return PROTOCOL_VERSION;
    }
}
