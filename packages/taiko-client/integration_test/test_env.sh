#!/bin/bash

source internal/docker/docker_env.sh
source scripts/common.sh

# get deployed contract address.
DEPLOYMENT_JSON=$(cat ../protocol/deployments/deploy_l1.json)
export TAIKO_L1_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.taiko' | sed 's/\"//g')
export TAIKO_L2_ADDRESS=0x1670010000000000000000000000000000010001
export SEQUENCER_REGISTRY=$(echo "$DEPLOYMENT_JSON" | jq '.sequencer_registry' | sed 's/\"//g')
export TAIKO_TOKEN_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.taiko_token' | sed 's/\"//g')
export TIMELOCK_CONTROLLER=$(echo "$DEPLOYMENT_JSON" | jq '.timelock_controller' | sed 's/\"//g')
export ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.rollup_address_manager' | sed 's/\"//g')
export GUARDIAN_PROVER_CONTRACT_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.guardian_prover' | sed 's/\"//g')
export GUARDIAN_PROVER_MINORITY_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.guardian_prover_minority' | sed 's/\"//g')
export PROVER_SET_ADDRESS=$(echo "$DEPLOYMENT_JSON" | jq '.prover_set' | sed 's/\"//g')
export L1_CONTRACT_OWNER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export L1_SECURITY_COUNCIL_PRIVATE_KEY=0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
export L1_PROPOSER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export L2_SUGGESTED_FEE_RECIPIENT=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
export L1_PROVER_PRIVATE_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
export TREASURY=0x1670010000000000000000000000000000010001
export VERBOSITY=3

# show the integration test environment variables.
# L1_BEACON_HTTP_ENDPOINT=$L1_BEACON_HTTP_ENDPOINT
echo "RUN_TESTS=true
L1_NODE_HTTP_ENDPOINT=$L1_NODE_HTTP_ENDPOINT
L1_NODE_WS_ENDPOINT=$L1_NODE_WS_ENDPOINT
L2_SUGGESTED_FEE_RECIPIENT=$L2_SUGGESTED_FEE_RECIPIENT
L2_EXECUTION_ENGINE_HTTP_ENDPOINT=$L2_EXECUTION_ENGINE_HTTP_ENDPOINT
L2_EXECUTION_ENGINE_WS_ENDPOINT=$L2_EXECUTION_ENGINE_WS_ENDPOINT
L2_EXECUTION_ENGINE_AUTH_ENDPOINT=$L2_EXECUTION_ENGINE_AUTH_ENDPOINT
TAIKO_L1_ADDRESS=$TAIKO_L1_ADDRESS
TAIKO_L2_ADDRESS=$TAIKO_L2_ADDRESS
SEQUENCER_REGISTRY=$SEQUENCER_REGISTRY
TAIKO_TOKEN_ADDRESS=$TAIKO_TOKEN_ADDRESS
PROVER_SET_ADDRESS=$PROVER_SET_ADDRESS
TIMELOCK_CONTROLLER=$TIMELOCK_CONTROLLER
ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS=$ROLLUP_ADDRESS_MANAGER_CONTRACT_ADDRESS
GUARDIAN_PROVER_CONTRACT_ADDRESS=$GUARDIAN_PROVER_CONTRACT_ADDRESS
GUARDIAN_PROVER_MINORITY_ADDRESS=$GUARDIAN_PROVER_MINORITY_ADDRESS
L1_CONTRACT_OWNER_PRIVATE_KEY=$L1_CONTRACT_OWNER_PRIVATE_KEY
L1_SECURITY_COUNCIL_PRIVATE_KEY=$L1_SECURITY_COUNCIL_PRIVATE_KEY
L1_PROPOSER_PRIVATE_KEY=$L1_PROPOSER_PRIVATE_KEY
L1_PROVER_PRIVATE_KEY=$L1_PROVER_PRIVATE_KEY
TREASURY=$TREASURY
JWT_SECRET=$JWT_SECRET
VERBOSITY=$VERBOSITY" > integration_test/.env
