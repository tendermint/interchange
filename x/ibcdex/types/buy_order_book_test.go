package types_test

import (
	"testing"
	"sort"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/ibcdex/types"
)

func OrderListToBuyOrderBook(list []types.Order) types.BuyOrderBook {
	listCopy := make([]*types.Order, len(list))
	for i, order := range list {
		order := order
		listCopy[i] = &order
	}

	book := types.BuyOrderBook{
		AmountDenom: "foo",
		PriceDenom:  "bar",
		Book: &types.OrderBook{
			IdCount: 0,
			Orders:  listCopy,
		},
	}
	return book
}

func TestAppendOrder(t *testing.T) {
	buyBook := types.NewBuyOrderBook(GenPair())

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, err := buyBook.AppendOrder(seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, err = buyBook.AppendOrder(seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, err = buyBook.AppendOrder(seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, err = buyBook.AppendOrder(seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)

	// Can append buy orders
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			Id:      buyBook.Book.IdCount,
			Creator: creator,
			Amount:  amount,
			Price:   price,
		}
		orderID, err := buyBook.AppendOrder(creator, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, buyBook.Book.Orders, &newOrder)
		require.Equal(t, newOrder.Id, orderID)
	}


	require.Len(t, buyBook.Book.Orders, 20)
	require.True(t, sort.SliceIsSorted(buyBook.Book.Orders, func(i, j int) bool {
		return buyBook.Book.Orders[i].Price < buyBook.Book.Orders[j].Price
	}))
}