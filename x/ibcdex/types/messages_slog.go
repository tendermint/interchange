package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSlog{}

func NewMsgCreateSlog(creator string, log string) *MsgCreateSlog {
	return &MsgCreateSlog{
		Creator: creator,
		Log:     log,
	}
}

func (msg *MsgCreateSlog) Route() string {
	return RouterKey
}

func (msg *MsgCreateSlog) Type() string {
	return "CreateSlog"
}

func (msg *MsgCreateSlog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSlog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSlog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSlog{}

func NewMsgUpdateSlog(creator string, id string, log string) *MsgUpdateSlog {
	return &MsgUpdateSlog{
		Id:      id,
		Creator: creator,
		Log:     log,
	}
}

func (msg *MsgUpdateSlog) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSlog) Type() string {
	return "UpdateSlog"
}

func (msg *MsgUpdateSlog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSlog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSlog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateSlog{}

func NewMsgDeleteSlog(creator string, id string) *MsgDeleteSlog {
	return &MsgDeleteSlog{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSlog) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSlog) Type() string {
	return "DeleteSlog"
}

func (msg *MsgDeleteSlog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSlog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSlog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
