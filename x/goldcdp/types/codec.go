package types

import (
	"github.com/bandprotocol/bandchain/chain/x/oracle"
	"github.com/cosmos/cosmos-sdk/codec"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/types"
)

// ModuleCdc is the codec for the module.
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
	channel.RegisterCodec(ModuleCdc)
	commitmenttypes.RegisterCodec(ModuleCdc)
	oracle.RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetSourceChannel{}, "goldcdp/SetSourceChannel", nil)
	cdc.RegisterConcrete(MsgBuyGold{}, "goldcdp/BuyGold", nil)
	cdc.RegisterConcrete(MsgCreateProduct{}, "goldcdp/CreateProduct", nil)
	cdc.RegisterConcrete(MsgUpdateProduct{}, "goldcdp/UpdateProduct", nil)

	cdc.RegisterConcrete(MsgCreateSell{}, "goldcdp/CreateSell", nil)
	cdc.RegisterConcrete(MsgUpdateSell{}, "goldcdp/UpdateSell", nil)
	cdc.RegisterConcrete(MsgDeleteSell{}, "goldcdp/DeleteSell", nil)

	cdc.RegisterConcrete(MsgCreateReservation{}, "goldcdp/CreateReservation", nil)
	cdc.RegisterConcrete(MsgUpdateReservation{}, "goldcdp/UpdateReservation", nil)
	cdc.RegisterConcrete(MsgDeleteReservation{}, "goldcdp/DeleteReservation", nil)
}
