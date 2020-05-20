package sunchain

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"

	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

// NewHandler creates the msg handler of this module, as required by Cosmos-SDK standard.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgBuyGold:
			return handleBuyGold(ctx, msg, keeper)
		case MsgCreateProduct:
			return handleMsgCreateProduct(ctx, keeper, msg)
		case MsgUpdateProduct:
			return handleMsgUpdateProduct(ctx, keeper, msg)
		case MsgCreateSell:
			return handleMsgCreateSell(ctx, keeper, msg)
		case MsgUpdateSell:
			return handleMsgUpdateSell(ctx, keeper, msg)
		case MsgDeleteSell:
			return handleMsgDeleteSell(ctx, keeper, msg)
		case MsgDecideSell:
			return handleMsgDecideSell(ctx, keeper, msg)
		case MsgCreateReservation:
			return handleMsgCreateReservation(ctx, keeper, msg)
		case MsgUpdateReservation:
			return handleMsgUpdateReservation(ctx, keeper, msg)
		case MsgDeleteReservation:
			return handleMsgDeleteReservation(ctx, keeper, msg)
		case MsgPayReservation:
			return handleMsgPayReservation(ctx, keeper, msg)
		case MsgPayReservationByAtom:
			return handleMsgPayReservationByAtom(ctx, keeper, msg)
		case MsgSetSourceChannel:
			return handleSetSourceChannel(ctx, msg, keeper)
		case channeltypes.MsgPacket:
			var responsePacket oracle.OracleResponsePacketData
			if err := types.ModuleCdc.UnmarshalJSON(msg.GetData(), &responsePacket); err == nil {
				return handleOracleRespondPacketData(ctx, responsePacket, keeper)
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal oracle packet data")
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleBuyGold(ctx sdk.Context, msg MsgBuyGold, keeper Keeper) (*sdk.Result, error) {
	orderID, err := keeper.AddOrder(ctx, msg.Buyer, msg.Amount)
	if err != nil {
		return nil, err
	}

	bandChainID := "bandchain"
	port := "sunchain"
	oracleScriptID := oracle.OracleScriptID(3)
	calldata := make([]byte, 8)
	binary.LittleEndian.PutUint64(calldata, 1000000)
	askCount := int64(1)
	minCount := int64(1)

	channelID, err := keeper.GetChannel(ctx, bandChainID, port)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"not found channel to bandchain",
		)
	}
	sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, port, channelID)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port goldcdp",
			channelID,
		)
	}
	destinationPort := sourceChannelEnd.Counterparty.PortID
	destinationChannel := sourceChannelEnd.Counterparty.ChannelID
	sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
		ctx, port, channelID,
	)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			channelID,
		)
	}
	packet := oracle.NewOracleRequestPacketData(
		fmt.Sprintf("Order:%d", orderID), oracleScriptID, hex.EncodeToString(calldata),
		askCount, minCount,
	)
	err = keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
		sequence, port, channelID, destinationPort, destinationChannel,
		1000000000, // Arbitrarily high timeout for now
	))
	if err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleSetSourceChannel(ctx sdk.Context, msg MsgSetSourceChannel, keeper Keeper) (*sdk.Result, error) {
	keeper.SetChannel(ctx, msg.ChainName, msg.SourcePort, msg.SourceChannel)
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

// func handleOracleRespondPacketData(ctx sdk.Context, packet oracle.OracleResponsePacketData, keeper Keeper) (*sdk.Result, error) {
// 	clientID := strings.Split(packet.ClientID, ":")
// 	if len(clientID) != 2 {
// 		return nil, sdkerrors.Wrapf(types.ErrUnknownClientID, "unknown client id %s", packet.ClientID)
// 	}
// 	orderID, err := strconv.ParseUint(clientID[1], 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rawResult, err := hex.DecodeString(packet.Result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result, err := types.DecodeResult(rawResult)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Assume multiplier should be 1000000
// 	order, err := keeper.GetOrder(ctx, orderID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// TODO: Calculate collateral percentage
// 	goldAmount := order.Amount[0].Amount.Int64() / int64(result.Px)
// 	if goldAmount == 0 {
// 		escrowAddress := types.GetEscrowAddress()
// 		err = keeper.BankKeeper.SendCoins(ctx, escrowAddress, order.Owner, order.Amount)
// 		if err != nil {
// 			return nil, err
// 		}
// 		order.Status = types.Completed
// 		keeper.SetOrder(ctx, orderID, order)
// 	} else {
// 		goldToken := sdk.NewCoin("gold", sdk.NewInt(goldAmount))
// 		keeper.BankKeeper.AddCoins(ctx, order.Owner, sdk.NewCoins(goldToken))
// 		order.Gold = goldToken
// 		order.Status = types.Active
// 		keeper.SetOrder(ctx, orderID, order)
// 	}
// 	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
// }

// handleCreateProduct handles a message to set product
func handleMsgCreateProduct(ctx sdk.Context, keeper Keeper, msg MsgCreateProduct) (*sdk.Result, error) {

	key := "Product-" + msg.ProductID

	if keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.ProductID)
	}

	var product = Product{
		ProductID:   msg.ProductID,
		Title:       msg.Title,
		Description: msg.Description,
		Category:    msg.Category,
		Images:      msg.Images,
		Owner:       msg.Signer,
		Selling:     false,
		SellID:      "",
	}

	keeper.SetProduct(ctx, key, product)
	return &sdk.Result{}, nil // return
}

