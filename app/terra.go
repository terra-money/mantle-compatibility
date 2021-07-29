package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cast"
	"io/ioutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	terra "github.com/terra-money/core/app"
	wasmconfig "github.com/terra-money/core/x/wasm/config"
)

func NewTerraApp(db dbm.DB) *terra.TerraApp {
	return terra.NewTerraApp(
		log.NewTMLogger(ioutil.Discard),
		db,
		nil,
		true, // need this so KVStores are set
		make(map[int64]bool),
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		terra.MakeEncodingConfig(), // Ideally, we would reuse the one created by NewRootCmd.
		appOpts,
		&wasmconfig.Config{BaseConfig: wasmconfig.BaseConfig{
			ContractQueryGasLimit:   uint64(30000000),
			ContractDebugMode: false,
			ContractMemoryCacheSize: cast.ToUint32(appOpts.Get(wasmconfig.FlagContractMemoryCacheSize)),
		}},
		setPruningOptions(),
		fauxMerkleModeOpt,
		// baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		// baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		// baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		// baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		// baseapp.SetInterBlockCache(cache),
		// baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		// baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
	)
}

// Pass this in as an option to use a dbStoreAdapter instead of an IAVLStore for simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

func setPruningOptions() func(*baseapp.BaseApp) {
	// prune nothing
	pruningOptions := sdk.PruningOptions{
		KeepRecent: 0,
		KeepEvery:  0,
		Interval:   10,
	}
	return baseapp.SetPruning(pruningOptions)
}
