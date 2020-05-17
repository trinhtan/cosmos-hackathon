package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinNamePrice is Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// MinProductPrice is Initial Starting Price for a product that was never previously owned
var MinProductPrice = sdk.Coins{sdk.NewInt64Coin("producttoken", 1)}

// Product is struct that contains all the metadata of a product
type Product struct {
	ProductID   string         `json:"productID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Images      string         `json:"images"`
	Owner       sdk.AccAddress `json:"owner"`
	Selling     bool           `json:"selling"`
	SellID      string         `json:"sellID"`
}

//NewProduct returns a new product
func NewProduct() Product {
	return Product{}
}

// implement fmt.Stringer
func (product Product) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
	ProductID: %s
	Title: %s
	Description: %s
	Category: %s
	Images: %s
	Owner: %s`, product.ProductID, product.Title, product.Description, product.Category, product.Images, product.Owner))
}

// Sell is a struct contains all the metadata of a sell
type Sell struct {
	SellID    string         `json:"sellID"`
	ProductID string         `json:"productID"`
	Seller    sdk.AccAddress `json:"seller"`
	MinPrice  sdk.Coins      `json:"minPrice"`
}

//NewSell returns a new sell
func NewSell() Sell {
	return Sell{}
}

// implement fmt.Stringer
func (sell Sell) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
	SellID: %s
	ProductID: %s
	Seller: %s
	MinPrice: %s`, sell.SellID, sell.ProductID, sell.Seller, sell.MinPrice))
}

// Reservation is a struct contains all the metadata of a reservation
type Reservation struct {
	ReservationID string         `json:"reservationID"`
	SellID        string         `json:"sellID"`
	Buyer         sdk.AccAddress `json:"buyer"`
	Price         sdk.Coins      `json:"price"`
	Decide        bool           `json:"decide"`
}

//NewReservation returns a new Reservation
func NewReservation() Reservation {
	return Reservation{}
}

// implement fmt.Stringer
func (reservation Reservation) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
	ReservationID: %s
	SellID: %s
	Buyer: %s
	Price: %s`, reservation.ReservationID, reservation.SellID, reservation.Buyer, reservation.Price))
}