// handleMsgUpdateProduct handles a message to set product
func handleMsgUpdateProduct(ctx sdk.Context, keeper Keeper, msg MsgUpdateProduct) (*sdk.Result, error) {

	key := "Product-" + msg.ProductID

	if !keeper.IsProductPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	owner, err := keeper.GetProductOwner(ctx, key)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	var newInfo = Product{
		ProductID:   msg.ProductID,
		Title:       msg.Title,
		Description: msg.Description,
		Category:    msg.Category,
		Images:      msg.Images,
		Owner:       msg.Signer,
		Selling:     false,
	}

	keeper.SetProduct(ctx, key, newInfo)
	return &sdk.Result{}, nil // return
}

// handleMsgCreateSell handles a message to set sell
func handleMsgCreateSell(ctx sdk.Context, keeper Keeper, msg MsgCreateSell) (*sdk.Result, error) {

	keySell := "Sell-" + msg.SellID
	keyProduct := "Product-" + msg.ProductID

	if keeper.IsSellPresent(ctx, keySell) {
		return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.SellID)
	}

	if !keeper.IsProductPresent(ctx, keyProduct) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	product, err := keeper.GetProduct(ctx, keyProduct)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(product.Owner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var sell = Sell{
		SellID:    msg.SellID,
		ProductID: msg.ProductID,
		Seller:    msg.Signer,
		MinPrice:  msg.MinPrice,
	}

	product.Selling = true
	product.SellID = msg.SellID

	keeper.SetProduct(ctx, keyProduct, product)
	keeper.SetSell(ctx, keySell, sell)
	return &sdk.Result{}, nil // return
}

// handleMsgUpdateSell handles a message to update sell
func handleMsgUpdateSell(ctx sdk.Context, keeper Keeper, msg MsgUpdateSell) (*sdk.Result, error) {

	key := "Sell-" + msg.SellID

	if !keeper.IsSellPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrSellDoesNotExist, msg.SellID)
	}

	sell, err := keeper.GetSell(ctx, key)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(sell.Seller) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	sell.MinPrice = msg.MinPrice

	keeper.SetSell(ctx, key, sell)
	return &sdk.Result{}, nil // return
}

