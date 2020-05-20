package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	BankKeeper    types.BankKeeper
	ChannelKeeper types.ChannelKeeper
}

// NewKeeper creates a new band consumer Keeper instance.
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		BankKeeper:    bankKeeper,
		ChannelKeeper: channelKeeper,
	}
}

// GetOrderCount returns the current number of all orders ever exist.
func (k Keeper) GetOrderCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.OrdersCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNextOrderCount increments and returns the current number of orders.
// If the global order count is not set, it initializes it with value 0.
func (k Keeper) GetNextOrderCount(ctx sdk.Context) uint64 {
	orderCount := k.GetOrderCount(ctx)
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(orderCount + 1)
	store.Set(types.OrdersCountStoreKey, bz)
	return orderCount + 1
}

// GetProduct gets the entire Product metadata struct for a name
func (k Keeper) GetProduct(ctx sdk.Context, key string) (types.Product, error) {
	store := ctx.KVStore(k.storeKey)

	if !k.IsProductPresent(ctx, key) {
		return types.NewProduct(), sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Key not found", key)
	}

	bz := store.Get([]byte(key))

	var product types.Product

	k.cdc.MustUnmarshalBinaryBare(bz, &product)

	return product, nil
}

// CreateProduct sets the entire Product metadata struct for a product
func (k Keeper) SetProduct(ctx sdk.Context, key string, product types.Product) {

	if product.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(product))
}

// GetProductTitle gets product title
func (k Keeper) GetProductTitle(ctx sdk.Context, key string) (string, error) {
	product, err := k.GetProduct(ctx, key)
	if err != nil {
		return "Product does not exist!", err
	}
	return product.Title, nil
}

// GetProductOwner gets product owner
func (k Keeper) GetProductOwner(ctx sdk.Context, key string) (sdk.AccAddress, error) {
	product, err := k.GetProduct(ctx, key)
	if err != nil {
		return product.Owner, err
	}

	return product.Owner, nil
}

// GetProductDescription gets product description
func (k Keeper) GetProductDescription(ctx sdk.Context, key string) (string, error) {
	product, err := k.GetProduct(ctx, key)
	if err != nil {
		return "Product does not exist!", err
	}
	return product.Description, nil
}

// GetProductCategory gets product description
func (k Keeper) GetProductCategory(ctx sdk.Context, key string) (string, error) {
	product, err := k.GetProduct(ctx, key)
	if err != nil {
		return "Product does not exist!", err
	}
	return product.Category, nil
}

// GetProductImages gets product description
func (k Keeper) GetProductImages(ctx sdk.Context, key string) (string, error) {
	product, err := k.GetProduct(ctx, key)
	if err != nil {
		return "Product does not exist!", err
	}

	return product.Images, nil
}

// GetProductsIterator gets an iterator over all product in which the keys are the productID and the values are the product
func (k Keeper) GetProductsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// IsProductPresent checks if the product is present in the store or not
func (k Keeper) IsProductPresent(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key))
}

// GetSell gets the entire Sell metadata struct for a name
func (k Keeper) GetSell(ctx sdk.Context, key string) (types.Sell, error) {
	store := ctx.KVStore(k.storeKey)

	if !k.IsSellPresent(ctx, key) {
		return types.NewSell(), sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Key not found", key)
	}

	bz := store.Get([]byte(key))

	var sell types.Sell

	k.cdc.MustUnmarshalBinaryBare(bz, &sell)

	return sell, nil
}

// SetSell sets the entire sell metadata struct for a sell
func (k Keeper) SetSell(ctx sdk.Context, key string, sell types.Sell) {
	if sell.Seller.Empty() || len(sell.ProductID) == 0 || sell.MinPrice.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(sell))
}

// IsSellPresent checks if the sell is present in the store or not
func (k Keeper) IsSellPresent(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key))
}

// GetSellsIterator gets an iterator over all sell in which the keys are the sellID and the values are the sell
func (k Keeper) GetSellsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// DeleteSell deletes the entire Sell metadata struct for a sell
func (k Keeper) DeleteSell(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}

// GetReservation gets the entire reservation metadata struct for a reservation
func (k Keeper) GetReservation(ctx sdk.Context, key string) (types.Reservation, error) {
	store := ctx.KVStore(k.storeKey)

	if !k.IsReservationPresent(ctx, key) {
		return types.NewReservation(), sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Key not found", key)
	}

	bz := store.Get([]byte(key))

	var reservation types.Reservation

	k.cdc.MustUnmarshalBinaryBare(bz, &reservation)

	return reservation, nil
}

// SetReservation sets the entire sell metadata struct for a reservation
func (k Keeper) SetReservation(ctx sdk.Context, key string, reservation types.Reservation) {
	if reservation.Buyer.Empty() || len(reservation.SellID) == 0 || reservation.Price.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(reservation))
}

// IsReservationPresent checks if the reservation is present in the store or not
func (k Keeper) IsReservationPresent(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(key))
}

// GetReservationsIterator gets an iterator over all reservations in which the keys are the reservationID and the values are the reservation
func (k Keeper) GetReservationsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// DeleteSell deletes the entire Sell metadata struct for a sell
func (k Keeper) DeleteReservation(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(key))
}
