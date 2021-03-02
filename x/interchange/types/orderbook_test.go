package types_test

import (
	"github.com/tendermint/interchange/x/interchange/types"
	"github.com/stretchr/testify/require"
	"sort"
	"math/rand"
	"testing"
)

// ------------------------------ Utils ------------------------------

func GenAmount() uint64 {
	return uint64(rand.Intn(int(types.MaxAmount))+1)
}

func GenPrice() uint64 {
	return uint64(rand.Intn(int(types.MaxPrice))+1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (types.Account, uint64, uint64)  {
	return GenLocalAccount(), GenAmount(), GenPrice()
}

// ------------------------------ Sell Order ------------------------------

func TestNewSellOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	sob := types.NewSellOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint64(0), sob.OrderIDTrack)
	require.Equal(t, amountDenom, sob.AmountDenom)
	require.Equal(t, priceDenom, sob.PriceDenom)
	require.Empty(t, sob.Orders)
}

func TestAppendSellOrder(t *testing.T) {
	sob := types.NewSellOrderBook(GenPair())

	for i := 0; i < 20; i++ {
		// Append a new sell order
		seller, amount, price := GenOrder()
		newOrder := types.SellOrder{
			ID: sob.OrderIDTrack,
			Seller: seller,
			Amount: amount,
			Price: price,
		}
		newSOB, orderID, err := types.AppendSellOrder(sob, seller, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, newSOB.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)

		sob = newSOB
	}
	require.Len(t, sob.Orders, 20)
	require.True(t, sort.IsSorted(sob))

	// Prevent zero amount
	seller, amount, price := GenOrder()
	_, _, err := types.AppendSellOrder(sob, seller, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendSellOrder(sob, seller, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendSellOrder(sob, seller, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendSellOrder(sob, seller, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)
}

// ------------------------------ Buy Order ------------------------------

func TestNewBuyOrderBook(t *testing.T) {
	amountDenom, priceDenom := GenPair()
	bob := types.NewBuyOrderBook(amountDenom, priceDenom)
	require.Equal(t, uint64(1), bob.OrderIDTrack)
	require.Equal(t, amountDenom, bob.AmountDenom)
	require.Equal(t, priceDenom, bob.PriceDenom)
	require.Empty(t, bob.Orders)
}

func TestAppendBuyOrder(t *testing.T) {
	bob := types.NewBuyOrderBook(GenPair())

	for i := 0; i < 20; i++ {
		// Append a new sell order
		buyer, amount, price := GenOrder()
		newOrder := types.BuyOrder{
			ID: bob.OrderIDTrack,
			Buyer: buyer,
			Amount: amount,
			Price: price,
		}
		newBOB, orderID, err := types.AppendBuyOrder(bob, buyer, amount, price)

		// Checks
		require.NoError(t, err)
		require.Contains(t, newBOB.Orders, newOrder)
		require.Equal(t, newOrder.ID, orderID)

		bob = newBOB
	}
	require.Len(t, bob.Orders, 20)
	require.True(t, sort.IsSorted(bob))

	// Prevent zero amount
	buyer, amount, price := GenOrder()
	_, _, err := types.AppendBuyOrder(bob, buyer, 0, price)
	require.ErrorIs(t, err, types.ErrZeroAmount)

	// Prevent big amount
	_, _, err = types.AppendBuyOrder(bob, buyer, types.MaxAmount+1, price)
	require.ErrorIs(t, err, types.ErrMaxAmount)

	// Prevent zero price
	_, _, err = types.AppendBuyOrder(bob, buyer, amount, 0)
	require.ErrorIs(t, err, types.ErrZeroPrice)

	// Prevent big price
	_, _, err = types.AppendBuyOrder(bob, buyer, amount, types.MaxPrice+1)
	require.ErrorIs(t, err, types.ErrMaxPrice)
}