// Handle a message to delete sell
func handleMsgDeleteSell(ctx sdk.Context, keeper Keeper, msg MsgDeleteSell) (*sdk.Result, error) {
	keySell := "Sell-" + msg.SellID

	if !keeper.IsSellPresent(ctx, keySell) {
		return nil, sdkerrors.Wrap(types.ErrSellDoesNotExist, msg.SellID)
	}

	sell, err := keeper.GetSell(ctx, keySell)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(sell.Seller) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	iterator := keeper.GetReservationsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			record, err := keeper.GetReservation(ctx, key)
			if err != nil {
				continue
			}
			if record.SellID == msg.SellID {
				keeper.DeleteReservation(ctx, key)
			}
		}
	}

	keyProduct := "Product-" + sell.ProductID
	product, err := keeper.GetProduct(ctx, keyProduct)
	if err != nil {
		return &sdk.Result{}, err
	}

	product.Selling = false
	product.SellID = ""

	keeper.SetProduct(ctx, keyProduct, product)
	keeper.DeleteSell(ctx, keySell)
	return &sdk.Result{}, nil
}

// handleMsgCreateReservation handles a message to set reservation
func handleMsgCreateReservation(ctx sdk.Context, keeper Keeper, msg MsgCreateReservation) (*sdk.Result, error) {

	key := "Reservation-" + msg.ReservationID

	if keeper.IsReservationPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrReservationAlreadyExists, msg.ReservationID)
	}

	if !keeper.IsSellPresent(ctx, "Sell-"+msg.SellID) {
		return nil, sdkerrors.Wrap(types.ErrSellDoesNotExist, msg.SellID)
	}

	var reservation = Reservation{
		ReservationID: msg.ReservationID,
		SellID:        msg.SellID,
		Buyer:         msg.Signer,
		Price:         msg.Price,
		Decide:        false,
	}

	keeper.SetReservation(ctx, key, reservation)
	return &sdk.Result{}, nil // return
}

// handleMsgUpdateReservation handles a message to set reservation
func handleMsgUpdateReservation(ctx sdk.Context, keeper Keeper, msg MsgUpdateReservation) (*sdk.Result, error) {

	keyReservation := "Reservation-" + msg.ReservationID

	if !keeper.IsReservationPresent(ctx, keyReservation) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	reservation, err := keeper.GetReservation(ctx, keyReservation)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(reservation.Buyer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	reservation.Price = msg.Price

	keeper.SetReservation(ctx, keyReservation, reservation)
	return &sdk.Result{}, nil // return
}

// Handle a message to delete reservation
func handleMsgDeleteReservation(ctx sdk.Context, keeper Keeper, msg MsgDeleteReservation) (*sdk.Result, error) {
	keyReservation := "Reservation-" + msg.ReservationID

	if !keeper.IsReservationPresent(ctx, keyReservation) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	reservation, err := keeper.GetReservation(ctx, keyReservation)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(reservation.Buyer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	keeper.DeleteReservation(ctx, keyReservation)
	return &sdk.Result{}, nil
}

// Handle a message to delete reservation
func handleMsgDecideSell(ctx sdk.Context, keeper Keeper, msg MsgDecideSell) (*sdk.Result, error) {
	keyReservation := "Reservation-" + msg.ReservationID

	if !keeper.IsReservationPresent(ctx, keyReservation) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	reservation, err := keeper.GetReservation(ctx, keyReservation)
	if err != nil {
		return &sdk.Result{}, err
	}

	keySell := "Sell-" + reservation.SellID
	sell, err := keeper.GetSell(ctx, keySell)

	if !msg.Signer.Equals(sell.Seller) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	reservation.Decide = true

	keeper.SetReservation(ctx, keyReservation, reservation)
	return &sdk.Result{}, nil
}

// Handle a message to delete reservation
func handleMsgPayReservation(ctx sdk.Context, keeper Keeper, msg MsgPayReservation) (*sdk.Result, error) {

	keyReservation := "Reservation-" + msg.ReservationID

	if !keeper.IsReservationPresent(ctx, keyReservation) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	reservation, err := keeper.GetReservation(ctx, keyReservation)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !reservation.Decide {
		return nil, sdkerrors.Wrap(types.ErrReservationNotDecided, msg.ReservationID)
	}

	keySell := "Sell-" + reservation.SellID
	if !keeper.IsSellPresent(ctx, keySell) {
		return nil, sdkerrors.Wrap(types.ErrSellDoesNotExist, reservation.SellID)
	}

	if !msg.Signer.Equals(reservation.Buyer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	sell, err := keeper.GetSell(ctx, keySell)
	if err != nil {
		return &sdk.Result{}, err
	}

	err = keeper.BankKeeper.SendCoins(ctx, reservation.Buyer, sell.Seller, reservation.Price)
	if err != nil {
		return nil, err
	}

	iterator := keeper.GetReservationsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			record, err := keeper.GetReservation(ctx, key)
			if err != nil {
				continue
			}
			if record.SellID == reservation.SellID {
				keeper.DeleteReservation(ctx, key)
			}
		}
	}

	keyProduct := "Product-" + sell.ProductID
	product, err := keeper.GetProduct(ctx, keyProduct)
	if err != nil {
		return &sdk.Result{}, err
	}

	product.Selling = false
	product.SellID = ""
	product.Owner = reservation.Buyer

	keeper.DeleteSell(ctx, keySell)
	keeper.SetProduct(ctx, keyProduct, product)
	return &sdk.Result{}, nil
}

// Handle a message to delete reservation
func handleMsgPayReservationByAtom(ctx sdk.Context, keeper Keeper, msg MsgPayReservationByAtom) (*sdk.Result, error) {

	bandChainID := "bandchain"
	port := "sunchain"
	oracleScriptID := oracle.OracleScriptID(2)
	calldata := encodeRequestParams("ATOM", 1000000)
	askCount := int64(1)
	minCount := int64(1)

	channelID, err := keeper.GetChannel(ctx, bandChainID, port)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"not found channel to bandchain",
		)
	}
	sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, port, channelID)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port goldcdp",
			channelID,
		)
	}

	destinationPort := sourceChannelEnd.Counterparty.PortID
	destinationChannel := sourceChannelEnd.Counterparty.ChannelID
	sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
		ctx, port, channelID,
	)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			channelID,
		)
	}
	packet := oracle.NewOracleRequestPacketData(
		fmt.Sprintf("Msg:%s", msg.ReservationID), oracleScriptID, calldata,
		askCount, minCount,
	)
	err = keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
		sequence, port, channelID, destinationPort, destinationChannel,
		1000000000,
	))
	if err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil

}

