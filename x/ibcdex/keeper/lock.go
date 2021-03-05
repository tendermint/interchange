package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
)

func (k Keeper) LockTokens(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	sender sdk.AccAddress,
	tokens sdk.Coin,
) error {
	// create the escrow address for the tokens
	escrowAddress := ibctransfertypes.GetEscrowAddress(sourcePort, sourceChannel)

	// escrow source tokens. It fails if balance insufficient
	if err := k.bankKeeper.SendCoins(
		ctx, sender, escrowAddress, sdk.NewCoins(tokens),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) UnlockTokens(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	receiver sdk.AccAddress,
	tokens sdk.Coin,
) error {
	// create the escrow address for the tokens
	escrowAddress := ibctransfertypes.GetEscrowAddress(sourcePort, sourceChannel)

	// escrow source tokens. It fails if balance insufficient
	if err := k.bankKeeper.SendCoins(
		ctx, escrowAddress, receiver, sdk.NewCoins(tokens),
	); err != nil {
		return err
	}

	return nil
}