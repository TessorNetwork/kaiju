#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(kaijud status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(kaijud status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
GAS_PRICES="0.025ukaiju"
CHAIN_ID="pandora-4"

kaijud_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  kaijud tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

kaijud_q() {
  kaijud q "$@" --output=json | jq .
}

NEW_DID="$RANDOM"
FULL_DID="did:earth:pandora-4:$NEW_DID"

echo "Creating DID..."
DID=$(yes $PASSWORD | kaijud tx iid create-iid "$NEW_DID" "pandora-4" --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y | jq .)
echo $DID

#echo "Adding 2 contexts.."
#CONTEXT1=$(yes $PASSWORD | kaijud tx iid add-iid-context "$NEW_DID" "kaiju" "https://w3id.org/kaiju/NS/" --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y --output json | jq .)
#echo $CONTEXT1

#CONTEXT2=$(yes $PASSWORD | kaijud tx iid add-iid-context "$NEW_DID" "iid" "https://w3id.org/kaiju/NS/"  --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y --output json | jq .)
#echo $CONTEXT2

echo "Adding metadata..."
META3=$(yes $PASSWORD | kaijud tx iid update-iid-meta "$NEW_DID" '{"versionID":"1","deactivated":false,"entityType":"nft","startDate":null,"endDate":null,"status":1,"stage":"yes","relayerNode":"yes","verifiableCredential":"yes","credentials":[]}'  --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y --output json)
echo $META3

echo "Querying DID..."
echo $FULL_DID
QUERY_DID=$(kaijud query iid iid "$FULL_DID" --chain-id pandora-4 --output json | jq .)

echo $QUERY_DID

echo "Changing metadata..."
META3=$(yes $PASSWORD | kaijud tx iid update-iid-meta "$NEW_DID" '{"versionID":"2","deactivated":false,"entityType":"stove","startDate":null,"endDate":null,"status":1,"stage":"yes","relayerNode":"yes","verifiableCredential":"yes","credentials":[]}'  --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y --output json)
echo "Querying IID METADATA"
QUERY_DID=$(kaijud query iid metadata "$FULL_DID" --chain-id pandora-4 --output json | jq .)
echo "Deactivating IID"
DEAC=$Fnft(yes $PASSWORD | kaijud tx iid deactivate-iid "$NEW_DID" "true"  --from miguel --from miguel --chain-id pandora-4 --fees 5000ukaiju -y --output json)
echo "Querying IID METADATA"
QUERY_DID=$(kaijud query iid metadata "$FULL_DID" --chain-id pandora-4 --output json | jq .)
echo $QUERY_DID
