package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/interchange/x/ibcdex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SlogAll(c context.Context, req *types.QueryAllSlogRequest) (*types.QueryAllSlogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var slogs []*types.Slog
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	slogStore := prefix.NewStore(store, types.KeyPrefix(types.SlogKey))

	pageRes, err := query.Paginate(slogStore, req.Pagination, func(key []byte, value []byte) error {
		var slog types.Slog
		if err := k.cdc.UnmarshalBinaryBare(value, &slog); err != nil {
			return err
		}

		slogs = append(slogs, &slog)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSlogResponse{Slog: slogs, Pagination: pageRes}, nil
}

func (k Keeper) Slog(c context.Context, req *types.QueryGetSlogRequest) (*types.QueryGetSlogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var slog types.Slog
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SlogKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SlogKey+req.Id)), &slog)

	return &types.QueryGetSlogResponse{Slog: &slog}, nil
}
