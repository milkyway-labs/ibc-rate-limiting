package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}
