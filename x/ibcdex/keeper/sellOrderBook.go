package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

// SetSellOrderBook set a specific sellOrderBook in the store from its index
func (k Keeper) SetSellOrderBook(ctx sdk.Context, sellOrderBook types.SellOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKey))
	b := k.cdc.MustMarshalBinaryBare(&sellOrderBook)
	store.Set(types.KeyPrefix(sellOrderBook.Index), b)
}

// GetSellOrderBook returns a sellOrderBook from its index
func (k Keeper) GetSellOrderBook(ctx sdk.Context, index string) (val types.SellOrderBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteSellOrderBook removes a sellOrderBook from the store
func (k Keeper) RemoveSellOrderBook(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllSellOrderBook returns all sellOrderBook
func (k Keeper) GetAllSellOrderBook(ctx sdk.Context) (list []types.SellOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SellOrderBookKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SellOrderBook
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
