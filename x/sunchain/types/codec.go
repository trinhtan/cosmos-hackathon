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
	cdc.RegisterConcrete(MsgSetSourceChannel{}, "sunchain/SetSourceChannel", nil)
	cdc.RegisterConcrete(MsgBuyGold{}, "sunchain/BuyGold", nil)
	cdc.RegisterConcrete(MsgCreateProduct{}, "sunchain/CreateProduct", nil)
	cdc.RegisterConcrete(MsgUpdateProduct{}, "sunchain/UpdateProduct", nil)

	cdc.RegisterConcrete(MsgCreateSell{}, "sunchain/CreateSell", nil)
	cdc.RegisterConcrete(MsgUpdateSell{}, "sunchain/UpdateSell", nil)
	cdc.RegisterConcrete(MsgDeleteSell{}, "sunchain/DeleteSell", nil)
	cdc.RegisterConcrete(MsgDecideSell{}, "sunchain/DecideSell", nil)

	cdc.RegisterConcrete(MsgCreateReservation{}, "sunchain/CreateReservation", nil)
	cdc.RegisterConcrete(MsgUpdateReservation{}, "sunchain/UpdateReservation", nil)
	cdc.RegisterConcrete(MsgDeleteReservation{}, "sunchain/DeleteReservation", nil)
	cdc.RegisterConcrete(MsgPayReservation{}, "sunchain/PayReservation", nil)
	cdc.RegisterConcrete(MsgPayReservationByAtom{}, "sunchain/PayReservationByAtom", nil)

}
