package types_test

import (
	"github.com/tendermint/interchange/x/ibcdex/types"
	"testing"
	"github.com/stretchr/testify/require"
	"sort"
)

func TestSellOrderBook_AppendOrder(t *testing.T) {
	sellBook := types.NewSellOrderBook(GenPair())

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, err := sellBook.AppendOrder(seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, err = sellBook.AppendOrder(seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, err = sellBook.AppendOrder(seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, err = sellBook.AppendOrder(seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)

	// Can append sell orders
	for i := 0; i < 20; i++ {
		// Append a new order
		creator, amount, price := GenOrder()
		newOrder := types.Order{
			Id:      sellBook.Book.IdCount,
			Creator: creator,
			Amount:  amount,
			Price:   price,
		}
		orderID, err := sellBook.AppendOrder(creator, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, sellBook.Book.Orders, &newOrder)
		require.Equal(t, newOrder.Id, orderID)
	}
	require.Len(t, sellBook.Book.Orders, 20)
	require.True(t, sort.SliceIsSorted(sellBook.Book.Orders, func(i, j int) bool {
		return sellBook.Book.Orders[i].Price > sellBook.Book.Orders[j].Price
	}))
}