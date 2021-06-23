package types_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/interchange/x/ibcdex/types"
	"math/rand"
)

func GenString(n int) string {
	alpha := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	buf := make([]rune, n)
	for i := range buf {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func GenAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func GenAmount() int32 {
	return int32(rand.Intn(int(types.MaxAmount)) + 1)
}

func GenPrice() int32 {
	return int32(rand.Intn(int(types.MaxPrice)) + 1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (string, int32, int32) {
	return GenLocalAccount(), GenAmount(), GenPrice()
}

func GenLocalAccount() string {
	return GenAddress()
}

func MockAccount(str string) string {
	return str
}

func OrderListToSellOrderBook(list []types.Order) types.SellOrderBook {
	listCopy := make([]*types.Order, len(list))
	for i, order := range list {
		order := order
		listCopy[i] = &order
	}

	book := types.SellOrderBook{
		AmountDenom:  "foo",
		PriceDenom:   "bar",
		Book: &types.OrderBook{
			IdCount: 0,
			Orders:       listCopy,
		},
	}
	return book
}

func OrderListToBuyOrderBook(list []types.Order) types.BuyOrderBook {
	listCopy := make([]*types.Order, len(list))
	for i, order := range list {
		order := order
		listCopy[i] = &order
	}

	book := types.BuyOrderBook{
		AmountDenom:  "foo",
		PriceDenom:   "bar",
		Book: &types.OrderBook{
			IdCount: 0,
			Orders:       listCopy,
		},
	}
	return book
}


