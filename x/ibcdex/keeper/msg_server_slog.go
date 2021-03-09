package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) CreateSlog(goCtx context.Context, msg *types.MsgCreateSlog) (*types.MsgCreateSlogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendSlog(
		ctx,
		msg.Creator,
		msg.Log,
	)

	return &types.MsgCreateSlogResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSlog(goCtx context.Context, msg *types.MsgUpdateSlog) (*types.MsgUpdateSlogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var slog = types.Slog{
		Creator: msg.Creator,
		Id:      msg.Id,
		Log:     msg.Log,
	}

	// Checks that the element exists
	if !k.HasSlog(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetSlogOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSlog(ctx, slog)

	return &types.MsgUpdateSlogResponse{}, nil
}

func (k msgServer) DeleteSlog(goCtx context.Context, msg *types.MsgDeleteSlog) (*types.MsgDeleteSlogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasSlog(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetSlogOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSlog(ctx, msg.Id)

	return &types.MsgDeleteSlogResponse{}, nil
}
