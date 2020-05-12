package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteName defines a DeleteName message
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteName is a constructor function for MsgDeleteName
func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteName) Type() string { return "delete_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetDescription defines a SetDescription message
type MsgSetDescription struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"`
}

// NewMsgSetDescription is a constructor function for MsgSetDescription
func NewMsgSetDescription(name string, description string, owner sdk.AccAddress) MsgSetDescription {
	return MsgSetDescription{
		Name:        name,
		Description: description,
		Owner:       owner,
	}
}

// Route should return the name of the module
func (msg MsgSetDescription) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetDescription) Type() string { return "set_description" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetDescription) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetDescription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetProduct defines a SetProduct message
type MsgCreateProduct struct {
	ProductID   string         `json:"productID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgSetProduct is a constructor function for MsgSetProduct
func NewMsgCreateProduct(productID string, title string, description string, signer sdk.AccAddress) MsgCreateProduct {
	return MsgCreateProduct{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgCreateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateProduct) Type() string { return "create_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Title) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Title and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgUpdateProduct defines a UpdateProduct message
type MsgUpdateProduct struct {
	ProductID   string         `json:"productID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgUpdateProduct is a constructor function for MsgSetProduct
func NewMsgUpdateProduct(productID string, title string, description string, signer sdk.AccAddress) MsgUpdateProduct {
	return MsgUpdateProduct{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgUpdateProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateProduct) Type() string { return "update_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Title) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Title and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgChangeProductOwner defines the ChangeProductOwner message
type MsgChangeProductOwner struct {
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"singer"`
	NewOwner  sdk.AccAddress `json:"newOnwer"`
}

// NewMsgChangeProductOwner is the constructor function for MsgBuyName
func NewMsgChangeProductOwner(productID string, signer sdk.AccAddress, newOwner sdk.AccAddress) MsgChangeProductOwner {
	return MsgChangeProductOwner{
		ProductID: productID,
		Signer:    signer,
		NewOwner:  newOwner,
	}
}

// Route should return the name of the module
func (msg MsgChangeProductOwner) Route() string { return RouterKey }

// Type should return the action
func (msg MsgChangeProductOwner) Type() string { return "change_product_owner" }

// ValidateBasic runs stateless checks on the message
func (msg MsgChangeProductOwner) ValidateBasic() error {

	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}

	if msg.NewOwner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.NewOwner.String())
	}

	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgChangeProductOwner) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgChangeProductOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDeleteProduct defines a DeleteProduct message
type MsgDeleteProduct struct {
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
}

// NewMsgDeleteProduct is a constructor function for NewMsgDeleteProduct
func NewMsgDeleteProduct(productID string, signer sdk.AccAddress) MsgDeleteProduct {
	return MsgDeleteProduct{
		ProductID: productID,
		Signer:    signer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProduct) Type() string { return "delete_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProduct) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgSetSell defines a SetSell message
type MsgSetSell struct {
	SellID    string         `json:"sellID"`
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
	MinPrice  sdk.Coins      `json:"minPrice"`
}

// NewMsgSetSell is a constructor function for MsgSetSell
func NewMsgSetSell(sellID string, productID string, signer sdk.AccAddress, minPrice sdk.Coins) MsgSetSell {
	return MsgSetSell{
		SellID:    sellID,
		ProductID: productID,
		Signer:    signer,
		MinPrice:  minPrice,
	}
}

// Route should return the name of the module
func (msg MsgSetSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetSell) Type() string { return "set_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetSell) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.SellID) == 0 || len(msg.ProductID) == 0 || msg.MinPrice.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or SellID and/or MinPrice cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgSetSellMinPrice defines a SetSellMinPrice message
type MsgSetSellMinPrice struct {
	SellID   string         `json:"sellID"`
	Signer   sdk.AccAddress `json:"signer"`
	MinPrice sdk.Coins      `json:"minPrice"`
}

// NewMsgSetSellMinPrice is a constructor function for MsgSetSellMinPrice
func NewMsgSetSellMinPrice(sellID string, signer sdk.AccAddress, minPrice sdk.Coins) MsgSetSellMinPrice {
	return MsgSetSellMinPrice{
		SellID:   sellID,
		Signer:   signer,
		MinPrice: minPrice,
	}
}

// Route should return the name of the module
func (msg MsgSetSellMinPrice) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetSellMinPrice) Type() string { return "set_sell_minPrice" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetSellMinPrice) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.SellID) == 0 || msg.MinPrice.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SellID and/or MinPrice cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetSellMinPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetSellMinPrice) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDeleteSell defines a DeleteSell message
type MsgDeleteSell struct {
	SellID string         `json:"sellID"`
	Signer sdk.AccAddress `json:"signer"`
}

// NewMsgDeleteSell is a constructor function for MsgDeleteSell
func NewMsgDeleteSell(sellID string, signer sdk.AccAddress) MsgDeleteSell {
	return MsgDeleteSell{
		SellID: sellID,
		Signer: signer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteSell) Type() string { return "delete_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteSell) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.SellID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SellID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgSetReservation defines a SetReservation message
type MsgSetReservation struct {
	ReservationID string         `json:"reservationID"`
	SellID        string         `json:"sellID"`
	Signer        sdk.AccAddress `json:"buyer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgSetReservation is a constructor function for MsgSetReservation
func NewMsgSetReservation(reservationID string, sellID string, signer sdk.AccAddress, price sdk.Coins) MsgSetReservation {
	return MsgSetReservation{
		ReservationID: reservationID,
		SellID:        sellID,
		Signer:        signer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgSetReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetReservation) Type() string { return "set_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetReservation) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 || len(msg.SellID) == 0 || msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SellID and/or ReservationID and/or Price cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetReservation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetReservation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgSetReservationPrice defines a SetReservationPrice message
type MsgSetReservationPrice struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgSetReservationPrice is a constructor function for MsgSetReservation
func NewMsgSetReservationPrice(reservationID string, signer sdk.AccAddress, price sdk.Coins) MsgSetReservationPrice {
	return MsgSetReservationPrice{
		ReservationID: reservationID,
		Signer:        signer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgSetReservationPrice) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetReservationPrice) Type() string { return "set_reservation_price" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetReservationPrice) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 || msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID and/or Price cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetReservationPrice) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetReservationPrice) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDeleteReservation defines a DeleteReservation message
type MsgDeleteReservation struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
}

// NewMsgDeleteReservation is a constructor function for MsgDeleteReservation
func NewMsgDeleteReservation(reservationId string, signer sdk.AccAddress) MsgDeleteReservation {
	return MsgDeleteReservation{
		ReservationID: reservationId,
		Signer:        signer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteReservation) Type() string { return "delete_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteReservation) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteReservation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteReservation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
