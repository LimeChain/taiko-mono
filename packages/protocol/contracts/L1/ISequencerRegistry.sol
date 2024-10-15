// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

interface ISequencerRegistry {
    struct ValidatorProof {
        uint64 currentEpoch;
        uint64 activationEpoch;
        uint64 exitEpoch;
        uint256 validatorIndex;
        bool slashed;
        uint256 proofSlot;
        bytes sszProof;
    }

    struct Sequencer {
        bytes pubkey;
        bytes metadata;
        address signer;
        uint256 activationBlock;
        uint256 deactivationBlock;
    }

    event SequencerRegistered(address indexed signer, bytes pubkey);
    event SequencerChanged(address indexed oldSigner, address indexed newSigner, bytes pubkey);
    event SequencerActivated(address indexed signer);
    event SequencerDeactivated(address indexed signer);

    /**
     *     Registers the sequencer without activating them.
     *
     *     Authorised operation. Requires a signature over a digest by the validator via their BLS
     * pubkey.
     *     Requires EIP-2537 in order to verify the signature.
     *
     *     Rollup contracts will use the signer address when enforcing the primary or secondary
     * selection.
     *     The signature must authenticate the validator and must be over an authorisation hash.
     *     The authorisation hash must be verifiable by the contract and must include a nonce in
     * order to guard against replay attacks.
     *     The nonce derivation is a decision of the implementer but can be as simple as an
     * incremental counter.
     *     The implementation MUST check if the authHash is a
     * keccak256(protocol_version,contract_address,chain_id,nonce,function_selector,signer,metadata).
     *
     *     @param signer - the secp256k1 wallet that will be representing the sequencer via its
     * signatures
     *     @param metadata - metadata of the sequencer - including but not limited to version and
     * endpoint URL
     *     @param authHash - the authorisation hash -
     * keccak256(protocol_version,contract_address,chain_id,nonce,function_selector,signer,metadata).
     * The authorisation signature was
     * created by signing over these bytes.
     *     @param signature - the signature over the authHash performed by the validator key
     *     @param validatorProof - all the data needed to validate the existence of the validator in
     * the state tree of the beacon chain
     */
    function register(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature,
        ValidatorProof calldata validatorProof
    )
        external;

    /**
     *     Changes the sequencer signer and/or metadata.
     *
     *     Authorised operation. Similar requirements apply as in `register`
     *
     *     @param signer - the new wallet that will be representing the sequencer via its signatures
     *     @param metadata - the new metadata of the sequencer - including but not limited to
     * version and endpoint URL
     *     @param authHash - the authorisation hash -
     * keccak256(protocol_version,contract_address,chain_id,nonce,function_selector,signer,metadata).
     * The authorisation signature was
     * created by signing over these bytes.
     *     @param signature - the signature over the authHash performed by the validator key
     */
    function changeRegistration(
        address signer,
        bytes calldata metadata,
        bytes32 authHash,
        bytes calldata signature
    )
        external;

    /**
     *     Activates the sequencer finalising the registration process
     *     Implementers must make sure that the sequencer meets the activation (i.e. stake)
     * requirements before changing their status.
     *
     *     @param pubkey - the validator a BLS12-381 public key - 48 bytes
     *     @param validatorProof - all the data needed to validate the existence of the validator in
     * the state tree of the beacon chain
     */
    function activate(bytes calldata pubkey, ValidatorProof calldata validatorProof) external;

    /**
     *     Deactivates the sequencer.
     *
     *     Authorised operation. Similar requirements apply as in `register`.
     *     Implementers of the staking process must make sure that the sequencer is no longer active
     * before withdrawal disbursal
     *
     *     @param authHash - the authorisation hash -
     * keccak256(protocol_version,contract_address,chain_id,nonce,function_selector). The
     * authorisation signature was
     * created by signing over these bytes.
     *     @param signature - the signature over the authHash performed by the validator key
     */
    function deactivate(bytes32 authHash, bytes calldata signature) external;

    /**
     *     Forcefully deactivates a sequencer.
     *
     *     The caller must provide proof that the validator is no longer active or has been slashed.
     *
     *     @param pubkey - the validator a BLS12-381 public key - 48 bytes
     *     @param validatorProof - all the data needed to validate the existence and state of the
     * validator in the state tree of the beacon chain
     */
    function forceDeactivate(
        bytes calldata pubkey,
        ValidatorProof calldata validatorProof
    )
        external;

    /**
     *     Used to get the eligibility status of the sequencer identified by this pubkey
     *     @param pubkey - the validator a BLS12-381 public key - 48 bytes
     */
    function isEligible(bytes calldata pubkey) external view returns (bool);

    /**
     *     Returns the saved data for the sequencer identified by this pubkey
     *     @param pubkey - the validator a BLS12-381 public key - 48 bytes
     */
    function statusOf(bytes calldata pubkey) external view returns (Sequencer memory metadata);

    /**
     *     Used to get the activation status of the sequencer with this signer address
     *     @param signer - the associated signer address of a sequencer
     */
    function isEligibleSigner(address signer) external view returns (bool);

    /**
     *     Used to get the activation status of the sequencer with this signer address
     *     in a given block number.
     *     @param signer - the associated signer address of a sequencer
     *     @param blockNumber - the target block number.
     */
    function isEligibleSignerIn(address signer, uint256 blockNumber) external view returns (bool);

    /**
     *     Returns the data for a sequencer by its index
     */
    function sequencerByIndex(uint256 index)
        external
        view
        returns (address signer, bytes memory metadata, bytes memory pubkey);

    /**
     *     Number of Blocks after activation that the sequencer becomes eligible for sequencing
     */
    function activationTimeout() external view returns (uint8);

    /**
     *     Number of Blocks after deactivation that the sequencer becomes ineligible for sequencing
     */
    function deactivationPeriod() external view returns (uint8);

    /**
     *     Returns the total count of sequencers at this block number
     */
    function eligibleCountAt(uint256 blockNum) external view returns (uint256);

    /**
     *     Returns the protocol version used for authorising the digests.
     */
    function protocolVersion() external view returns (uint8);

    /**
     *     Returns if the signer is registered as a sequencer.
     */
    function isRegistered(address signer) external view returns (bool);

    /**
     *     Returns the deterministically selected fallback sequencer address.
     */
    function fallbackSigner(uint256 blockNum) external view returns (address);
}
