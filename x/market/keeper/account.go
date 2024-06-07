package keeper

import (
	"github.com/ollo-station/ollo/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetMarketplaceAccount returns marketplace ModuleAccount
func (k Keeper) GetMarketplaceAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}
