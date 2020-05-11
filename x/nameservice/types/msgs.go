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
type MsgSetProduct struct {
	ProductID   string         `json:"productID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"`
}

// NewMsgSetProduct is a constructor function for MsgSetProduct
func NewMsgSetProduct(productID string, title string, description string, owner sdk.AccAddress) MsgSetProduct {
	return MsgSetProduct{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Owner:       owner,
	}
}

// Route should return the name of the module
func (msg MsgSetProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetProduct) Type() string { return "set_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetProduct) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Title) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Title cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetProduct) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetProductTitle defines a SetProductTitle message
type MsgSetProductTitle struct {
	ProductID string         `json:"productID"`
	Title     string         `json:"title"`
	Owner     sdk.AccAddress `json:"owner"`
}

// NewMsgSetProductTitle is a constructor function for MsgSetProductTitle
func NewMsgSetProductTitle(productID string, title string, owner sdk.AccAddress) MsgSetProductTitle {
	return MsgSetProductTitle{
		ProductID: productID,
		Title:     title,
		Owner:     owner,
	}
}

// Route should return the name of the module
func (msg MsgSetProductTitle) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetProductTitle) Type() string { return "set_product_title" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetProductTitle) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Title) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Title cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetProductTitle) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetProductTitle) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetProductDescription defines a SetProductDescription message
type MsgSetProductDescription struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"`
}

// NewMsgSetProductDescription is a constructor function for MsgSetProductDescription
func NewMsgSetProductDescription(productID string, description string, owner sdk.AccAddress) MsgSetProductDescription {
	return MsgSetProductDescription{
		ProductID:   productID,
		Description: description,
		Owner:       owner,
	}
}

// Route should return the name of the module
func (msg MsgSetProductDescription) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetProductDescription) Type() string { return "set_product_description" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetProductDescription) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.ProductID) == 0 || len(msg.Description) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or Description cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetProductDescription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetProductDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgChangeProductOwner defines the ChangeProductOwner message
type MsgChangeProductOwner struct {
	ProductID    string         `json:"productID"`
	Bid          sdk.Coins      `json:"bid"`
	CurrentOwner sdk.AccAddress `json:"currentOwner"`
	NewOwner     sdk.AccAddress `json:"newOnwer"`
}

// NewMsgChangeProductOwner is the constructor function for MsgBuyName
func NewMsgChangeProductOwner(productID string, bid sdk.Coins, currentOwner sdk.AccAddress, newOwner sdk.AccAddress) MsgChangeProductOwner {
	return MsgChangeProductOwner{
		ProductID:    productID,
		Bid:          bid,
		CurrentOwner: currentOwner,
		NewOwner:     newOwner,
	}
}

// Route should return the name of the module
func (msg MsgChangeProductOwner) Route() string { return RouterKey }

// Type should return the action
func (msg MsgChangeProductOwner) Type() string { return "change_product_owner" }

// ValidateBasic runs stateless checks on the message
func (msg MsgChangeProductOwner) ValidateBasic() error {

	if msg.CurrentOwner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.CurrentOwner.String())
	}

	if msg.NewOwner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.NewOwner.String())
	}

	if len(msg.ProductID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID cannot be empty")
	}

	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgChangeProductOwner) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgChangeProductOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.CurrentOwner}
}

// MsgDeleteProduct defines a DeleteProduct message
type MsgDeleteProduct struct {
	ProductID string         `json:"productID"`
	Owner     sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteProduct is a constructor function for NewMsgDeleteProduct
func NewMsgDeleteProduct(productID string, owner sdk.AccAddress) MsgDeleteProduct {
	return MsgDeleteProduct{
		ProductID: productID,
		Owner:     owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProduct) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProduct) Type() string { return "delete_product" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProduct) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
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
	return []sdk.AccAddress{msg.Owner}
}

// MsgSetSell defines a SetSell message
type MsgSetSell struct {
	SellID    string         `json:"sellID"`
	ProductID string         `json:"productID"`
	Seller    sdk.AccAddress `json:"seller"`
	MinPrice  sdk.Coins      `json:"minPrice"`
}

