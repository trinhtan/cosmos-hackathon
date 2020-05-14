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

	accName = "name"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), namesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), setNameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/names/setDescription", storeName), setDescriptionHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}", storeName, restName), resolveNameHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/whois", storeName, restName), whoIsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/description", storeName, restName), getDescriptionHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), deleteNameHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/%s/tx/sign", storeName), signTxHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/tx/sign", storeName), signTxHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), createProductHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), createProductHandler(cliCtx)).Methods("OPTIONS")

	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), productsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/products/{%s}", storeName, restProduct), getProductHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/products", storeName), updateProductHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/products/decideSell", storeName), changeProductOwnerHandler(cliCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), createSellHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), sellsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells/{%s}", storeName, restSell), getSellHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), updateSellHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), deleteSellHandler(cliCtx)).Methods("DELETE")
	r.HandleFunc(fmt.Sprintf("/%s/sells", storeName), sellsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/sells/{%s}/reservations", storeName, restSell), reservationsBySellIDHandler(cliCtx, storeName)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), createReservationHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), reservationsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reservations/{%s}", storeName, restReservation), getSellHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), updateReservationHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/reservations", storeName), deleteReservationHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/%s/accAddress/{%s}", storeName, accName), accAddressHandler(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/accAddress/{%s}/products", storeName, restOwner), productsByOwnerHandler(cliCtx, storeName)).Methods("GET")
}
