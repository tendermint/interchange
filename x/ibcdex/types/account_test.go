package types_test

import "github.com/tendermint/interchange/x/interchange/types"

func GenLocalAccount() types.Account {
	return types.Account{
		Address: types.GenAddress(),
		Chain:   "",
	}
}

func GenRemoteAccount() types.Account {
	return types.Account{
		Address: types.GenAddress(),
		Chain:   types.GenString(10),
	}
}

func MockAccount(str string) types.Account {
	return types.Account{
		Address: str,
		Chain:   "",
	}
}
