// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../TaikoTest.sol";

contract SequencerRegistryTest is Test {
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
    }

    function _genValidatorProof()
        internal
        pure
        returns (ISequencerRegistry.ValidatorProof memory)
    {
        return ISequencerRegistry.ValidatorProof(0, 0, 0, 0, false, 0, bytes(""));
    }

    function _genAuthHash(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    function testInit() public {
        assertEq(registry.owner(), owner);
    }

    function _registerSequencer() internal returns (bytes memory) {
        bytes32 authHash = _genAuthHash(sequencer);
        bytes memory signature = "dummy_signature";
        bytes memory metadata = "metadata";
        ISequencerRegistry.ValidatorProof memory validatorProof = _genValidatorProof();
        registry.register(sequencer, metadata, authHash, signature, validatorProof);
        bytes memory pubkey = new bytes(48);
        assembly {
            mstore(add(pubkey, 32), authHash)
        }
        return pubkey;
    }

    function testRegisterSequencer() public {
        bytes32 authHash = _genAuthHash(sequencer);
        bytes memory signature = "dummy_signature";
        bytes memory metadata = "metadata";
        ISequencerRegistry.ValidatorProof memory validatorProof = _genValidatorProof();

        vm.expectEmit(address(registry));
        emit SequencerRegistry.SequencerUpdated(sequencer, false);

        registry.register(sequencer, metadata, authHash, signature, validatorProof);
        assertTrue(registry.isEligibleSigner(sequencer) == false);
        assertTrue(registry.registered(sequencer));
    }

    function testActivateSequencer() public {
        bytes memory pubkey = _registerSequencer();
        ISequencerRegistry.ValidatorProof memory validatorProof = _genValidatorProof();
        vm.prank(owner);
        registry.activate(pubkey, validatorProof);
        assertTrue(registry.isEligible(pubkey));
    }

    function testDeactivateSequencer() public {
        bytes memory signature = "dummy_signature";
        bytes memory pubkey = _registerSequencer();
        ISequencerRegistry.ValidatorProof memory validatorProof = _genValidatorProof();
        bytes32 authHash = _genAuthHash(sequencer);
        vm.prank(owner);
        registry.activate(pubkey, validatorProof);
        registry.deactivate(authHash, signature);
        assertFalse(registry.isEligible(pubkey));
    }

    function testIsIneligible() public {
        bytes memory pubkey = _registerSequencer();
        assertFalse(registry.isEligible(pubkey));
    }

    function testChangeRegistration() public {
        bytes memory pubkey = _registerSequencer();
        bytes memory newMetadata = "new_metadata";
        bytes32 authHash = _genAuthHash(sequencer);
        bytes memory signature = "new_dummy_signature";

        registry.changeRegistration(sequencer, newMetadata, authHash, signature);

        SequencerRegistry.Sequencer memory seq = registry.statusOf(pubkey);
        assertEq(string(seq.metadata), string(newMetadata), "Metadata should be updated.");

        vm.expectRevert("unathorized");
        registry.changeRegistration(address(0x3), newMetadata, authHash, signature);
    }

    function testMultipleSequencerRegistrations() public {
        for (uint256 i = 0; i < 5; i++) {
            address newSequencer = address(uint160(uint256(uint160(0x10)) + i));
            bytes32 authHash = _genAuthHash(newSequencer);
            bytes memory signature = bytes(abi.encodePacked("signature", i));
            bytes memory metadata = bytes(abi.encodePacked("metadata", i));
            ISequencerRegistry.ValidatorProof memory validatorProof = _genValidatorProof();

            registry.register(newSequencer, metadata, authHash, signature, validatorProof);

            assertTrue(
                registry.isEligibleSigner(newSequencer) == false,
                "Should not be eligible initially."
            );
        }

        bytes memory pubkey = _registerSequencer();
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        assertTrue(registry.eligibleCountAt(block.number) == 1, "Only one should be active.");
    }

    function testReactivation() public {
        bytes memory pubkey = _registerSequencer();
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        vm.expectRevert("already activated");
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());
    }

    function testRedeactivation() public {
        bytes memory pubkey = _registerSequencer();
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());
        registry.deactivate(_genAuthHash(sequencer), "dummy_signature");

        vm.expectRevert("already deactivated");
        registry.deactivate(_genAuthHash(sequencer), "dummy_signature");
    }

    function testForceDeactivate() public {
        bytes memory pubkey = _registerSequencer();
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        vm.expectEmit(true, true, false, true);
        emit SequencerRegistry.SequencerUpdated(sequencer, false);

        registry.forceDeactivate(pubkey, _genValidatorProof());

        assertFalse(
            registry.isEligible(pubkey),
            "Sequencer should not be eligible after forceful deactivation."
        );
    }

    function testEligibilityOverTime() public {
        bytes memory pubkey = _registerSequencer();
        vm.warp(100);
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        assertTrue(
            registry.isEligible(pubkey), "Sequencer should be eligible right after activation."
        );

        vm.warp(200);
        assertTrue(registry.isEligible(pubkey), "Sequencer should remain eligible.");
    }

    function testStatusChecks() public {
        bytes memory pubkey = _registerSequencer();
        vm.prank(owner);
        registry.activate(pubkey, _genValidatorProof());

        SequencerRegistry.Sequencer memory seqStatus = registry.statusOf(pubkey);
        assertEq(seqStatus.signer, sequencer, "Status should reflect the correct signer.");
        assertTrue(seqStatus.activationBlock > 0, "Activation block should be set.");
    }

    function testErrorHandlingInvalidRegistration() public {
        bytes memory signature = "dummy_signature";
        bytes32 authHash = _genAuthHash(sequencer);
        vm.expectRevert("invalid address");
        registry.register(address(0), "metadata", authHash, signature, _genValidatorProof());
    }
}
