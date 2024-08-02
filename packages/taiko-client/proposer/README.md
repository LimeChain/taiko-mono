# Taiko Sequencer Registration and Activation

This guide explains how to register and activate a sequencer for proposing blocks in the TaikoL1 smart contract.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- Environment variables set up (as listed below)

## Environment Variables

Set the following environment variables in your environment or in a `.env` file in the `taiko-mono/packages/taiko-client/proposer/sequencer_registry` directory:

CHAIN_ID=l1_chain_id
PRIVATE_KEY=your_private_key
RPC_URL=your_l1_rpc_url
TAIKOL1=your_taikol1_contract_address
SEQUENCER_REGISTRY=your_sequencer_registry_address
PROPOSER_ADDRESS=proposer_address

### Example `.env` File

CHAIN_ID=1
PRIVATE_KEY=0xabc123...your_private_key
RPC_URL=http://127.0.0.1:<l1_rpc_port>
TAIKOL1=0xYourTaikoL1ContractAddress
SEQUENCER_REGISTRY=0xYourSequencerRegistryAddress
PROPOSER_ADDRESS=0xYourProposerAddress

```sh
CHAIN_ID=3151908
PRIVATE_KEY=0xbcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31
RPC_URL=http://127.0.0.1:56029
TAIKOL1=0x38394c86870065a9df8b81acf9e4001f3fd1aa04
SEQUENCER_REGISTRY=0x4bb4c70df5601d726f7a738c2a4588b444356f33
PROPOSER_ADDRESS=0x8943545177806ED17B9F23F0a21ee5948eCaa776
```

## Steps to Register and Activate the Sequencer

1. **Navigate to the `sequencer_registry` directory:**

```sh
cd taiko-mono/packages/taiko-client/proposer/sequencer_registry
```

2. **Register the Sequencer**

Run the following command to register the sequencer:

```sh
go run ./sequencer_registry register
```

Additional Information

- Ensure that the smart contracts are deployed and accessible at the addresses specified in the environment variables.
- The private key used should have sufficient funds to pay for the gas fees during the transactions.
