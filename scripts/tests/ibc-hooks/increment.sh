#!/bin/bash

echo ""
echo "#################"
echo "# IBC Hook call #"
echo "#################"
echo ""

BINARY=feather-cored
CHAIN_DIR=$(pwd)/.test-data
WALLET_1=$($BINARY keys show wallet1 -a --keyring-backend test --home $CHAIN_DIR/test-1)
WALLET_2=$($BINARY keys show wallet2 -a --keyring-backend test --home $CHAIN_DIR/test-2)

# Deploy the smart contract on chain to test the callbacks. (find the source code under the following url: `~/scripts/tests/ibc-hooks/counter/src/contract.rs`)
echo "Deploying counter contract"
CONTRACT_DEPLOYMENT_HASH_RES=$($BINARY tx wasm store $(pwd)/scripts/tests/ibc-hooks/counter/artifacts/counter.wasm --from $WALLET_2 --broadcast-mode sync --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test -y --gas 10000000 -o json | jq -r '.txhash')
sleep 2
echo "Querying contract code id"
CODE_ID=$($BINARY query tx $CONTRACT_DEPLOYMENT_HASH_RES --type hash --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json | jq -r '.events[4].attributes[1].value')

# Use Instantiate2 to instantiate the previous smart contract with a random hash to enable multiple instances of the same contract (when needed).
echo "Instantiating counter contract with code id $CODE_ID"
RANDOM_HASH=$(hexdump -vn16 -e'4/4 "%08X" 1 "\n"' /dev/urandom)
CONTRACT_INSTANTIATE_HASH_RES=$($BINARY tx wasm instantiate2 $CODE_ID '{"count": 0}' $RANDOM_HASH --no-admin --label="Label with $RANDOM_HASH" --from $WALLET_2 --broadcast-mode sync --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test -y --gas 10000000 -o json | jq -r '.txhash')
sleep 2
echo "Querying contract address"
CONTRACT_ADDRESS=$($BINARY query tx $CONTRACT_INSTANTIATE_HASH_RES --type hash --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json | jq -r '.events[4].attributes[0].value')
sleep 2
echo "Executing the IBC Hook to increment the counter from contract $CONTRACT_ADDRESS"
# First execute an IBC transfer to create the entry in the smart contract with the sender address ...
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $CONTRACT_ADDRESS 1stake --memo='{"wasm":{"contract": "'"$CONTRACT_ADDRESS"'" ,"msg": {"increment": {}}}}' --chain-id test-1 --home $CHAIN_DIR/test-1 --node tcp://localhost:16657 --keyring-backend test --from $WALLET_1 -y -o json)
sleep 2
# ... then send another transfer to increments the count value from 0 to 1, send 1 more stake to the contract address to validate that it increased the value correctly.
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $CONTRACT_ADDRESS 1stake --memo='{"wasm":{"contract": "'"$CONTRACT_ADDRESS"'" ,"msg": {"increment": {}}}}' --chain-id test-1 --home $CHAIN_DIR/test-1 --node tcp://localhost:16657 --keyring-backend test --from $WALLET_1 -y -o json)
export WALLET_1_WASM_SENDER=$($BINARY q ibchooks wasm-sender channel-0 "$WALLET_1" --chain-id test-1 --home $CHAIN_DIR/test-1 --node tcp://localhost:16657)

COUNT_RES=""
COUNT_FUNDS_RES=""
while [ "$COUNT_RES" != "1" ] || [ "$COUNT_FUNDS_RES" != "2" ]; do
    sleep 2
    # Query to assert that the counter value is 1 and the fund send are 2stake (remeber that the first time fund are send to the contract the counter is set to 0 instead of 1)
    COUNT_RES=$($BINARY query wasm contract-state smart "$CONTRACT_ADDRESS" '{"get_count": {"addr": "'"$WALLET_1_WASM_SENDER"'"}}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json |  jq -r '.data.count')
    COUNT_FUNDS_RES=$($BINARY query wasm contract-state smart "$CONTRACT_ADDRESS" '{"get_total_funds": {"addr": "'"$WALLET_1_WASM_SENDER"'"}}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json |  jq -r '.data.total_funds[0].amount')
    echo "transaction relayed count: $COUNT_RES and relayed funds: $COUNT_FUNDS_RES"
done

echo "Executing the IBC Hook to increment the counter on callback"
# Execute an IBC transfer with ibc_callback to test the callback acknowledgement twice.
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $WALLET_1_WASM_SENDER 1stake --memo='{"ibc_callback":"'"$CONTRACT_ADDRESS"'"}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test --from $WALLET_2 --broadcast-mode sync -y -o json)
sleep 2
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $WALLET_1_WASM_SENDER 1stake --memo='{"ibc_callback":"'"$CONTRACT_ADDRESS"'"}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test --from $WALLET_2 --broadcast-mode sync -y -o json)
export WALLET_2_WASM_SENDER=$($BINARY q ibchooks wasm-sender channel-0 "$WALLET_2" --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657)

COUNT_RES=""
while [ "$COUNT_RES" != "2" ]; do
    sleep 2
    # Query the smart contract to validate that it received the callback twice (notice that the queried addess is the contract address itself).
    COUNT_RES=$($BINARY query wasm contract-state smart "$CONTRACT_ADDRESS" '{"get_count": {"addr": "'"$CONTRACT_ADDRESS"'"}}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json |  jq -r '.data.count')
    echo "relayed callback transaction count: $COUNT_RES"
done

echo "Executing the IBC Hook to increment the counter on callback with timeout"
# Prepare two callback queries but this time with a timeout height that is unreachable (0-1) to test the timeout callback.
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $WALLET_1_WASM_SENDER 1stake --packet-timeout-height="0-1" --memo='{"ibc_callback":"'"$CONTRACT_ADDRESS"'"}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test --from $WALLET_2 --broadcast-mode sync -y -o json)
sleep 2
IBC_HOOK_RES=$($BINARY tx ibc-transfer transfer transfer channel-0 $WALLET_1_WASM_SENDER 1stake --packet-timeout-height="0-1" --memo='{"ibc_callback":"'"$CONTRACT_ADDRESS"'"}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 --keyring-backend test --from $WALLET_2 --broadcast-mode sync -y -o json)
export WALLET_2_WASM_SENDER=$($BINARY q ibchooks wasm-sender channel-0 "$WALLET_2" --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657)

COUNT_RES=""
while [ "$COUNT_RES" != "22" ]; do
    sleep 2
    # Query the smart contract to validate that it received the timeout callback twice and keep in mind that per each timeout the contract increases 10 counts (notice that the queried addess is the contract address itself).
    COUNT_RES=$($BINARY query wasm contract-state smart "$CONTRACT_ADDRESS" '{"get_count": {"addr": "'"$CONTRACT_ADDRESS"'"}}' --chain-id test-2 --home $CHAIN_DIR/test-2 --node tcp://localhost:26657 -o json |  jq -r '.data.count')
    echo "relayed timeout callback transaction count: $COUNT_RES"
done

echo ""
echo "##########################"
echo "# SUCCESS: IBC Hook call #"
echo "##########################"
echo ""
