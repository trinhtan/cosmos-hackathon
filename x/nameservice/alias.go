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

	NewSell          = types.NewSell
	NewMsgCreateSell = types.NewMsgCreateSell
	NewMsgUpdateSell = types.NewMsgUpdateSell
	NewMsgDeleteSell = types.NewMsgDeleteSell

	NewReservation          = types.NewReservation
	NewMsgCreateReservation = types.NewMsgCreateReservation
	NewMsgUpdateReservation = types.NewMsgUpdateReservation
	NewMsgDeleteReservation = types.NewMsgDeleteReservation

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

	Sell          = types.Sell
	MsgCreateSell = types.MsgCreateSell
	MsgUpdateSell = types.MsgUpdateSell
	MsgDeleteSell  = types.MsgDeleteSell

	Reservation          = types.Reservation
	MsgCreateReservation = types.MsgCreateReservation
	MsgUpdateReservation = types.MsgUpdateReservation
	MsgDeleteReservation = types.MsgDeleteReservation
)
