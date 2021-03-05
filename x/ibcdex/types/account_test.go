package types_test

import "github.com/tendermint/interchange/x/ibcdex/types"

func GenLocalAccount() *types.Account {
	return &types.Account{
		Address: GenAddress(),
		Remote:   false,
	}
}

func GenRemoteAccount() *types.Account {
	return &types.Account{
		Address: GenAddress(),
		Remote:   true,
	}
}

func MockAccount(str string) *types.Account {
	return &types.Account{
		Address: str,
		Remote: false,
	}
}
