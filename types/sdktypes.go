package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/terra-money/core/x/genaccounts"
)

type (
	Coin           = sdk.Coin
	Coins          = sdk.Coins
	Int            = sdk.Int
	Dec            = sdk.Dec
	DecCoin        = sdk.DecCoin
	AccAddress     = sdk.AccAddress
	ValAddress     = sdk.ValAddress
	Msg            = sdk.Msg
	GenesisDoc     = tmtypes.GenesisDoc
	GenesisAccount = genaccounts.GenesisAccount
	HexBytes       = cmn.HexBytes
)

var (
	NewInt = sdk.NewInt
)
