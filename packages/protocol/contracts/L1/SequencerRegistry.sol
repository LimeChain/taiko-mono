// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../common/EssentialContract.sol";
import "./ISequencerRegistry.sol";

/// @title SequencerRegistry
/// @notice This contract serves as registry of L1 proposers that opt-in
/// to be Taiko sequencers. It stores the information about the status of
/// the sequencers and their metadata.
contract SequencerRegistry is EssentialContract, ISequencerRegistry {
    /// @notice Protocol version
    uint8 public constant PROTOCOL_VERSION = 1;
    /// @notice Number of blocks after activation when the sequencer becomes eligible for sequencing
    uint8 public constant ACTIVATION_TIMEOUT = 1;
    /// @notice Number of blocks after deactivation when the sequencer is no longer eligible for
    /// sequencing
    uint8 public constant DEACTIVATION_PERIOD = 2;

    bytes32[] public allValidators;
    mapping(bytes pubkey => uint256) public nonces;
    mapping(address sequencer => bytes32 pubkeyHash) public sequencersToPubkeyHash;
    mapping(bytes32 pubkeyHash => Sequencer sequencer) public validators;

    uint256[46] private __gap;

    error SR_BLOCK_TOO_LOW();
    error SR_INDEX_OUT_OF_BOUNDS();
    error SR_INVALID_ADDRESS();
    error SR_INVALID_AUTH_HASH();
    error SR_INVALID_AUTH_SIGNATURE();
    error SR_INVALID_PROOF();
    error SR_NO_ELIGIBLE_SEQUENCERS();
    error SR_SIGNER_REGISTERED();
    error SR_VALIDATOR_NOT_REGISTERED();
    error SR_VALIDATOR_REGISTERED();
    error SR_VALIDATOR_DEACTIVATED();

    /// @notice Initializes the contract.
    /// @param _owner The address of the owner.
    function init(address _owner) external initializer {
        __Essential_init(_owner);
    }

    /// @inheritdoc ISequencerRegistry
    function register(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature,
        ValidatorProof calldata validatorProof
    )
        external
        override
    {
        if (signer == address(0)) {
            revert SR_INVALID_ADDRESS();
        }
        if (sequencersToPubkeyHash[signer] != bytes32(0)) {
            revert SR_SIGNER_REGISTERED();
        }

        bytes memory pubkey = _recoverPubkey(authHash, signature);
        bytes32 _authenticationHash = _authHash(
            PROTOCOL_VERSION,
            address(this),
            block.chainid,
            nonces[pubkey],
            this.register.selector,
            signer,
            metadata
        );

        if (authHash != _authenticationHash) {
            revert SR_INVALID_AUTH_HASH();
        }
        if (!_verifySignature(pubkey, authHash, signature)) {
            revert SR_INVALID_AUTH_SIGNATURE();
        }

        if (!_verifyValidatorProof(pubkey, validatorProof)) {
            revert SR_INVALID_PROOF();
        }

        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer storage sequencer = validators[pubkeyHash];
        if (sequencer.signer != address(0)) {
            revert SR_VALIDATOR_REGISTERED();
        }

        unchecked {
            sequencer.pubkey = pubkey;
            sequencer.metadata = metadata;
            sequencer.signer = signer;
            sequencer.activationBlock = 0;
            sequencer.deactivationBlock = 0;

            nonces[pubkey]++;
            sequencersToPubkeyHash[signer] = pubkeyHash;
            allValidators.push(pubkeyHash);
        }

        emit SequencerRegistered(signer, pubkey);
    }

    /// @inheritdoc ISequencerRegistry
    function changeRegistration(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature
    )
        external
        override
    {
        if (signer == address(0)) {
            revert SR_INVALID_ADDRESS();
        }
        if (sequencersToPubkeyHash[signer] != bytes32(0)) {
            revert SR_SIGNER_REGISTERED();
        }

        bytes memory pubkey = _recoverPubkey(authHash, signature);
        bytes32 _authenticationHash = _authHash(
            PROTOCOL_VERSION,
            address(this),
            block.chainid,
            nonces[pubkey],
            this.changeRegistration.selector,
            signer,
            metadata
        );

        if (authHash != _authenticationHash) {
            revert SR_INVALID_AUTH_HASH();
        }
        if (!_verifySignature(pubkey, authHash, signature)) {
            revert SR_INVALID_AUTH_SIGNATURE();
        }

        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer storage sequencer = validators[pubkeyHash];

        if (sequencer.signer == address(0)) {
            revert SR_VALIDATOR_NOT_REGISTERED();
        }

        address oldSigner = sequencer.signer;

        unchecked {
            sequencer.signer = signer;
            sequencer.metadata = metadata;

            nonces[pubkey]++;
            sequencersToPubkeyHash[signer] = pubkeyHash;
            delete sequencersToPubkeyHash[oldSigner];
        }

        emit SequencerChanged(signer, oldSigner, pubkey);
    }

    /// @inheritdoc ISequencerRegistry
    function activate(
        bytes calldata pubkey,
        ValidatorProof calldata validatorProof
    )
        external
        override
        onlyOwner
    {
        if (!_verifyValidatorProof(pubkey, validatorProof)) {
            revert SR_INVALID_PROOF();
        }

        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer storage seq = validators[pubkeyHash];

        if (seq.signer == address(0)) {
            revert SR_VALIDATOR_NOT_REGISTERED();
        }

        seq.activationBlock = block.number;
        seq.deactivationBlock = 0;

        emit SequencerActivated(seq.signer);
    }

    /// @inheritdoc ISequencerRegistry
    function deactivate(bytes32 authHash, bytes calldata signature) external override {
        bytes memory pubkey = _recoverPubkey(authHash, signature);
        bytes32 _authenticationHash = _authDeactivationHash(
            PROTOCOL_VERSION, address(this), block.chainid, nonces[pubkey], this.deactivate.selector
        );

        if (authHash != _authenticationHash) {
            revert SR_INVALID_AUTH_HASH();
        }
        if (!_verifySignature(pubkey, authHash, signature)) {
            revert SR_INVALID_AUTH_SIGNATURE();
        }

        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer storage seq = validators[pubkeyHash];

        if (seq.signer == address(0)) {
            revert SR_VALIDATOR_NOT_REGISTERED();
        }
        if (seq.deactivationBlock != 0) {
            revert SR_VALIDATOR_DEACTIVATED();
        }

        unchecked {
            nonces[pubkey]++;

            seq.deactivationBlock = block.number;
        }

        emit SequencerDeactivated(seq.signer);
    }

    /// @inheritdoc ISequencerRegistry
    function forceDeactivate(
        bytes calldata pubkey,
        ValidatorProof calldata validatorProof
    )
        external
        override
    {
        if (!_verifyValidatorProof(pubkey, validatorProof)) {
            revert SR_INVALID_PROOF();
        }

        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer storage seq = validators[pubkeyHash];
        if (seq.signer == address(0)) {
            revert SR_VALIDATOR_NOT_REGISTERED();
        }
        if (seq.deactivationBlock != 0) {
            revert SR_VALIDATOR_DEACTIVATED();
        }

        seq.deactivationBlock = block.number;

        emit SequencerDeactivated(seq.signer);
    }

    /// @inheritdoc ISequencerRegistry
    function isEligible(bytes calldata pubkey) external view override returns (bool) {
        bytes32 pubkeyHash = keccak256(pubkey);
        Sequencer memory sequencer = validators[pubkeyHash];

        return _isEligibleSequencer(sequencer, block.number);
    }

    /// @inheritdoc ISequencerRegistry
    function statusOf(bytes calldata pubkey)
        external
        view
        override
        returns (Sequencer memory data)
    {
        bytes32 pubkeyHash = keccak256(pubkey);
        return validators[pubkeyHash];
    }

    /// @inheritdoc ISequencerRegistry
    function isEligibleSigner(address signer) external view override returns (bool) {
        return isEligibleSignerIn(signer, block.number);
    }

    function isEligibleSignerIn(
        address signer,
        uint256 blockNumber
    )
        public
        view
        override
        returns (bool)
    {
        bytes32 pubkeyHash = sequencersToPubkeyHash[signer];
        if (pubkeyHash == bytes32(0)) {
            return false;
        }

        Sequencer memory sequencer = validators[pubkeyHash];

        return _isEligibleSequencer(sequencer, blockNumber);
    }

    /// @inheritdoc ISequencerRegistry
    function sequencerByIndex(uint256 index)
        external
        view
        override
        returns (address signer, bytes memory metadata, bytes memory pubkey)
    {
        if (index >= allValidators.length) {
            revert SR_INDEX_OUT_OF_BOUNDS();
        }

        bytes32 pubkeyHash = allValidators[index];
        Sequencer memory seq = validators[pubkeyHash];

        return (seq.signer, seq.metadata, seq.pubkey);
    }

    /// @inheritdoc ISequencerRegistry
    function eligibleCountAt(uint256 blockNumber) public view override returns (uint256) {
        uint256 count = 0;
        for (uint256 i = 0; i < allValidators.length; i++) {
            if (_isEligibleSequencer(validators[allValidators[i]], blockNumber)) {
                count++;
            }
        }

        return count;
    }

    /// @inheritdoc ISequencerRegistry
    function activationTimeout() external pure override returns (uint8) {
        return ACTIVATION_TIMEOUT;
    }

    /// @inheritdoc ISequencerRegistry
    function deactivationPeriod() external pure override returns (uint8) {
        return DEACTIVATION_PERIOD;
    }

    /// @inheritdoc ISequencerRegistry
    function protocolVersion() external pure override returns (uint8) {
        return PROTOCOL_VERSION;
    }

    /// @inheritdoc ISequencerRegistry
    function isRegistered(address signer) external view returns (bool) {
        return sequencersToPubkeyHash[signer] != bytes32(0);
    }

    /// @notice Deterministic fallback selection.
    /// @dev Takes the 10th parent block of the given block as a seed.
    /// Using the seed it takes the mod of the current eligible sequencers.
    /// Works only for the last 256 blocks.
    /// @param _blockNum Target block number.
    /// @return The address of the fallback proposer
    function fallbackSigner(uint256 _blockNum) external view override returns (address) {
        if (_blockNum < 10) {
            revert SR_BLOCK_TOO_LOW();
        }
        bytes32 parentBlockHash = blockhash(_blockNum - 10);

        address[] memory eligibleAt = _eligibleAt(_blockNum);
        if (eligibleAt.length == 0) {
            revert SR_NO_ELIGIBLE_SEQUENCERS();
        }

        uint256 fallbackIndex = uint256(parentBlockHash) % eligibleAt.length;

        return eligibleAt[fallbackIndex];
    }

    function _eligibleAt(uint256 _blockNum) internal view returns (address[] memory) {
        uint256 count = eligibleCountAt(_blockNum);

        address[] memory eligible = new address[](count);
        uint256 index = 0;

        for (uint256 i = 0; i < allValidators.length; i++) {
            Sequencer memory seq = validators[allValidators[i]];
            if (_isEligibleSequencer(seq, _blockNum)) {
                eligible[index] = seq.signer;
                index++;
            }
        }

        return eligible;
    }

    /// @dev Recovers the 48 bytes BLS12 pub key - pubkey - of the caller
    /// through the signature recovery process over the signature and authHash.
    function _recoverPubkey(bytes32, bytes calldata) private pure returns (bytes memory) {
        // TODO: Implement once BLS signature recovery is supported (after EIP-2573) is merged.
        bytes32 one = bytes32(uint256(1));
        bytes memory pubkey = new bytes(48);
        assembly {
            mstore(add(pubkey, 32), one)
        }

        return pubkey;
    }

    /// @dev Verifies the 48 bytes BLS12 pub key through signature verification
    function _verifySignature(bytes memory, bytes32, bytes calldata) private pure returns (bool) {
        // TODO: Implement once BLS signatures are supported (after EIP-2573)
        return true;
    }

    /// @notice Verifies that the validator exists on the beacon chain.
    function _verifyValidatorProof(
        bytes memory,
        ValidatorProof calldata
    )
        private
        pure
        returns (bool)
    {
        /// TODO: Out of proposal scope.
        return true;
    }

    function _isEligibleSequencer(
        Sequencer memory seq,
        uint256 blockNumber
    )
        internal
        pure
        returns (bool)
    {
        if (seq.deactivationBlock == 0) {
            return
                seq.activationBlock > 0 && blockNumber >= seq.activationBlock + ACTIVATION_TIMEOUT;
        }
        return seq.activationBlock > 0 && blockNumber >= seq.activationBlock + ACTIVATION_TIMEOUT
            && blockNumber < seq.deactivationBlock + DEACTIVATION_PERIOD;
    }

    function _authHash(
        uint8 _protocolVersion,
        address _contract,
        uint256 _chainid,
        uint256 _nonce,
        bytes4 _functionSelector,
        address _signer,
        bytes calldata _metadata
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _protocolVersion, _contract, _chainid, _nonce, _functionSelector, _signer, _metadata
            )
        );
    }

    function _authDeactivationHash(
        uint8 _protocolVersion,
        address _contract,
        uint256 _chainid,
        uint256 _nonce,
        bytes4 _functionSelector
    )
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(_protocolVersion, _contract, _chainid, _nonce, _functionSelector)
        );
    }
}
