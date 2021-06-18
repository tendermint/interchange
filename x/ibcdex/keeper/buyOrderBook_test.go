package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/interchange/x/ibcdex/types"
)

func createNBuyOrderBook(keeper *Keeper, ctx sdk.Context, n int) []types.BuyOrderBook {
	items := make([]types.BuyOrderBook, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetBuyOrderBook(ctx, items[i])
	}
	return items
}

func TestBuyOrderBookGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBuyOrderBook(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestBuyOrderBookRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBuyOrderBook(ctx, item.Index)
		_, found := keeper.GetBuyOrderBook(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestBuyOrderBookGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBuyOrderBook(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllBuyOrderBook(ctx))
}