func handleOracleRespondPacketData(ctx sdk.Context, packet oracle.OracleResponsePacketData, keeper Keeper) (*sdk.Result, error) {
	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return nil, sdkerrors.Wrapf(types.ErrUnknownClientID, "unknown client id %s", packet.ClientID)
	}

	reservationID := clientID[1]
	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return nil, err
	}
	result, err := types.DecodeResult(rawResult)
	if err != nil {
		return nil, err
	}

	reservation, err := keeper.GetReservation(ctx, "Reservation-"+reservationID)
	if err != nil {
		return nil, err
	}

	denom := "transfer/ruahosxkxc/uatom"
	amount := reservation.Price[0].Amount.Int64() * 1000000000000 / int64(result.Px)
	token := sdk.NewCoin(denom, sdk.NewInt(amount))

	fmt.Print("============================", amount, "\n")

	keySell := "Sell-" + reservation.SellID
	sell, err := keeper.GetSell(ctx, keySell)

	keyProduct := "Product-" + sell.ProductID
	product, err := keeper.GetProduct(ctx, keyProduct)
	if err != nil {
		return nil, err
	}

	iterator := keeper.GetReservationsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())
		if "Reservation-" <= key && key <= "Reservation-zzzzzzzz" {
			record, err := keeper.GetReservation(ctx, key)
			if err != nil {
				continue
			}
			if record.SellID == reservation.SellID {
				keeper.DeleteReservation(ctx, key)
			}
		}

	}

	keeper.BankKeeper.SendCoins(ctx, reservation.Buyer, sell.Seller, sdk.NewCoins(token))

	product.Selling = false
	product.SellID = ""
	product.Owner = reservation.Buyer

	keeper.DeleteSell(ctx, keySell)
	keeper.SetProduct(ctx, keyProduct, product)

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
