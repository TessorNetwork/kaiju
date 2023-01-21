#!/usr/bin/env bash
PASSWORD=“12345678”
# kaijud init local --chain-id devnet-1
/app/kaijud init local --chain-id devnet-1
# yes ‘y’ | kaijud keys delete miguel --force
# yes ‘y’ | kaijud keys delete francesco --force
# yes ‘y’ | kaijud keys delete shaun --force
# yes ‘y’ | kaijud keys delete fee --force
# yes ‘y’ | kaijud keys delete fee2 --force
# yes ‘y’ | kaijud keys delete fee3 --force
# yes ‘y’ | kaijud keys delete fee4 --force
# yes ‘y’ | kaijud keys delete fee5 --force
# yes ‘y’ | kaijud keys delete reserveOut --force
# yes $PASSWORD | kaijud keys add miguel
# yes $PASSWORD | kaijud keys add francesco
# yes $PASSWORD | kaijud keys add shaun
# yes $PASSWORD | kaijud keys add fee
# yes $PASSWORD | kaijud keys add fee2
# yes $PASSWORD | kaijud keys add fee3
# yes $PASSWORD | kaijud keys add fee4
# yes $PASSWORD | kaijud keys add fee5
# yes $PASSWORD | kaijud keys add reserveOut
yes 12345678 | /app/kaijud keys add dev-main
# Note: important to add ‘miguel’ as a genesis-account since this is the chain’s validator
yes 12345678 | /app/kaijud add-genesis-account $(/app/kaijud keys show dev-main -a) 100000000000000000ukaiju

# yes $PASSWORD | kaijud add-genesis-account “$(kaijud keys show miguel -a)” 1000000000000ukaiju,1000000000000res,1000000000000rez,1000000000000uxgbp
# yes $PASSWORD | kaijud add-genesis-account “$(kaijud keys show francesco -a)” 1000000000000ukaiju,1000000000000res,1000000000000rez
# yes $PASSWORD | kaijud add-genesis-account “$(kaijud keys show shaun -a)” 1000000000000ukaiju,1000000000000res,1000000000000rez
# Add pubkey-based genesis accounts
# MIGUEL_ADDR=“kaiju1acltgu0kwgnuqdgewracms3nhz8c6n2g3wfakx”    # address from did:kaiju:4XJLBfGtWSGKSz4BeRxdun’s pubkey
# FRANCESCO_ADDR=“kaiju1zyaz6rkpxa9mdlzazc9uuch4hqc7l5eaeg6jde” # address from did:kaiju:UKzkhVSHc3qEFva5EY2XHt’s pubkey
# SHAUN_ADDR=“kaiju10uqnjz60h3lkxxsgnlvxql3yfpkgv6gcgk6rp3”     # address from did:kaiju:U4tSpzzv91HHqWW1YmFkHJ’s pubkey
# yes $PASSWORD | kaijud add-genesis-account “$MIGUEL_ADDR” 1000000000000ukaiju,1000000000000res,1000000000000rez
# yes $PASSWORD | kaijud add-genesis-account “$FRANCESCO_ADDR” 1000000000000ukaiju,1000000000000res,1000000000000rez
# yes $PASSWORD | kaijud add-genesis-account “$SHAUN_ADDR” 1000000000000ukaiju,1000000000000res,1000000000000rez
# yes $PASSWORD | kaijud add-genesis-account “kaiju14zqgpctk6hnnms77427gyahpawnmj5aj5ntem3” 1000000000000ukaiju,1000000000000res,1000000000000rez
# Add kaiju did
yes 12345678 | /app/kaijud gentx dev-main 10000000000ukaiju --chain-id devnet-1
HOME=/root
/app/kaijud collect-gentxs

# KAIJU_DID=“did:kaiju:U4tSpzzv91HHqWW1YmFkHJ”
# export FROM=\”kaiju_did\“: \“\”
# export TO=“\”kaiju_did\“: \“$KAIJU_DID\“”
sed -i 's/"kaiju_did" : ""/"kaiju_did" : ""/;s/"bond_denom" : "stake"/"bond_denom" : "ukaiju"/;s/"mint_denom" : "stake"/"mint_denom" : "ukaiju"/;s/stake/ukaiju/;s/"Reserved_bond_tokens" : "\[\]"/"Reserved_bond_tokens" : "\[\]"/;s/"minimum-gas-prices" : ""/"minimum-gas-prices" : "0.025ukaiju"/;s/"enable" : "false"/"enable" : "true"/;s/"swagger" : "false"/"swagger" : "true"/;' $HOME/.kaijud/config/genesis.json
MAX_VOTING_PERIOD="30s"  # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.kaijud/config/genesis.json
# sed -i '/kaiju_did/c\   \"i xo_did\" : \"did:kaiju:aaaaaaaaaaa\",' $HOME/.kaijud/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
# export STAKING_TOKEN=“ukaiju”
# export FROM=“\”bond_denom\“: \“stake\“”
# export TO=“\”bond_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/bond_denom/c\   \"bond_denom\" : \"ukaiju\",' $HOME/.kaijud/config/genesis.json

# export FROM=“\”mint_denom\“: \“stake\“”
# export TO=“\”mint_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/mint_denom/c\   \"mint_denom\" : \"ukaiju\",' $HOME/.kaijud/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
# export FEE_TOKEN=“ukaiju”
# export FROM=“\”stake\“”
# export TO=“\”$FEE_TOKEN\“”
# sed 's/stake/ukaiju' $HOME/.kaijud/config/genesis.json

# Set reserved bond tokens
# export RESERVED_BOND_TOKENS=“”  # example: ” \“abc\“, \“def\“, \“ghi\” ”
# export FROM=“\”reserved_bond_tokens\“: \[\]”
# export TO=“\”reserved_bond_tokens\“: \[$RESERVED_BOND_TOKENS\]”
# sed -i '/reserved_bond_tokens/c\   \"Reserved_bond_tokens\" : \"[]\",' $HOME/.kaijud/config/genesis.json

# Set min-gas-prices (using fee token)
# export FROM=“minimum-gas-prices = \“\”
# export TO=“minimum-gas-prices = \“0.025$FEE_TOKEN\“”
# sed -i '/minimum-gas-prices/c\   \"minimum-gas-prices\" : \"0.025ukaiju\",' $HOME/.kaijud/config/genesis.json
# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
#kaijud config chain-id devnet-1
#kaijud config output jsonW
#kaijud config indent true
#kaijud config trust-node true
# kaijud gentx miguel 1000000ukaiju --chain-id devnet-1
/app/kaijud validate-genesis
# Enable REST API (assumed to be at line 104 of app.toml)
# export FROM=“enable = false”
# TO=“enable = true”
# sed -i '/enable/c\   enable = true' $HOME/.kaijud/config/genesis.json
# Enable Swagger docs (assumed to be at line 107 of app.toml)
# export FROM=“swagger = false”
# export TO=“swagger = true”
# sed -i '/swagger/c\   swagger = true' $HOME/.kaijud/config/genesis.json
# Uncomment the below to broadcast node RPC endpoint
#FROM=“laddr = \“tcp:\/\/127.0.0.1:26657\“”
#TO=“laddr = \“tcp:\/\/0.0.0.0:26657\“”
#sed -i “s/$FROM/$TO/” “$HOME”/.kaijud/config/config.toml
# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i ‘s/timeout_commit = “5s”/timeout_commit = “1s”/g’ “$HOME”/.kaijud/config/config.toml
#sed -i ‘s/timeout_propose = “3s”/timeout_propose = “1s”/g’ “$HOME”/.kaijud/config/config.toml
# kaijud start --pruning “nothing”