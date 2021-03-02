package keeper

import (
	"github.com/tendermint/interchange/x/interchange/types"
)

var _ types.QueryServer = Keeper{}
