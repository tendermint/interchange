package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func (k msgServer) CancelBuyOrder(goCtx context.Context, msg *types.MsgCancelBuyOrder) (*types.MsgCancelBuyOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the book
	pairIndex := types.OrderBookIndex(msg.Port, msg.Channel, msg.AmountDenom, msg.PriceDenom)
	book, found := k.GetBuyOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgCancelBuyOrderResponse{}, errors.New("the pair doesn't exist")
	}

	// Check order creator
	order, err := book.GetOrderFromID(msg.OrderID)
	if err != nil {
		return &types.MsgCancelBuyOrderResponse{}, err
	}
	if order.Creator != msg.Creator {
		return &types.MsgCancelBuyOrderResponse{}, errors.New("canceller must be creator")
	}

	// Remove order
	newBook, err := book.RemoveOrderFromID(msg.OrderID)
	if err != nil {
		return &types.MsgCancelBuyOrderResponse{}, err
	}
	book = newBook.(types.BuyOrderBook)
	k.SetBuyOrderBook(ctx, book)

	// Refund buyer with remaining price amount
	buyer, err := sdk.AccAddressFromBech32(order.Creator)
	if err != nil {
		return &types.MsgCancelBuyOrderResponse{}, err
	}
	if err := k.SafeMint(
		ctx,
		msg.Port,
		msg.Channel,
		buyer,
		msg.PriceDenom,
		order.Amount*order.Price,
	); err != nil {
		return &types.MsgCancelBuyOrderResponse{}, err
	}


	return &types.MsgCancelBuyOrderResponse{}, nil
}
