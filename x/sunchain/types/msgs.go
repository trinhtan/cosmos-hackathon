package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the goldcdp module
const RouterKey = ModuleName

// MsgSetSoruceChannel is a message for setting source channel to other chain
type MsgSetSourceChannel struct {
	ChainName     string         `json:"chain_name"`
	SourcePort    string         `json:"source_port"`
	SourceChannel string         `json:"source_channel"`
	Signer        sdk.AccAddress `json:"signer"`
}

func NewMsgSetSourceChannel(
	chainName, sourcePort, sourceChannel string,
	signer sdk.AccAddress,
) MsgSetSourceChannel {
	return MsgSetSourceChannel{
		ChainName:     chainName,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Signer:        signer,
	}
}

// Route implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Type() string { return "set_source_channel" }

// ValidateBasic implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) ValidateBasic() error {
	// TODO: Add validate basic
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// GetSignBytes implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgBuyGold is a message for creating order to buy gold
type MsgBuyGold struct {
	Buyer  sdk.AccAddress `json:"buyer"`
	Amount sdk.Coins      `json:"amount"`
}

// NewMsgBuyGold creates a new MsgBuyGold instance.
func NewMsgBuyGold(
	buyer sdk.AccAddress,
	amount sdk.Coins,
) MsgBuyGold {
	return MsgBuyGold{
		Buyer:  buyer,
		Amount: amount,
	}
}

// Route implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) Type() string { return "buy_gold" }

// ValidateBasic implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBuyGold: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgBuyGold: Amount must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// GetSignBytes implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgBuyGold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgCreateProduct defines a SetProduct message
type MsgCreateProduct struct {
	ProductID   string         `json:"productID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Images      string         `json:"images"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgCreateProduct is a constructor function for MsgSetProduct
func NewMsgCreateProduct(productID string, title string, description string, category string, images string, signer sdk.AccAddress) MsgCreateProduct {
	return MsgCreateProduct{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Category:    category,
		Images:      images,
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
	Category    string         `json:"category"`
	Images      string         `json:"images"`
	Signer      sdk.AccAddress `json:"signer"`
}

// NewMsgUpdateProduct is a constructor function for MsgSetProduct
func NewMsgUpdateProduct(productID string, title string, description string, category string, images string, signer sdk.AccAddress) MsgUpdateProduct {
	return MsgUpdateProduct{
		ProductID:   productID,
		Title:       title,
		Description: description,
		Category:    category,
		Images:      images,
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

// MsgSetSell defines a SetSell message
type MsgCreateSell struct {
	SellID    string         `json:"sellID"`
	ProductID string         `json:"productID"`
	Signer    sdk.AccAddress `json:"signer"`
	MinPrice  sdk.Coins      `json:"minPrice"`
}

// NewMsgCreateSell is a constructor function for MsgCreateSell
func NewMsgCreateSell(sellID string, productID string, signer sdk.AccAddress, minPrice sdk.Coins) MsgCreateSell {
	return MsgCreateSell{
		SellID:    sellID,
		ProductID: productID,
		Signer:    signer,
		MinPrice:  minPrice,
	}
}

// Route should return the name of the module
func (msg MsgCreateSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSell) Type() string { return "create_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSell) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.SellID) == 0 || len(msg.ProductID) == 0 || msg.MinPrice.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ProductID and/or SellID and/or MinPrice cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgUpdateSell defines a SetSellMinPrice message
type MsgUpdateSell struct {
	SellID   string         `json:"sellID"`
	Signer   sdk.AccAddress `json:"signer"`
	MinPrice sdk.Coins      `json:"minPrice"`
}

// NewMsgUpdateSell is a constructor function for MsgUpdateSell
func NewMsgUpdateSell(sellID string, signer sdk.AccAddress, minPrice sdk.Coins) MsgUpdateSell {
	return MsgUpdateSell{
		SellID:   sellID,
		Signer:   signer,
		MinPrice: minPrice,
	}
}

// Route should return the name of the module
func (msg MsgUpdateSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateSell) Type() string { return "update_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateSell) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.SellID) == 0 || msg.MinPrice.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SellID and/or MinPrice cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateSell) GetSigners() []sdk.AccAddress {
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
type MsgCreateReservation struct {
	ReservationID string         `json:"reservationID"`
	SellID        string         `json:"sellID"`
	Signer        sdk.AccAddress `json:"buyer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgCreateReservation is a constructor function for MsgCreateReservation
func NewMsgCreateReservation(reservationID string, sellID string, signer sdk.AccAddress, price sdk.Coins) MsgCreateReservation {
	return MsgCreateReservation{
		ReservationID: reservationID,
		SellID:        sellID,
		Signer:        signer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgCreateReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateReservation) Type() string { return "set_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateReservation) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 || len(msg.SellID) == 0 || msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "SellID and/or ReservationID and/or Price cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateReservation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateReservation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgUpdateReservation defines a SetReservationPrice message
type MsgUpdateReservation struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
	Price         sdk.Coins      `json:"price"`
}

// NewMsgUpdateReservation is a constructor function for MsgUpdateReservation
func NewMsgUpdateReservation(reservationID string, signer sdk.AccAddress, price sdk.Coins) MsgUpdateReservation {
	return MsgUpdateReservation{
		ReservationID: reservationID,
		Signer:        signer,
		Price:         price,
	}
}

// Route should return the name of the module
func (msg MsgUpdateReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUpdateReservation) Type() string { return "update_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateReservation) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 || msg.Price.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID and/or Price cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateReservation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateReservation) GetSigners() []sdk.AccAddress {
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

// MsgDecideSell defines a DecideSell message
type MsgDecideSell struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
}

// NewMsgDecideSell is a constructor function for MsgDecideSell
func NewMsgDecideSell(reservationID string, signer sdk.AccAddress) MsgDecideSell {
	return MsgDecideSell{
		ReservationID: reservationID,
		Signer:        signer,
	}
}

// Route should return the name of the module
func (msg MsgDecideSell) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDecideSell) Type() string { return "decide_sell" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDecideSell) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDecideSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDecideSell) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDecideSell defines a DecideSell message
type MsgPayReservation struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
}

// NewMsgPayReservation is a constructor function for MsgPayReservation
func NewMsgPayReservation(reservationID string, signer sdk.AccAddress) MsgPayReservation {
	return MsgPayReservation{
		ReservationID: reservationID,
		Signer:        signer,
	}
}

// Route should return the name of the module
func (msg MsgPayReservation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPayReservation) Type() string { return "pay_reservation" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPayReservation) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPayReservation) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPayReservation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// MsgDecideSell defines a DecideSell message
type MsgPayReservationByAtom struct {
	ReservationID string         `json:"reservationID"`
	Signer        sdk.AccAddress `json:"signer"`
}

// NewMsgPayReservationByAtom is a constructor function for MsgPayReservationByAtom
func NewMsgPayReservationByAtom(reservationID string, signer sdk.AccAddress) MsgPayReservationByAtom {
	return MsgPayReservationByAtom{
		ReservationID: reservationID,
		// Reciver:       reciver,
		Signer: signer,
	}
}

// Route should return the name of the module
func (msg MsgPayReservationByAtom) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPayReservationByAtom) Type() string { return "pay_reservation_by_atom" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPayReservationByAtom) ValidateBasic() error {
	if msg.Signer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}
	if len(msg.ReservationID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "ReservationID cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPayReservationByAtom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPayReservationByAtom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
