package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)
	cdc.RegisterConcrete(MsgSetDescription{}, "nameservice/SetDescription", nil)

	cdc.RegisterConcrete(MsgCreateProduct{}, "nameservice/CreateProduct", nil)
	// cdc.RegisterConcrete(MsgDeleteProduct{}, "nameservice/DeleteProduct", nil)
	cdc.RegisterConcrete(MsgChangeProductOwner{}, "nameservice/ChangeProductOwner", nil)
	cdc.RegisterConcrete(MsgUpdateProduct{}, "nameservice/UpdateProduct", nil)

	cdc.RegisterConcrete(MsgCreateSell{}, "nameservice/CreateSell", nil)
	cdc.RegisterConcrete(MsgUpdateSell{}, "nameservice/UpdateSell", nil)
	cdc.RegisterConcrete(MsgDeleteSell{}, "nameservice/DeleteSell", nil)

	cdc.RegisterConcrete(MsgCreateReservation{}, "nameservice/CreateReservation", nil)
	cdc.RegisterConcrete(MsgUpdateReservation{}, "nameservice/UpdateReservation", nil)
	cdc.RegisterConcrete(MsgDeleteReservation{}, "nameservice/DeleteReservation", nil)
}
