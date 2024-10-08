package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ollo-station/ollo/x/exchange/types"
)

// Keeper of the module's store.
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	paramSpace paramstypes.Subspace

	bankKeeper  types.BankKeeper
	sources     map[string]types.OrderSource
	sourceNames []string
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramSpace paramstypes.Subspace,
	bankKeeper types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) SetOrderSources(sources ...types.OrderSource) *Keeper {
	if k.sources != nil {
		panic("cannot set order sources twice")
	}
	k.sources = map[string]types.OrderSource{}
	for _, source := range sources {
		sourceName := source.Name()
		if _, ok := k.sources[sourceName]; ok {
			panic(fmt.Sprintf("duplicate order source name: %s", sourceName))
		}
		k.sources[sourceName] = source
		k.sourceNames = append(k.sourceNames, sourceName)
	}
	return k
}
