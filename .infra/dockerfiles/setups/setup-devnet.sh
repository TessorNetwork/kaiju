#!/usr/bin/env bash
PASSWORD=“12345678”
# xcod init local --chain-id devnet-1
/app/xcod init local --chain-id devnet-1
# yes ‘y’ | xcod keys delete miguel --force
# yes ‘y’ | xcod keys delete francesco --force
# yes ‘y’ | xcod keys delete shaun --force
# yes ‘y’ | xcod keys delete fee --force
# yes ‘y’ | xcod keys delete fee2 --force
# yes ‘y’ | xcod keys delete fee3 --force
# yes ‘y’ | xcod keys delete fee4 --force
# yes ‘y’ | xcod keys delete fee5 --force
# yes ‘y’ | xcod keys delete reserveOut --force
# yes $PASSWORD | xcod keys add miguel
# yes $PASSWORD | xcod keys add francesco
# yes $PASSWORD | xcod keys add shaun
# yes $PASSWORD | xcod keys add fee
# yes $PASSWORD | xcod keys add fee2
# yes $PASSWORD | xcod keys add fee3
# yes $PASSWORD | xcod keys add fee4
# yes $PASSWORD | xcod keys add fee5
# yes $PASSWORD | xcod keys add reserveOut
yes 12345678 | /app/xcod keys add dev-main
# Note: important to add ‘miguel’ as a genesis-account since this is the chain’s validator
yes 12345678 | /app/xcod add-genesis-account $(/app/xcod keys show dev-main -a) 100000000000000000uxco

# yes $PASSWORD | xcod add-genesis-account “$(xcod keys show miguel -a)” 1000000000000uxco,1000000000000res,1000000000000rez,1000000000000uxgbp
# yes $PASSWORD | xcod add-genesis-account “$(xcod keys show francesco -a)” 1000000000000uxco,1000000000000res,1000000000000rez
# yes $PASSWORD | xcod add-genesis-account “$(xcod keys show shaun -a)” 1000000000000uxco,1000000000000res,1000000000000rez
# Add pubkey-based genesis accounts
# MIGUEL_ADDR=“xco1acltgu0kwgnuqdgewracms3nhz8c6n2grk0uz0”    # address from did:xco:4XJLBfGtWSGKSz4BeRxdun’s pubkey
# FRANCESCO_ADDR=“xco1zyaz6rkpxa9mdlzazc9uuch4hqc7l5eatsunes” # address from did:xco:UKzkhVSHc3qEFva5EY2XHt’s pubkey
# SHAUN_ADDR=“xco10uqnjz60h3lkxxsgnlvxql3yfpkgv6gc6wuz4c”     # address from did:xco:U4tSpzzv91HHqWW1YmFkHJ’s pubkey
# yes $PASSWORD | xcod add-genesis-account “$MIGUEL_ADDR” 1000000000000uxco,1000000000000res,1000000000000rez
# yes $PASSWORD | xcod add-genesis-account “$FRANCESCO_ADDR” 1000000000000uxco,1000000000000res,1000000000000rez
# yes $PASSWORD | xcod add-genesis-account “$SHAUN_ADDR” 1000000000000uxco,1000000000000res,1000000000000rez
# yes $PASSWORD | xcod add-genesis-account “xco14zqgpctk6hnnms77427gyahpawnmj5ajxtdc0c” 1000000000000uxco,1000000000000res,1000000000000rez
# Add xco did
yes 12345678 | /app/xcod gentx dev-main 10000000000uxco --chain-id devnet-1
HOME=/root
/app/xcod collect-gentxs

# XCO_DID=“did:xco:U4tSpzzv91HHqWW1YmFkHJ”
# export FROM=\”xco_did\“: \“\”
# export TO=“\”xco_did\“: \“$XCO_DID\“”
sed -i 's/"xco_did" : ""/"xco_did" : ""/;s/"bond_denom" : "stake"/"bond_denom" : "uxco"/;s/"mint_denom" : "stake"/"mint_denom" : "uxco"/;s/stake/uxco/;s/"Reserved_bond_tokens" : "\[\]"/"Reserved_bond_tokens" : "\[\]"/;s/"minimum-gas-prices" : ""/"minimum-gas-prices" : "0.025uxco"/;s/"enable" : "false"/"enable" : "true"/;s/"swagger" : "false"/"swagger" : "true"/;' $HOME/.xcod/config/genesis.json
MAX_VOTING_PERIOD="30s"  # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.xcod/config/genesis.json
# sed -i '/xco_did/c\   \"i xo_did\" : \"did:xco:aaaaaaaaaaa\",' $HOME/.xcod/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
# export STAKING_TOKEN=“uxco”
# export FROM=“\”bond_denom\“: \“stake\“”
# export TO=“\”bond_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/bond_denom/c\   \"bond_denom\" : \"uxco\",' $HOME/.xcod/config/genesis.json

# export FROM=“\”mint_denom\“: \“stake\“”
# export TO=“\”mint_denom\“: \“$STAKING_TOKEN\“”
# sed -i '/mint_denom/c\   \"mint_denom\" : \"uxco\",' $HOME/.xcod/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
# export FEE_TOKEN=“uxco”
# export FROM=“\”stake\“”
# export TO=“\”$FEE_TOKEN\“”
# sed 's/stake/uxco' $HOME/.xcod/config/genesis.json

# Set reserved bond tokens
# export RESERVED_BOND_TOKENS=“”  # example: ” \“abc\“, \“def\“, \“ghi\” ”
# export FROM=“\”reserved_bond_tokens\“: \[\]”
# export TO=“\”reserved_bond_tokens\“: \[$RESERVED_BOND_TOKENS\]”
# sed -i '/reserved_bond_tokens/c\   \"Reserved_bond_tokens\" : \"[]\",' $HOME/.xcod/config/genesis.json

# Set min-gas-prices (using fee token)
# export FROM=“minimum-gas-prices = \“\”
# export TO=“minimum-gas-prices = \“0.025$FEE_TOKEN\“”
# sed -i '/minimum-gas-prices/c\   \"minimum-gas-prices\" : \"0.025uxco\",' $HOME/.xcod/config/genesis.json
# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
#xcod config chain-id devnet-1
#xcod config output jsonW
#xcod config indent true
#xcod config trust-node true
# xcod gentx miguel 1000000uxco --chain-id devnet-1
/app/xcod validate-genesis
# Enable REST API (assumed to be at line 104 of app.toml)
# export FROM=“enable = false”
# TO=“enable = true”
# sed -i '/enable/c\   enable = true' $HOME/.xcod/config/genesis.json
# Enable Swagger docs (assumed to be at line 107 of app.toml)
# export FROM=“swagger = false”
# export TO=“swagger = true”
# sed -i '/swagger/c\   swagger = true' $HOME/.xcod/config/genesis.json
# Uncomment the below to broadcast node RPC endpoint
#FROM=“laddr = \“tcp:\/\/127.0.0.1:26657\“”
#TO=“laddr = \“tcp:\/\/0.0.0.0:26657\“”
#sed -i “s/$FROM/$TO/” “$HOME”/.xcod/config/config.toml
# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i ‘s/timeout_commit = “5s”/timeout_commit = “1s”/g’ “$HOME”/.xcod/config/config.toml
#sed -i ‘s/timeout_propose = “3s”/timeout_propose = “1s”/g’ “$HOME”/.xcod/config/config.toml
# xcod start --pruning “nothing”