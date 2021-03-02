package types_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
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