#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(xcod status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(xcod status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
GAS_PRICES="0.025uxco"
CHAIN_ID="pandora-4"
FEE=$(yes $PASSWORD | xcod keys show fee -a)
RESERVE_OUT=$(yes $PASSWORD | xcod keys show reserveOut -a)

xcod_tx() {
  # This function first approximates the gas (adjusted to 105%) and then
  # supplies this for the actual transaction broadcasting as the --gas.
  # This might fail sometimes: https://github.com/cosmos/cosmos-sdk/issues/4938
  cmd="$1 $2"
  shift
  shift
  # APPROX=$(xcod tx $cmd --gas=auto --gas-adjustment=1.05 --fees=1uxco --chain-id="$CHAIN_ID" --dry-run "$@" 2>&1)
  # APPROX=${APPROX//gas estimate: /}
  echo "Gas estimate: $APPROX"
  yes "12345678" | xcod tx $cmd \
    --fees 5000uxco \
    --chain-id="$CHAIN_ID" \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

xcod_q() {
  xcod q "$@" --output=json | jq .
}


MIGUEL_DID_FULL='{
  "did":"did:xco:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'

FRANCESCO_DID="did:xco:UKzkhVSHc3qEFva5EY2XHt"
FRANCESCO_DID_FULL='{
  "did":"did:xco:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'

INST='{
"minter": "xco14nevcuw8sfz5ltsq4f6x4fr56cvlhcklraucvn",
}'

#xcod tx wasm instantiate 2 '{"minter":"xco14nevcuw8sfz5ltsq4f6x4fr56cvlhcklraucvn"}' --label abc --from miguel --admin "xco14nevcuw8sfz5ltsq4f6x4fr56cvlhcklraucvn" $TXFLAG
# echo $ENTITY | jq

echo "Minting 25x abc123/CARBON"
xcod_tx wasm execute 'xco1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqg98jca' '{"mint":{"to":"xco14nevcuw8sfz5ltsq4f6x4fr56cvlhcklraucvn","token_id":"CARBON:abc124","value":"1","uri":"did:xco:entity12345"}}' --from miguel
echo "Getting balance"
xcod_q wasm contract-state smart 'xco1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqg98jca' '{"balance":{"owner":"xco14nevcuw8sfz5ltsq4f6x4fr56cvlhcklraucvn","token_id":"CARBON:abc124"}}'
echo "Getting URI (Token Info)"
xcod_q wasm contract-state smart 'xco1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqg98jca' '{"token_info":{"token_id":"CARBON:abc124"}}'

