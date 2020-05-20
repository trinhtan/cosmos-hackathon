package rest

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

const (
	dir = "/home/ngo.van.nghia/Documents/POC/cosmos-hackathon/"
)

type createProducteReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Category    string       `json:"category"`
	Images      string       `json:"images"`
}

func createProductHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req createProducteReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		signer := req.BaseReq.From

		addr, err := sdk.AccAddressFromBech32(signer)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		b := make([]byte, 16)
		_, err = rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		productID := fmt.Sprintf("%x-%x-%x-%x-%x",
			b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		msg := types.NewMsgCreateProduct(productID, req.Title, req.Description, req.Category, req.Images, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type updateProductReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	ProductID   string       `json:"productID"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Category    string       `json:"category"`
	Images      string       `json:"images"`
}

func updateProductHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateProductReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgUpdateProduct(req.ProductID, req.Title, req.Description, req.Category, req.Images, addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type createSellReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	ProductID string       `json:"productID"`
	// Signer    string       `json:"signer"`
	MinPrice string `json:"minPrice"`
}

func createSellHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req createSellReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.MinPrice)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		b := make([]byte, 16)
		_, err = rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		sellID := fmt.Sprintf("%x-%x-%x-%x-%x",
			b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		// create the message
		msg := types.NewMsgCreateSell(sellID, req.ProductID, addr, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type updateSellReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	SellID   string       `json:"sellID"`
	MinPrice string       `json:"minPrice"`
}

func updateSellHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateSellReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.MinPrice)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgUpdateSell(req.SellID, addr, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteSellReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	SellID  string       `json:"sellID"`
}

func deleteSellHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req deleteSellReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgDeleteSell(req.SellID, addr)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type createReservationReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	SellID  string       `json:"sellID"`
	Price   string       `json:"price"`
}

func createReservationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req createReservationReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.Price)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		b := make([]byte, 16)
		_, err = rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		reservationID := fmt.Sprintf("%x-%x-%x-%x-%x",
			b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		// create the message
		msg := types.NewMsgCreateReservation(reservationID, req.SellID, addr, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type updateReservationReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReservationID string       `json:"reservationID"`
	Price         string       `json:"price"`
}

func updateReservationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateReservationReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.Price)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgUpdateReservation(req.ReservationID, addr, coins)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteReservationReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReservationID string       `json:"reservationID"`
}

func deleteReservationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req deleteReservationReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgDeleteReservation(req.ReservationID, addr)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type signTxReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Tx            string       `json:"tx"`
	Sequence      string       `json:"sequence"`
	AccountNumber string       `json:"accountNumber"`
}

func signTxHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var req signTxReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		filePath := dir + "unsignedTx.json"
		data := []byte(req.Tx)

		err := ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("bccli", "tx", "sign", filePath, "--from", req.BaseReq.From, "--offline", "--chain-id", "band-consumer", "--sequence", req.Sequence, "--account-number", req.AccountNumber)
		stdout, err := cmd.Output()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		filePath = dir + "signedTx.json"
		data = []byte(string(stdout))
		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			panic(err)
		}

		cmd = exec.Command("bccli", "tx", "broadcast", filePath)
		stdout, err = cmd.Output()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, string(stdout))
	}
}

type decideSellReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReservationID string       `json:"reservationID"`
}

func decideSellHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req deleteReservationReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgDecideSell(req.ReservationID, addr)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type payReservationReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReservationID string       `json:"reservationID"`
}

func payReservationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req deleteReservationReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgPayReservation(req.ReservationID, addr)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type payReservationByAtomReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReservationID string       `json:"reservationID"`
}

func payReservationByAtomHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var req payReservationByAtomReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgPayReservationByAtom(req.ReservationID, addr)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
