package app

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	terra "github.com/terra-money/core/app"
)

func NewTerraApp(db dbm.DB) *terra.TerraApp {
	return terra.NewTerraApp(
		log.NewTMLogger(ioutil.Discard),
		db,
		nil,
		true, // need this so KVStores are set
		0,
		// make(map[int64]bool),
		// &wasmconfig.Config{BaseConfig: wasmconfig.BaseConfig{
		// 	ContractQueryGasLimit: viper.GetUint64(wasmconfig.FlagContractQueryGasLimit),
		// 	CacheSize:             viper.GetUint64(wasmconfig.FlagCacheSize),
		// }},
		// fauxMerkleModeOpt, // error
		// setPruningOptions(),
	)
}

// Pass this in as an option to use a dbStoreAdapter instead of an IAVLStore for simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}
