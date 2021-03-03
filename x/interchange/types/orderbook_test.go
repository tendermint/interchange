package types_test

import (
	"github.com/tendermint/interchange/x/interchange/types"
	"math/rand"
)

// ------------------------------ Utils ------------------------------

func GenAmount() uint64 {
	return uint64(rand.Intn(int(types.MaxAmount)) + 1)
}

func GenPrice() uint64 {
	return uint64(rand.Intn(int(types.MaxPrice)) + 1)
}

func GenPair() (string, string) {
	return GenString(10), GenString(10)
}

func GenOrder() (types.Account, uint64, uint64) {
	return GenLocalAccount(), GenAmount(), GenPrice()
}