// NewMsgSetSell is a constructor function for MsgSetSell
func NewMsgSetSell(sellID string, productID string, seller sdk.AccAddress, minPrice sdk.Coins) MsgSetSell {
	return MsgSetSell{
		SellID:    sellID,
		ProductID: productID,
		Seller:    seller,
		MinPrice:  minPrice,
	}
}

// Route should return the name of the module
func (msg MsgSetSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetSell) Type() string { return "set_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetSell) ValidateBasic() error {
	if msg.Seller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Seller.String())
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
	return []sdk.AccAddress{msg.Seller}
}

// MsgSetSellMinPrice defines a SetSellMinPrice message
type MsgSetSellMinPrice struct {
	SellID   string         `json:"sellID"`
	Seller   sdk.AccAddress `json:"seller"`
	MinPrice sdk.Coins      `json:"minPrice"`
}

// NewMsgSetSellMinPrice is a constructor function for MsgSetSellMinPrice
func NewMsgSetSellMinPrice(sellID string, seller sdk.AccAddress, minPrice sdk.Coins) MsgSetSellMinPrice {
	return MsgSetSellMinPrice{
		SellID:   sellID,
		Seller:   seller,
		MinPrice: minPrice,
	}
}

// Route should return the name of the module
func (msg MsgSetSellMinPrice) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetSellMinPrice) Type() string { return "set_sell_minPrice" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetSellMinPrice) ValidateBasic() error {
	if msg.Seller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Seller.String())
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
	return []sdk.AccAddress{msg.Seller}
}

// MsgDeleteSell defines a DeleteSell message
type MsgDeleteSell struct {
	SellID string         `json:"sellID"`
	Seller sdk.AccAddress `json:"seller"`
}

// NewMsgDeleteSell is a constructor function for MsgDeleteSell
func NewMsgDeleteSell(sellID string, seller sdk.AccAddress) MsgDeleteSell {
	return MsgDeleteSell{
		SellID: sellID,
		Seller: seller,
	}
}

// Route should return the name of the module
func (msg MsgDeleteSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteSell) Type() string { return "delete_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteSell) ValidateBasic() error {
	if msg.Seller.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Seller.String())
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
	return []sdk.AccAddress{msg.Seller}
}

// MsgSetReservation defines a SetReservation message
type MsgSetReservation struct {
	ReservationID string         `json:"reservationID"`
	SellID        string         `json:"sellID"`
	Buyer         sdk.AccAddress `json:"buyer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgSetReservation is a constructor function for MsgSetReservation
func NewMsgSetReservation(reservationID string, sellID string, buyer sdk.AccAddress, price sdk.Coins) MsgSetReservation {
	return MsgSetReservation{
		ReservationID: reservationID,
		SellID:        sellID,
		Buyer:         buyer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgSetReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetReservation) Type() string { return "set_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetReservation) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
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
	return []sdk.AccAddress{msg.Buyer}
}

// MsgSetReservationPrice defines a SetReservationPrice message
type MsgSetReservationPrice struct {
	ReservationID string         `json:"reservationID"`
	Buyer         sdk.AccAddress `json:"buyer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgSetReservationPrice is a constructor function for MsgSetReservation
func NewMsgSetReservationPrice(reservationID string, buyer sdk.AccAddress, price sdk.Coins) MsgSetReservationPrice {
	return MsgSetReservationPrice{
		ReservationID: reservationID,
		Buyer:         buyer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgSetReservationPrice) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetReservationPrice) Type() string { return "set_reservation_price" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetReservationPrice) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
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
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteReservation defines a DeleteReservation message
type MsgDeleteReservation struct {
	ReservationID string         `json:"reservationID"`
	Buyer         sdk.AccAddress `json:"buyer"`
}

// NewMsgDeleteReservation is a constructor function for MsgDeleteReservation
func NewMsgDeleteReservation(reservationId string, buyer sdk.AccAddress) MsgDeleteReservation {
	return MsgDeleteReservation{
		ReservationID: reservationId,
		Buyer:         buyer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteReservation) Type() string { return "delete_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteReservation) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
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
	return []sdk.AccAddress{msg.Buyer}
}
