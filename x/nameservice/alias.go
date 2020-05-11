package nameservice

import (
	"github.com/trinhtan/cosmos-hackathon/x/nameservice/keeper"
	"github.com/trinhtan/cosmos-hackathon/x/nameservice/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier

	NewWhois             = types.NewWhois
	NewMsgBuyName        = types.NewMsgBuyName
	NewMsgSetName        = types.NewMsgSetName
	NewMsgDeleteName     = types.NewMsgDeleteName
	NewMsgSetDescription = types.NewMsgSetDescription

	NewProduct               = types.NewProduct
	NewMsgCreateProduct      = types.NewMsgCreateProduct
	NewMsgChangeProductOwner = types.NewMsgChangeProductOwner
	NewMsgDeleteProduct      = types.NewMsgDeleteProduct
	NewMsgUpdateProduct      = types.NewMsgUpdateProduct

	NewSell               = types.NewSell
	NewMsgSetSell         = types.NewMsgSetSell
	NewMsgSetSellMinPrice = types.NewMsgSetSellMinPrice

	NewReservation            = types.NewReservation
	NewMsgSetReservation      = types.NewMsgSetReservation
	NewMsgSetReservationPrice = types.NewMsgSetReservationPrice

	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper

	Whois               = types.Whois
	MsgSetName          = types.MsgSetName
	MsgBuyName          = types.MsgBuyName
	MsgDeleteName       = types.MsgDeleteName
	MsgSetDescription   = types.MsgSetDescription
	QueryResResolve     = types.QueryResResolve
	QueryResNames       = types.QueryResNames
	QueryResDescription = types.QueryResDescription

	Product               = types.Product
	MsgCreateProduct      = types.MsgCreateProduct
	MsgUpdateProduct      = types.MsgUpdateProduct
	MsgDeleteProduct      = types.MsgDeleteProduct
	MsgChangeProductOwner = types.MsgChangeProductOwner

	Sell               = types.Sell
	MsgSetSell         = types.MsgSetSell
	MsgSetSellMinPrice = types.MsgSetSellMinPrice

	Reservation            = types.Reservation
	MsgSetReservation      = types.MsgSetReservation
	MsgSetReservationPrice = types.MsgSetReservationPrice
)
