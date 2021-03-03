package types_test

import "github.com/tendermint/interchange/x/interchange/types"

func GenLocalAccount() types.Account {
	return types.Account{
		Address: GenAddress(),
		Chain:   "",
	}
}

func GenRemoteAccount() types.Account {
	return types.Account{
		Address: GenAddress(),
		Chain:   GenString(10),
	}
}
