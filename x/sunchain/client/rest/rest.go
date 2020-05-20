package rest

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

const (
	restName        = "name"
	restProduct     = "product"
	restSell        = "sell"
	restReservation = "reservation"
	restOwner       = "owner"

	accName    = "name"
	accAddress = "address"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {

	r.HandleFunc(fmt.Sprintf("/%s/tx/sign", storeName), signTxHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/tx/sign", storeName), signTxHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), createProductHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), createProductHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), productsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/products/{%s}", storeName, restProduct), getProductHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), updateProductHandler(cliCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), createSellHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), createSellHandler(cliCtx)).Methods("OPTIONS")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), sellsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells/{%s}", storeName, restSell), getSellHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), updateSellHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/cancelSell", storeName), deleteSellHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/cancelSell", storeName), deleteSellHandler(cliCtx)).Methods("OPTIONS")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), sellsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells/{%s}/reservations", storeName, restSell), reservationsBySellIDHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells/decideSell", storeName), decideSellHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/sells/decideSell", storeName), decideSellHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), createReservationHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), createReservationHandler(cliCtx)).Methods("OPTIONS")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), reservationsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/{%s}", storeName, restReservation), getSellHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), updateReservationHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), deleteReservationHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/payReservation", storeName), payReservationHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/payReservation", storeName), payReservationHandler(cliCtx)).Methods("OPTIONS")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/payReservationByAtom", storeName), payReservationByAtomHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/payReservationByAtom", storeName), payReservationByAtomHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/address", storeName, accName), accAddressHandler(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/products", storeName, accName), productsByOwnerHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/balance", storeName, accName), queryBalanceHandler(cliCtx)).Methods("GET")
}
