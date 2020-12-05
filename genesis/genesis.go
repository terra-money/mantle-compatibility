package genesis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/terra-project/core/app"
	core "github.com/terra-project/core/types"
	. "github.com/terra-project/mantle-compatibility/types"
	"time"
)

func NewGenesis(chainId string, genesisAccounts ...GenesisAccount) *GenesisDoc {
	codec := app.MakeCodec()
	appStates := app.ModuleBasics.DefaultGenesis()

	authDefaultState := authtypes.DefaultGenesisState()
	authDefaultState.Accounts = genesisAccounts

	appStates["auth"] = codec.MustMarshalJSON(authDefaultState)
	appStatesJson, err := codec.MarshalJSON(appStates)

	if err != nil {
		panic(fmt.Errorf("could not initialize appstate, %v", err))
	}

	gendoc := &GenesisDoc{
		ChainID:     chainId,
		Validators:  nil,
		AppState:    appStatesJson,
		GenesisTime: time.Now(),
	}

	if gendocErr := gendoc.ValidateAndComplete(); gendocErr != nil {
		panic(gendocErr)
	}

	return gendoc
}

func NewGenesisAccount(
	address AccAddress,
	coins Coins,
) GenesisAccount {
	config := sdk.GetConfig()
	config.SetCoinType(core.CoinType)
	config.SetFullFundraiserPath(core.FullFundraiserPath)
	config.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)

	acc := authtypes.NewBaseAccountWithAddress(address)
	acc.Coins = coins

	var genAcc authexported.GenesisAccount
	genAcc = &acc

	return genAcc
}
