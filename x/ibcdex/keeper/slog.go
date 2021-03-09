package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
	"strconv"
)

// GetSlogCount get the total number of slog
func (k Keeper) GetSlogCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogCountKey))
	byteKey := types.KeyPrefix(types.SlogCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetSlogCount set the total number of slog
func (k Keeper) SetSlogCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogCountKey))
	byteKey := types.KeyPrefix(types.SlogCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendSlog appends a slog in the store with a new id and update the count
func (k Keeper) AppendSlog(
	ctx sdk.Context,
	creator string,
	log string,
) string {
	// Create the slog
	count := k.GetSlogCount(ctx)
	id := strconv.FormatInt(count, 10)
	var slog = types.Slog{
		Creator: creator,
		Id:      id,
		Log:     log,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	key := types.KeyPrefix(types.SlogKey + slog.Id)
	value := k.cdc.MustMarshalBinaryBare(&slog)
	store.Set(key, value)

	// Update slog count
	k.SetSlogCount(ctx, count+1)

	return id
}

// SetSlog set a specific slog in the store
func (k Keeper) SetSlog(ctx sdk.Context, slog types.Slog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	b := k.cdc.MustMarshalBinaryBare(&slog)
	store.Set(types.KeyPrefix(types.SlogKey+slog.Id), b)
}

// GetSlog returns a slog from its id
func (k Keeper) GetSlog(ctx sdk.Context, key string) types.Slog {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	var slog types.Slog
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SlogKey+key)), &slog)
	return slog
}

// HasSlog checks if the slog exists in the store
func (k Keeper) HasSlog(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	return store.Has(types.KeyPrefix(types.SlogKey + id))
}

// GetSlogOwner returns the creator of the slog
func (k Keeper) GetSlogOwner(ctx sdk.Context, key string) string {
	return k.GetSlog(ctx, key).Creator
}

// DeleteSlog removes a slog from the store
func (k Keeper) RemoveSlog(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	store.Delete(types.KeyPrefix(types.SlogKey + key))
}

// GetAllSlog returns all slog
func (k Keeper) GetAllSlog(ctx sdk.Context) (list []types.Slog) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SlogKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Slog
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
