#!/usr/bin/env bash

# IT IS RECOMMENDED TO RUN THE BLOCKCHAIN USING run_with_all_data_dev.sh SINCE
# THIS SETS GOVERNANCE PERIODS TO 30 seconds FOR FASTER GOVERNANCE

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

echo "Query transfer params before param change"
kaijud_q ibc-transfer params

echo "Submitting param change proposal"
kaijud_tx gov submit-proposal param-change demo_gov_param_change_proposal.json --from=miguel

echo "Query proposal 1"
kaijud_q gov proposal 1

echo "Depositing 10000000ukaiju to reach minimum deposit"
kaijud_tx gov deposit 1 10000000ukaiju --from=miguel

echo "Query proposal 1 deposits"
kaijud_q gov deposits 1

echo "Voting yes for proposal"
kaijud_tx gov vote 1 yes --from=miguel

echo "Query proposal 1 tally"
kaijud_q gov tally 1

echo "Waiting for proposal to pass..."
while :; do
  RET=$(kaijud_q gov proposal 1 2>&1)
  if [[ ($RET == *'PROPOSAL_STATUS_VOTING_PERIOD'*) ]]; then
    sleep 1
  else
    echo "A few more seconds..."
    sleep 6
    break
  fi
done

echo "Query transfer params (expected to be true and false)"
kaijud_q ibc-transfer params
