package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/trinhtan/cosmos-hackathon/x/nameservice/types"
)

// query endpoints supported by the nameservice Querier
const (
	QueryResolve     = "resolve"
	QueryWhois       = "whois"
	QueryNames       = "names"
	QueryDescription = "description"

	QueryProduct  = "product"
	QueryProducts = "products"

	QuerySell  = "sell"
	QuerySells = "sells"

	QueryReservation          = "reservation"
	QueryReservations         = "reservations"
	QueryReservationsBySellID = "reservationsBySellID"
	QueryProductsByOwner      = "productsByOwner"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		case QueryDescription:
			return queryDescription(ctx, path[1:], req, keeper)
		case QueryProduct:
			return queryProduct(ctx, path[1:], req, keeper)
		case QueryProducts:
			return queryProducts(ctx, req, keeper)
		case QuerySell:
			return querySell(ctx, path[1:], req, keeper)
		case QuerySells:
			return querySells(ctx, req, keeper)
		case QueryReservation:
			return queryReservation(ctx, path[1:], req, keeper)
		case QueryReservations:
			return queryReservations(ctx, req, keeper)
		case QueryReservationsBySellID:
			return queryReservationsBySellID(ctx, path[1:], req, keeper)
		case QueryProductsByOwner:
			return queryProductsByOwner(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value := keeper.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryDescription(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	description := keeper.GetDescription(ctx, path[0])
	if description == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not query description")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResDescription{Description: description})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	whois := keeper.GetWhois(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var namesList types.QueryResNames

	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryProduct(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	key := "Product-" + path[0]

	product := keeper.GetProduct(ctx, key)

	res, err := codec.MarshalJSONIndent(keeper.cdc, product)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryProducts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var productsList types.QueryResProducts

	iterator := keeper.GetProductsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Product-" <= key && key <= "Product-zzzzzzzz" {
			productsList = append(productsList, keeper.GetProduct(ctx, key))
		}
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, productsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func querySell(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	key := "Sell-" + path[0]

	sell := keeper.GetSell(ctx, key)

	res, err := codec.MarshalJSONIndent(keeper.cdc, sell)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func querySells(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var sellsList types.QueryResSells

	iterator := keeper.GetSellsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Sell-" <= key && key <= "Sell-zzzzzzzz" {
			sellsList = append(sellsList, keeper.GetSell(ctx, key))
		}
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, sellsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryReservation(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	key := "Reservation-" + path[0]

	reservation := keeper.GetReservation(ctx, key)

	res, err := codec.MarshalJSONIndent(keeper.cdc, reservation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryReservations(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var reservationsList types.QueryResReservations

	iterator := keeper.GetReservationsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			reservationsList = append(reservationsList, keeper.GetReservation(ctx, key))
		}
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, reservationsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryReservationsBySellID(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var sellID = path[0]
	var reservationsList types.QueryResReservations

	iterator := keeper.GetSellsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			sell := keeper.GetReservation(ctx, key)
			if sell.SellID == sellID {
				reservationsList = append(reservationsList, sell)
			}
		}
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, reservationsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// nolint: unparam
func queryProductsByOwner(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var accAddress = path[0]
	var productsList types.QueryResProducts
	addr, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		// rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	iterator := keeper.GetSellsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Product-" <= key && key <= "Product-zzzzzzzz" {
			product := keeper.GetProduct(ctx, key)
			if product.Owner.Equals(addr) {
				productsList = append(productsList, product)
			}
		}
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, productsList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
