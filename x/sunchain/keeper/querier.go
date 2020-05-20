package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

const (
	QueryOrder = "order"

	QueryProduct  = "product"
	QueryProducts = "products"

	QuerySell  = "sell"
	QuerySells = "sells"

	QueryReservation          = "reservation"
	QueryReservations         = "reservations"
	QueryReservationsBySellID = "reservationsBySellID"
	QueryProductsByOwner      = "productsByOwner"
)

// NewQuerier is the module level router for state queries.
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryOrder:
			return queryOrder(ctx, path[1:], req, keeper)
		case QueryProduct:
			return queryProduct(ctx, path[1:], req, keeper)
		case QueryProducts:
			return queryProducts(ctx, req, keeper)
		case QueryProductsByOwner:
			return queryProductsByOwner(ctx, path[1:], req, keeper)
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
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown sunchain query endpoint")
		}
	}
}

// queryOrder is a query function to get order by order ID.
func queryOrder(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) ([]byte, error) {
	if len(path) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "must specify the order id")
	}
	id, err := strconv.ParseInt(path[0], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, fmt.Sprintf("wrong format for requestid %s", err.Error()))
	}
	order, err := keeper.GetOrder(ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	return keeper.cdc.MustMarshalJSON(order), nil
}

// nolint: unparam
func queryProduct(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	key := "Product-" + path[0]

	product, err := keeper.GetProduct(ctx, key)

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProductDoesNotExist, "product %s not found", path[0])
	}

	res := keeper.cdc.MustMarshalJSON(product)

	return res, nil
}

// nolint: unparam
func queryProducts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var productsList types.QueryResProducts

	iterator := keeper.GetProductsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Product-" <= key && key <= "Product-zzzzzzzz" {
			product, err := keeper.GetProduct(ctx, key)
			if err != nil {
				continue
			}
			productsList = append(productsList, product)
		}
	}

	res := keeper.cdc.MustMarshalJSON(productsList)

	return res, nil
}

// nolint: unparam
func querySell(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	key := "Sell-" + path[0]

	sell, err := keeper.GetSell(ctx, key)

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProductDoesNotExist, "sell %s not found", path[0])
	}

	res := keeper.cdc.MustMarshalJSON(sell)

	return res, nil
}

// nolint: unparam
func querySells(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var sellsList types.QueryResSells

	iterator := keeper.GetSellsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Sell-" <= key && key <= "Sell-zzzzzzzz" {
			sell, err := keeper.GetSell(ctx, key)
			if err != nil {
				continue
			}
			sellsList = append(sellsList, sell)
		}
	}

	res := keeper.cdc.MustMarshalJSON(sellsList)

	return res, nil
}

// nolint: unparam
func queryReservation(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	key := "Reservation-" + path[0]

	reservation, err := keeper.GetReservation(ctx, key)

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrReservationDoesNotExist, "reservation %s not found", path[0])
	}

	res := keeper.cdc.MustMarshalJSON(reservation)

	return res, nil
}

// nolint: unparam
func queryReservations(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var reservationsList types.QueryResReservations

	iterator := keeper.GetReservationsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			reservation, err := keeper.GetReservation(ctx, key)
			if err != nil {
				continue
			}
			reservationsList = append(reservationsList, reservation)
		}
	}

	res := keeper.cdc.MustMarshalJSON(reservationsList)

	return res, nil
}

// nolint: unparam
func queryReservationsBySellID(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var sellID = path[0]
	var reservationsList types.QueryResReservations

	iterator := keeper.GetReservationsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			reservation, err := keeper.GetReservation(ctx, key)
			if err != nil {
				continue
			}
			if reservation.SellID == sellID {
				reservationsList = append(reservationsList, reservation)
			}
		}
	}

	res := keeper.cdc.MustMarshalJSON(reservationsList)

	return res, nil
}

// nolint: unparam
func queryProductsByOwner(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	var accAddress = path[0]
	var productsList types.QueryResProducts
	addr, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	iterator := keeper.GetSellsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Product-" <= key && key <= "Product-zzzzzzzz" {
			product, err := keeper.GetProduct(ctx, key)
			if err != nil {
				continue
			}
			if product.Owner.Equals(addr) {
				productsList = append(productsList, product)
			}
		}
	}

	res := keeper.cdc.MustMarshalJSON(productsList)

	return res, nil
}
