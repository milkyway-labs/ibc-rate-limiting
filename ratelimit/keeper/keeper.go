package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
)

type (
	Keeper struct {
		storeKey  storetypes.StoreKey
		cdc       codec.BinaryCodec
		authority string

		bankKeeper    types.BankKeeper
		channelKeeper types.ChannelKeeper
		ics4Wrapper   types.ICS4Wrapper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	authority string,
	bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
	ics4Wrapper types.ICS4Wrapper,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      key,
		authority:     authority,
		bankKeeper:    bankKeeper,
		channelKeeper: channelKeeper,
		ics4Wrapper:   ics4Wrapper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
