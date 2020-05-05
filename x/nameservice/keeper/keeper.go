package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trinhtan/cosmos-hackathon/x/nameservice/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper types.BankKeeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, coinKeeper types.BankKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		CoinKeeper: coinKeeper,
	}
}

// GetWhois gets the entire Whois metadata struct for a name
func (k Keeper) GetWhois(ctx sdk.Context, name string) types.Whois {
	store := ctx.KVStore(k.storeKey)

	if !k.IsNamePresent(ctx, name) {
		return types.NewWhois()
	}

	bz := store.Get([]byte(name))

	var whois types.Whois

	k.cdc.MustUnmarshalBinaryBare(bz, &whois)

	return whois
}

// SetWhois sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}

// DeleteWhois deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteWhois(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Value
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

// GetDescription - gets the current description of a name
func (k Keeper) GetDescription(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Description
}

// SetDescription sets the current description of a name
func (k Keeper) SetDescription(ctx sdk.Context, name string, description string) {
	whois := k.GetWhois(ctx, name)
	whois.Description = description
	k.SetWhois(ctx, name, whois)
}

// GetNamesIterator gets an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// IsNamePresent checks if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// GetProduct gets the entire Product metadata struct for a name
func (k Keeper) GetProduct(ctx sdk.Context, productID string) types.Product {
	store := ctx.KVStore(k.storeKey)

	if !k.IsProductPresent(ctx, productID) {
		return types.NewProduct()
	}

	bz := store.Get([]byte(productID))

	var product types.Product

	k.cdc.MustUnmarshalBinaryBare(bz, &product)

	return product
}

// SetProduct sets the entire Product metadata struct for a product
func (k Keeper) SetProduct(ctx sdk.Context, productID string, product types.Product) {
	if product.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(productID), k.cdc.MustMarshalBinaryBare(product))
}

// DeleteProduct deletes the entire Proudct metadata struct for a product
func (k Keeper) DeleteProduct(ctx sdk.Context, productID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(productID))
}

// GetProductTitle gets product title
func (k Keeper) GetProductTitle(ctx sdk.Context, productID string) string {
	return k.GetProduct(ctx, productID).Title
}

// SetProductTitle sets product title
func (k Keeper) SetProductTitle(ctx sdk.Context, productID string, title string) {
	product := k.GetProduct(ctx, productID)
	product.Title = title
	k.SetProduct(ctx, productID, product)
}

// GetProductOwner gets product owner
func (k Keeper) GetProductOwner(ctx sdk.Context, productID string) sdk.AccAddress {
	return k.GetProduct(ctx, productID).Owner
}

// SetProductOwner sets product owner
func (k Keeper) SetProductOwner(ctx sdk.Context, productID string, owner sdk.AccAddress) {
	product := k.GetProduct(ctx, productID)
	product.Owner = owner
	k.SetProduct(ctx, productID, product)
}

// GetProductDescription gets product description
func (k Keeper) GetProductDescription(ctx sdk.Context, productID string) string {
	return k.GetProduct(ctx, productID).Description
}

// SetProductDescription sets product description
func (k Keeper) SetProductDescription(ctx sdk.Context, productID string, description string) {
	product := k.GetProduct(ctx, productID)
	product.Description = description
	k.SetProduct(ctx, productID, product)
}

// GetProductsIterator gets an iterator over all product in which the keys are the productID and the values are the product
func (k Keeper) GetProductsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// IsProductPresent checks if the product is present in the store or not
func (k Keeper) IsProductPresent(ctx sdk.Context, productID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(productID))
}

// GetSell gets the entire Sell metadata struct for a name
func (k Keeper) GetSell(ctx sdk.Context, sellID string) types.Sell {
	store := ctx.KVStore(k.storeKey)

	if !k.IsSellPresent(ctx, sellID) {
		return types.NewSell()
	}

	bz := store.Get([]byte(sellID))

	var sell types.Sell

	k.cdc.MustUnmarshalBinaryBare(bz, &sell)

	return sell
}

// SetSell sets the entire sell metadata struct for a sell
func (k Keeper) SetSell(ctx sdk.Context, sellID string, sell types.Sell) {
	if sell.Seller.Empty() || len(sell.ProductID) == 0 || sell.MinPrice.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(sellID), k.cdc.MustMarshalBinaryBare(sell))
}

