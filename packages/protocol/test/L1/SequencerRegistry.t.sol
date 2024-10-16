// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../TaikoTest.sol";

contract SequencerRegistryTest is Test {
    bytes signature = "dummy_signature";
    bytes metadata = "metadata";

    SequencerRegistry registry;
    address owner;
    address sequencer;

    function setUp() public {
        owner = address(0x1);
        sequencer = address(0x2);
        registry = new SequencerRegistry();
        vm.prank(address(0x0));
        registry.transferOwnership(owner);
        vm.prank(owner);
        registry.acceptOwnership();

        assertEq(1, registry.protocolVersion());
        assertEq(1, registry.activationTimeout());
        assertEq(2, registry.deactivationPeriod());
    }

    function _getPubkey(bytes32 data) internal pure returns (bytes memory) {
        bytes memory pubkey = new bytes(48);
        assembly {
            mstore(add(pubkey, 32), data)
        }
        return pubkey;
    }

    function _genValidatorProof()
        internal
        pure
        returns (ISequencerRegistry.ValidatorProof memory)
    {
        return ISequencerRegistry.ValidatorProof(0, 0, 0, 0, false, 0, bytes(""));
    }

    function _genAuthHash(
        uint256 _nonce,
        bytes4 _functionSelector,
        address _signer,
        bytes memory _metadata
    )
        internal
        view
        returns (bytes32)
    {
        return keccak256(
            abi.encodePacked(
                registry.protocolVersion(),
                address(registry),
                block.chainid,
                _nonce,
                _functionSelector,
                _signer,
                _metadata
            )
        );
    }

    function _genDeactivateHash(uint256 nonce) internal view returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                registry.protocolVersion(),
                address(registry),
                block.chainid,
                nonce,
                registry.deactivate.selector
            )
        );
    }

    function testInit() public {
        assertEq(registry.owner(), owner);
    }

    function test_sr_register() external {
        bytes32 authHash = _genAuthHash(0, registry.register.selector, sequencer, metadata);
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));
        bytes32 pubkeyHash = keccak256(pubkey);

        vm.expectEmit(true, true, true, true);
        emit ISequencerRegistry.SequencerRegistered(sequencer, pubkey);

        registry.register(sequencer, metadata, authHash, signature, _genValidatorProof());
        assertTrue(registry.isRegistered(sequencer));
        assertFalse(registry.isEligibleSigner(sequencer));
        assertFalse(registry.isEligibleSignerIn(sequencer, block.timestamp));

        ISequencerRegistry.Sequencer memory expect = ISequencerRegistry.Sequencer({
            pubkey: pubkey,
            metadata: metadata,
            signer: sequencer,
            activationBlock: 0,
            deactivationBlock: 0
        });

        assertStruct(expect, registry.statusOf(pubkey));
        assertEq(1, registry.nonces(pubkey));

        assertEq(pubkeyHash, registry.allValidators(0));
        assertEq(pubkeyHash, registry.sequencersToPubkeyHash(sequencer));

        assertFalse(registry.isEligible(pubkey));

        (address signer, bytes memory meta, bytes memory pkey) = registry.sequencerByIndex(0);
        assertEq(sequencer, signer);
        assertEq(metadata, meta);
        assertEq(pubkey, pkey);
    }

    function test_sr_register_invalid_address() external {
        bytes32 authHash = _genAuthHash(0, registry.register.selector, sequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_INVALID_ADDRESS.selector);
        registry.register(address(0), metadata, authHash, signature, _genValidatorProof());
    }

    function test_sr_register_signer_registered() external {
        _registerSequencer();

        bytes32 authHash = _genAuthHash(1, registry.register.selector, sequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_SIGNER_REGISTERED.selector);
        registry.register(sequencer, metadata, authHash, signature, _genValidatorProof());
    }

    function test_sr_register_invalid_auth_hash() external {
        vm.expectRevert(SequencerRegistry.SR_INVALID_AUTH_HASH.selector);
        registry.register(sequencer, metadata, bytes32(0), signature, _genValidatorProof());
    }

    function test_sr_register_validator_already_registered() external {
        _registerSequencer();

        address otherSequencer = address(0x5);
        bytes32 authHash = _genAuthHash(1, registry.register.selector, otherSequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_REGISTERED.selector);
        registry.register(otherSequencer, metadata, authHash, signature, _genValidatorProof());
    }

    function test_sr_changeRegistration() external {
        _registerSequencer();

        address newSequencer = address(0x5);
        bytes memory newMetadata = "new metadata";
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));
        bytes32 pubkeyHash = keccak256(pubkey);
        bytes32 authHash =
            _genAuthHash(1, registry.changeRegistration.selector, newSequencer, newMetadata);

        vm.expectEmit(true, true, true, true);
        emit ISequencerRegistry.SequencerChanged(newSequencer, sequencer, pubkey);

        registry.changeRegistration(newSequencer, newMetadata, authHash, signature);

        ISequencerRegistry.Sequencer memory expect = ISequencerRegistry.Sequencer({
            pubkey: pubkey,
            metadata: newMetadata,
            signer: newSequencer,
            activationBlock: 0,
            deactivationBlock: 0
        });

        assertStruct(expect, registry.statusOf(pubkey));
        assertEq(2, registry.nonces(pubkey));

        assertEq(pubkeyHash, registry.allValidators(0));
        assertEq(pubkeyHash, registry.sequencersToPubkeyHash(newSequencer));
        assertEq(bytes32(0), registry.sequencersToPubkeyHash(sequencer));

        assertTrue(registry.isRegistered(newSequencer));
        assertFalse(registry.isRegistered(sequencer));

        assertFalse(registry.isEligibleSigner(sequencer));
        assertFalse(registry.isEligibleSigner(newSequencer));

        assertFalse(registry.isEligible(pubkey));
    }

    function test_sr_changeRegistration_invalid_address() external {
        bytes32 authHash =
            _genAuthHash(0, registry.changeRegistration.selector, sequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_INVALID_ADDRESS.selector);
        registry.changeRegistration(address(0), metadata, authHash, signature);
    }

    function teset_sr_changeRegistration_signer_registered() external {
        _registerSequencer();

        bytes32 authHash =
            _genAuthHash(1, registry.changeRegistration.selector, sequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_SIGNER_REGISTERED.selector);
        registry.changeRegistration(sequencer, metadata, authHash, signature);
    }

    function test_sr_changeRegistration_invalid_auth_hash() external {
        vm.expectRevert(SequencerRegistry.SR_INVALID_AUTH_HASH.selector);
        registry.changeRegistration(sequencer, metadata, bytes32(0), signature);
    }

    function test_sr_changeRegistration_validator_not_registered() external {
        address otherSequencer = address(0x5);
        bytes32 authHash =
            _genAuthHash(0, registry.changeRegistration.selector, otherSequencer, metadata);

        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_NOT_REGISTERED.selector);
        registry.changeRegistration(otherSequencer, metadata, authHash, signature);
    }

    function test_sr_activate() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.expectEmit(true, true, true, true);
        emit ISequencerRegistry.SequencerActivated(sequencer);

        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);

        assertEq(block.number, seq.activationBlock);

        assertFalse(registry.isEligibleSigner(sequencer));
        assertFalse(registry.isEligible(pubkey));

        assertEq(0, registry.eligibleCountAt(block.number));
        assertEq(1, registry.eligibleCountAt(block.number + 1));
    }

    function test_sr_activate_eligible() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);

        assertEq(block.number, seq.activationBlock);

        assertFalse(registry.isEligibleSigner(sequencer));
        assertFalse(registry.isEligible(pubkey));

        vm.roll(2);

        assertTrue(registry.isEligibleSigner(sequencer));
        assertTrue(registry.isEligibleSignerIn(sequencer, block.number));
        assertTrue(registry.isEligible(pubkey));
    }

    function test_sr_activate_validator_not_registered() external {
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.prank(owner);
        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_NOT_REGISTERED.selector);
        registry.activate(pubkey, _genValidatorProof());
    }

    function test_sr_deactivate() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));
        bytes32 authHash = _genDeactivateHash(1);

        vm.expectEmit(true, true, true, true);
        emit ISequencerRegistry.SequencerDeactivated(sequencer);

        registry.deactivate(authHash, signature);

        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);
        assertEq(block.number, seq.deactivationBlock);

        assertEq(2, registry.nonces(pubkey));
    }

    function test_sr_deactivate_after_activate() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        vm.roll(block.number + 1);

        bytes32 authHash = _genDeactivateHash(1);
        registry.deactivate(authHash, signature);

        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);
        assertEq(block.number - 1, seq.activationBlock);
        assertEq(block.number, seq.deactivationBlock);

        vm.roll(block.number + 1);

        assertEq(2, registry.nonces(pubkey));
        assertTrue(registry.isEligibleSigner(sequencer));
    }

    function test_sr_deactivate_invalid_hash() external {
        _registerSequencer();
        bytes32 authHash = _genDeactivateHash(0);

        vm.expectRevert(SequencerRegistry.SR_INVALID_AUTH_HASH.selector);
        registry.deactivate(authHash, signature);
    }

    function test_sr_deactivate_validator_not_registered() external {
        bytes32 authHash = _genDeactivateHash(0);

        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_NOT_REGISTERED.selector);
        registry.deactivate(authHash, signature);
    }

    function test_sr_deactivate_validator_deactivated() external {
        _registerSequencer();
        bytes32 authHash = _genDeactivateHash(1);
        registry.deactivate(authHash, signature);

        authHash = _genDeactivateHash(2);
        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_DEACTIVATED.selector);
        registry.deactivate(authHash, signature);
    }

    function test_sr_forceDeactivate() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        registry.forceDeactivate(pubkey, _genValidatorProof());

        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);

        assertEq(block.timestamp, seq.deactivationBlock);
    }

    function test_sr_forceDeactivate_validator_not_registered() external {
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_NOT_REGISTERED.selector);
        registry.forceDeactivate(pubkey, _genValidatorProof());
    }

    function test_sr_forceDeactivate_validator_deactivated() external {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        registry.forceDeactivate(pubkey, _genValidatorProof());
        ISequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);
        assertEq(block.timestamp, seq.deactivationBlock);

        vm.expectRevert(SequencerRegistry.SR_VALIDATOR_DEACTIVATED.selector);
        registry.forceDeactivate(pubkey, _genValidatorProof());
    }

    function test_sr_sequencerByIndex_out_of_bounds() external {
        vm.expectRevert(SequencerRegistry.SR_INDEX_OUT_OF_BOUNDS.selector);
        registry.sequencerByIndex(0);
    }

    function test_sr_fallbackSigner_block_too_low() external {
        vm.expectRevert(SequencerRegistry.SR_BLOCK_TOO_LOW.selector);
        registry.fallbackSigner(block.number);
    }

    function test_sr_fallbackSigner_no_eligible_sequencers() external {
        vm.roll(11);
        vm.expectRevert(SequencerRegistry.SR_NO_ELIGIBLE_SEQUENCERS.selector);
        registry.fallbackSigner(block.number);
    }

    function test_eligible_over_time() public {
        _registerSequencer();
        bytes memory pubkey = _getPubkey(bytes32(uint256(1)));

        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        assertFalse(
            registry.isEligible(pubkey), "sequencer must be not be eligible right after activation."
        );

        vm.roll(2);
        assertTrue(registry.isEligible(pubkey), "sequencer must be eligible.");

        vm.roll(100);
        assertTrue(registry.isEligible(pubkey), "sequencer must remain eligible.");

        assertEq(sequencer, registry.fallbackSigner(block.number));
    }

    function _registerSequencer() internal {
        bytes32 authHash = _genAuthHash(0, registry.register.selector, sequencer, metadata);
        registry.register(sequencer, metadata, authHash, signature, _genValidatorProof());
    }

    function assertStruct(
        ISequencerRegistry.Sequencer memory a,
        ISequencerRegistry.Sequencer memory b
    )
        internal
    {
        bytes memory aPubkey = a.pubkey;
        bytes memory bPubkey = b.pubkey;

        assertEq(aPubkey, bPubkey);
        assertEq(a.metadata, b.metadata);
        assertEq(a.signer, b.signer);
        assertEq(a.activationBlock, b.activationBlock);
        assertEq(a.deactivationBlock, b.deactivationBlock);
    }
}
