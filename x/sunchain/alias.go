package sunchain

import (
	"github.com/trinhtan/cosmos-hackathon/x/sunchain/keeper"
	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier

	NewProduct          = types.NewProduct
	NewMsgCreateProduct = types.NewMsgCreateProduct
	NewMsgUpdateProduct = types.NewMsgUpdateProduct

	NewMsgCreateSell = types.NewMsgCreateSell
	NewMsgUpdateSell = types.NewMsgUpdateSell
	NewMsgDeleteSell = types.NewMsgDeleteSell
	NewMsgDecideSell = types.NewMsgDecideSell

	NewMsgCreateReservation    = types.NewMsgCreateReservation
	NewMsgUpdateReservation    = types.NewMsgUpdateReservation
	NewMsgDeleteReservation    = types.NewMsgDeleteReservation
	NewMsgPayReservation       = types.NewMsgPayReservation
	NewMsgPayReservationByAtom = types.NewMsgPayReservationByAtom
)

type (
	Keeper              = keeper.Keeper
	MsgBuyGold          = types.MsgBuyGold
	MsgSetSourceChannel = types.MsgSetSourceChannel

	Product          = types.Product
	MsgCreateProduct = types.MsgCreateProduct
	MsgUpdateProduct = types.MsgUpdateProduct

	Sell          = types.Sell
	MsgCreateSell = types.MsgCreateSell
	MsgUpdateSell = types.MsgUpdateSell
	MsgDeleteSell = types.MsgDeleteSell
	MsgDecideSell = types.MsgDecideSell

	Reservation             = types.Reservation
	MsgCreateReservation    = types.MsgCreateReservation
	MsgUpdateReservation    = types.MsgUpdateReservation
	MsgDeleteReservation    = types.MsgDeleteReservation
	MsgPayReservation       = types.MsgPayReservation
	MsgPayReservationByAtom = types.MsgPayReservationByAtom
)
