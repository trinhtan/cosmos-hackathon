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

	NewProduct                  = types.NewProduct
	NewMsgSetProduct            = types.NewMsgSetProduct
	NewMsgSetProductTitle       = types.NewMsgSetProductTitle
	NewMsgSetProductDescription = types.NewMsgSetProductDescription
	NewMsgDeleteProduct         = types.NewMsgDeleteProduct

	NewSell       = types.NewSell
	NewMsgSetSell = types.NewMsgSetSell

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

	Product                  = types.Product
	MsgSetProduct            = types.MsgSetProduct
	MsgSetProductTitle       = types.MsgSetProductTitle
	MsgSetProductDescription = types.MsgSetProductDescription
	MsgDeleteProduct         = types.MsgDeleteProduct

	Sell       = types.Sell
	MsgSetSell = types.MsgSetSell
)
