package ibcdex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/keeper"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the buyOrderBook
	for _, elem := range genState.BuyOrderBookList {
		k.SetBuyOrderBook(ctx, *elem)
	}

	// Set all the sellOrderBook
	for _, elem := range genState.SellOrderBookList {
		k.SetSellOrderBook(ctx, *elem)
	}

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all buyOrderBook
	buyOrderBookList := k.GetAllBuyOrderBook(ctx)
	for _, elem := range buyOrderBookList {
		elem := elem
		genesis.BuyOrderBookList = append(genesis.BuyOrderBookList, &elem)
	}

	// Get all sellOrderBook
	sellOrderBookList := k.GetAllSellOrderBook(ctx)
	for _, elem := range sellOrderBookList {
		elem := elem
		genesis.SellOrderBookList = append(genesis.SellOrderBookList, &elem)
	}

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
