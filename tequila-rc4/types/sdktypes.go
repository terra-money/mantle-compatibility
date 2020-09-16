package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/tendermint/tendermint/libs/bytes"
	tmtypes "github.com/tendermint/tendermint/types"
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
	GenesisAccount = authexported.GenesisAccount
	HexBytes       = bytes.HexBytes
)

var (
	NewInt = sdk.NewInt
)
