package goldcdp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"

	"github.com/bandprotocol/goldcdp/x/goldcdp/types"
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
		case MsgCreateReservation:
			return handleMsgCreateReservation(ctx, keeper, msg)
		case MsgUpdateReservation:
			return handleMsgUpdateReservation(ctx, keeper, msg)
		case MsgDeleteReservation:
			return handleMsgDeleteReservation(ctx, keeper, msg)
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
	// TODO: Set all bandchain parameter here
	bandChainID := "bandchain"
	port := "goldcdp"
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

func handleOracleRespondPacketData(ctx sdk.Context, packet oracle.OracleResponsePacketData, keeper Keeper) (*sdk.Result, error) {
	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return nil, sdkerrors.Wrapf(types.ErrUnknownClientID, "unknown client id %s", packet.ClientID)
	}
	orderID, err := strconv.ParseUint(clientID[1], 10, 64)
	if err != nil {
		return nil, err
	}
	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return nil, err
	}
	result, err := types.DecodeResult(rawResult)
	if err != nil {
		return nil, err
	}

	// Assume multiplier should be 1000000
	order, err := keeper.GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	// TODO: Calculate collateral percentage
	goldAmount := order.Amount[0].Amount.Int64() / int64(result.Px)
	if goldAmount == 0 {
		escrowAddress := types.GetEscrowAddress()
		err = keeper.BankKeeper.SendCoins(ctx, escrowAddress, order.Owner, order.Amount)
		if err != nil {
			return nil, err
		}
		order.Status = types.Completed
		keeper.SetOrder(ctx, orderID, order)
	} else {
		goldToken := sdk.NewCoin("gold", sdk.NewInt(goldAmount))
		keeper.BankKeeper.AddCoins(ctx, order.Owner, sdk.NewCoins(goldToken))
		order.Gold = goldToken
		order.Status = types.Active
		keeper.SetOrder(ctx, orderID, order)
	}
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

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
	}

	keeper.SetProduct(ctx, key, newInfo)
	return &sdk.Result{}, nil // return
}

// handleMsgCreateSell handles a message to set sell
func handleMsgCreateSell(ctx sdk.Context, keeper Keeper, msg MsgCreateSell) (*sdk.Result, error) {

	key := "Sell-" + msg.SellID

	if keeper.IsSellPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrProductAlreadyExists, msg.SellID)
	}

	if !keeper.IsProductPresent(ctx, "Product-"+msg.ProductID) {
		return nil, sdkerrors.Wrap(types.ErrProductDoesNotExist, msg.ProductID)
	}

	productOwner, err := keeper.GetProductOwner(ctx, "Product-"+msg.ProductID)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(productOwner) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var sell = Sell{
		SellID:    msg.SellID,
		ProductID: msg.ProductID,
		Seller:    msg.Signer,
		MinPrice:  msg.MinPrice,
	}

	keeper.SetSell(ctx, key, sell)
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

	var newInfo = Sell{
		SellID:    msg.SellID,
		ProductID: sell.ProductID,
		Seller:    sell.Seller,
		MinPrice:  msg.MinPrice,
	}

	keeper.SetSell(ctx, key, newInfo)
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
	}

	keeper.SetReservation(ctx, key, reservation)
	return &sdk.Result{}, nil // return
}

// handleMsgUpdateReservation handles a message to set reservation
func handleMsgUpdateReservation(ctx sdk.Context, keeper Keeper, msg MsgUpdateReservation) (*sdk.Result, error) {

	key := "Reservation-" + msg.ReservationID

	if !keeper.IsReservationPresent(ctx, key) {
		return nil, sdkerrors.Wrap(types.ErrReservationDoesNotExist, msg.ReservationID)
	}

	reservation, err := keeper.GetReservation(ctx, key)
	if err != nil {
		return &sdk.Result{}, err
	}

	if !msg.Signer.Equals(reservation.Buyer) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	var newInfo = Reservation{
		SellID:        reservation.SellID,
		ReservationID: msg.ReservationID,
		Buyer:         reservation.Buyer,
		Price:         msg.Price,
	}

	keeper.SetReservation(ctx, key, newInfo)
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
