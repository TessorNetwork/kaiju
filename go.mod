module github.com/tessornetwork/kaiju

go 1.18

require (
	github.com/CosmWasm/wasmd v0.29.1
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.45.9
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ibc-go/v3 v3.3.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.13.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.21
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require github.com/regen-network/cosmos-proto v0.3.1

require (
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
	golang.org/x/net v0.3.0 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0

	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.0

	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

	// latest grpc doesn't work with with our modified proto compiler, so we need to enforce
	// the following version across all dependencies.
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
