package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendCancelOrder{}, "ibcdex/SendCancelOrder", nil)

	cdc.RegisterConcrete(&MsgSendTargetBuyOrder{}, "ibcdex/SendTargetBuyOrder", nil)

	cdc.RegisterConcrete(&MsgSendSourceBuyOrder{}, "ibcdex/SendSourceBuyOrder", nil)

	cdc.RegisterConcrete(&MsgSendTargetSellOrder{}, "ibcdex/SendTargetSellOrder", nil)

	cdc.RegisterConcrete(&MsgSendSourceSellOrder{}, "ibcdex/SendSourceSellOrder", nil)

	cdc.RegisterConcrete(&MsgSendCreatePair{}, "ibcdex/SendCreatePair", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCancelOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendTargetBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendSourceBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendTargetSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendSourceSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCreatePair{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
