package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

// SetBuyOrderBook set a specific buyOrderBook in the store from its index
func (k Keeper) SetBuyOrderBook(ctx sdk.Context, buyOrderBook types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKey))
	b := k.cdc.MustMarshalBinaryBare(&buyOrderBook)
	store.Set(types.KeyPrefix(buyOrderBook.Index), b)
}

// GetBuyOrderBook returns a buyOrderBook from its index
func (k Keeper) GetBuyOrderBook(ctx sdk.Context, index string) (val types.BuyOrderBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteBuyOrderBook removes a buyOrderBook from the store
func (k Keeper) RemoveBuyOrderBook(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllBuyOrderBook returns all buyOrderBook
func (k Keeper) GetAllBuyOrderBook(ctx sdk.Context) (list []types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuyOrderBook
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
