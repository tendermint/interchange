package types_test

import "github.com/tendermint/interchange/x/ibcdex/types"

func GenLocalAccount() *types.Account {
	return &types.Account{
		Address: GenAddress(),
		Chain:   "",
	}
}

func GenRemoteAccount() *types.Account {
	return &types.Account{
		Address: GenAddress(),
		Chain:   GenString(10),
	}
}

func MockAccount(str string) *types.Account {
	return &types.Account{
		Address: str,
		Chain:   "",
	}
}
