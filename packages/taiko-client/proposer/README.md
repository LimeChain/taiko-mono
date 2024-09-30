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

## Steps to Register and Activate the Sequencer

1. **Navigate to the `sequencer_registry` directory:**

```sh
cd taiko-mono/packages/taiko-client/proposer/sequencer_registry
```

2. **Register the Sequencer**

Run the following command to register the sequencer:

```sh
go run . register
```

3. **Activate the Sequencer**

Run the following command to activate the sequencer:

```sh
go run . activate
```

Additional Information

- Ensure that the smart contracts are deployed and accessible at the addresses specified in the environment variables.
- The private key used should have sufficient funds to pay for the gas fees during the transactions.