// GetSellProductID gets productID of sell
func (k Keeper) GetSellProductID(ctx sdk.Context, sellID string) string {
	return k.GetSell(ctx, sellID).ProductID
}

// SetSellProductID sets productID of sell
func (k Keeper) SetSellProductID(ctx sdk.Context, sellID string, productID string) {
	sell := k.GetSell(ctx, sellID)
	sell.ProductID = productID
	k.SetSell(ctx, sellID, sell)
}

// GetSellSeller gets seller of sell
func (k Keeper) GetSellSeller(ctx sdk.Context, sellID string) sdk.AccAddress {
	return k.GetSell(ctx, sellID).Seller
}

// SetSellSeller sets seller of sell
func (k Keeper) SetSellSeller(ctx sdk.Context, sellID string, seller sdk.AccAddress) {
	sell := k.GetSell(ctx, sellID)
	sell.Seller = seller
	k.SetSell(ctx, sellID, sell)
}

// GetSellMinPrice gets MinPrice of sell
func (k Keeper) GetSellMinPrice(ctx sdk.Context, sellID string) sdk.Coins {
	return k.GetSell(ctx, sellID).MinPrice
}

// SetSellMinPrice sets MinPrice of sell
func (k Keeper) SetSellMinPrice(ctx sdk.Context, sellID string, minPrice sdk.Coins) {
	sell := k.GetSell(ctx, sellID)
	sell.MinPrice = minPrice
	k.SetSell(ctx, sellID, sell)
}

// GetSellsIterator gets an iterator over all sell in which the keys are the sellID and the values are the sell
func (k Keeper) GetSellsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// IsSellPresent checks if the sell is present in the store or not
func (k Keeper) IsSellPresent(ctx sdk.Context, sellID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(sellID))
}

// GetReservation gets the entire reservation metadata struct for a reservation
func (k Keeper) GetReservation(ctx sdk.Context, reservationID string) types.Reservation {
	store := ctx.KVStore(k.storeKey)

	if !k.IsReservationPresent(ctx, reservationID) {
		return types.NewReservation()
	}

	bz := store.Get([]byte(reservationID))

	var reservation types.Reservation

	k.cdc.MustUnmarshalBinaryBare(bz, &reservation)

	return reservation
}

// SetReservation sets the entire sell metadata struct for a reservation
func (k Keeper) SetReservation(ctx sdk.Context, reservationID string, reservation types.Reservation) {
	if reservation.Buyer.Empty() || len(reservation.ProductID) == 0 || reservation.Price.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(reservationID), k.cdc.MustMarshalBinaryBare(reservation))
}

// GetReservationProductID gets productID of reservation
func (k Keeper) GetReservationProductID(ctx sdk.Context, reservationID string) string {
	return k.GetReservation(ctx, reservationID).ProductID
}

// SetReservationProductID sets productID of reservation
func (k Keeper) SetReservationProductID(ctx sdk.Context, reservationID string, productID string) {
	reservation := k.GetReservation(ctx, reservationID)
	reservation.ProductID = productID
	k.SetReservation(ctx, reservationID, reservation)
}

// GetReservationBuyer gets Buyer of reservation
func (k Keeper) GetReservationBuyer(ctx sdk.Context, reservationID string) sdk.AccAddress {
	return k.GetReservation(ctx, reservationID).Buyer
}

// SetReservationBuyer sets Buyer of reservation
func (k Keeper) SetReservationBuyer(ctx sdk.Context, reservationID string, buyer sdk.AccAddress) {
	reservation := k.GetReservation(ctx, reservationID)
	reservation.Buyer = buyer
	k.SetReservation(ctx, reservationID, reservation)
}

// GetReservationPrice gets price of reservation
func (k Keeper) GetReservationPrice(ctx sdk.Context, reservationID string) sdk.Coins {
	return k.GetReservation(ctx, reservationID).Price
}

// SetReservationPrice sets Price of reservation
func (k Keeper) SetReservationPrice(ctx sdk.Context, reservationID string, price sdk.Coins) {
	reservation := k.GetReservation(ctx, reservationID)
	reservation.Price = price
	k.SetReservation(ctx, reservationID, reservation)
}

// GetReservationsIterator gets an iterator over all reservations in which the keys are the reservationID and the values are the reservation
func (k Keeper) GetReservationsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// IsReservationPresent checks if the reservation is present in the store or not
func (k Keeper) IsReservationPresent(ctx sdk.Context, reservationID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(reservationID))
}
