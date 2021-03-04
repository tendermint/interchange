package keeper

import (
	"github.com/tendermint/interchange/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
