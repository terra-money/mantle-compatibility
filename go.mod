module github.com/terra-money/mantle-compatibility

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	github.com/terra-money/core v0.5.0-beta4
	github.com/CosmWasm/wasmvm v0.15.1

)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
