package nameservice

import (
	"fmt"

	"github.com/trinhtan/cosmos-hackathon/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		case MsgSetDescription:
			return handleMsgSetDescription(ctx, keeper, msg)
		case MsgCreateProduct:
			return handleMsgCreateProduct(ctx, keeper, msg)
		case MsgUpdateProduct:
			return handleMsgUpdateProduct(ctx, keeper, msg)
		case MsgDeleteProduct:
			return handleMsgDeleteProduct(ctx, keeper, msg)
		case MsgChangeProductOwner:
			return handleMsgChangeProductOwner(ctx, keeper, msg)
		case MsgSetSell:
			return handleMsgSetSell(ctx, keeper, msg)
		case MsgSetSellMinPrice:
			return handleMsgSetSellMinPrice(ctx, keeper, msg)
		case MsgSetReservation:
			return handleMsgSetReservation(ctx, keeper, msg)
		case MsgSetReservationPrice:
			return handleMsgSetReservationPrice(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil                // return
}

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough") // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return nil, err
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) (*sdk.Result, error) {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.Name)
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteWhois(ctx, msg.Name)
	return &sdk.Result{}, nil
}

// Handle a message to set description
func handleMsgSetDescription(ctx sdk.Context, keeper Keeper, msg MsgSetDescription) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	keeper.SetDescription(ctx, msg.Name, msg.Description) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil                             // return
}

// handleCreateProduct handles a message to set product
func handleMsgCreateProduct(ctx sdk.Context, keeper Keeper, msg MsgCreateProduct) (*sdk.Result, error) {
	if keeper.IsProductPresent(ctx, msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.ProductID)
	}

	var product = Product{
		ProductID:   msg.ProductID,
		Title:       msg.Title,
		Description: msg.Description,
		Owner:       msg.Signer,
	}

	keeper.SetProduct(ctx, msg.ProductID, product)
	return &sdk.Result{}, nil // return
}

// handleMsgUpdateProduct handles a message to set product
func handleMsgUpdateProduct(ctx sdk.Context, keeper Keeper, msg MsgUpdateProduct) (*sdk.Result, error) {
	if !keeper.IsProductPresent(ctx, msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	if !msg.Signer.Equals(keeper.GetProductOwner(ctx, msg.ProductID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	// var product = keeper.GetProduct(ctx, msg.ProductID)

	// if product.Title == msg.Title && product.Description == msg.Description {
	// 	return nil, sdkerrors.Wrap(types.ErrNothingChanges, msg.ProductID)
	// }

	var newInfo = Product{
		ProductID:   msg.ProductID,
		Title:       msg.Title,
		Description: msg.Description,
		Owner:       msg.Signer,
	}

	keeper.SetProduct(ctx, msg.ProductID, newInfo)
	return &sdk.Result{}, nil // return
}

// Handle a message to delete product
func handleMsgDeleteProduct(ctx sdk.Context, keeper Keeper, msg MsgDeleteProduct) (*sdk.Result, error) {
	if !keeper.IsProductPresent(ctx, msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.ProductID)
	}
	if !msg.Signer.Equals(keeper.GetProductOwner(ctx, msg.ProductID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteProduct(ctx, msg.ProductID)
	return &sdk.Result{}, nil
}

// handleMsgChangeProductOwner a message to delete product
func handleMsgChangeProductOwner(ctx sdk.Context, keeper Keeper, msg MsgChangeProductOwner) (*sdk.Result, error) {
	if !keeper.IsProductPresent(ctx, msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}
	if !msg.Signer.Equals(keeper.GetProductOwner(ctx, msg.ProductID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.ChangeProductOwner(ctx, msg.ProductID, msg.NewOwner)
	return &sdk.Result{}, nil
}

// handleMsgSetSell handles a message to set sell
func handleMsgSetSell(ctx sdk.Context, keeper Keeper, msg MsgSetSell) (*sdk.Result, error) {

	if !keeper.IsProductPresent(ctx, msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	if !msg.Signer.Equals(keeper.GetProductOwner(ctx, msg.ProductID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var sell = Sell{
		SellID:    msg.SellID,
		ProductID: msg.ProductID,
		Seller:    msg.Signer,
		MinPrice:  msg.MinPrice,
	}

	keeper.SetSell(ctx, msg.SellID, sell)
	return &sdk.Result{}, nil // return
}

// handleMsgSetSellMinPrice handles a message to set sell minprice
func handleMsgSetSellMinPrice(ctx sdk.Context, keeper Keeper, msg MsgSetSellMinPrice) (*sdk.Result, error) {

	if !keeper.IsSellPresent(ctx, msg.SellID) {
		return nil, sdkerrors.Wrap(types.ErrSellDoesNotExist, msg.SellID)
	}

	if !msg.Signer.Equals(keeper.GetSell(ctx, msg.SellID).Seller) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.SetSellMinPrice(ctx, msg.SellID, msg.MinPrice)
	return &sdk.Result{}, nil // return
}

// handleMsgSetReservation handles a message to set reservation
func handleMsgSetReservation(ctx sdk.Context, keeper Keeper, msg MsgSetReservation) (*sdk.Result, error) {

	if !keeper.IsSellPresent(ctx, msg.SellID) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.SellID)
	}

	if keeper.GetSellMinPrice(ctx, msg.SellID).IsAllGT(msg.Price) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Price not high enough") // If not, throw an error
	}

	var reservation = Reservation{
		ReservationID: msg.ReservationID,
		SellID:        msg.SellID,
		Buyer:         msg.Signer,
		Price:         msg.Price,
	}

	keeper.SetReservation(ctx, msg.ReservationID, reservation)
	return &sdk.Result{}, nil // return
}

// handleMsgSetReservation handles a message to set reservation
func handleMsgSetReservationPrice(ctx sdk.Context, keeper Keeper, msg MsgSetReservationPrice) (*sdk.Result, error) {

	if !keeper.IsReservationPresent(ctx, msg.ReservationID) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	keeper.SetReservationPrice(ctx, msg.ReservationID, msg.Price)
	return &sdk.Result{}, nil // return
}
