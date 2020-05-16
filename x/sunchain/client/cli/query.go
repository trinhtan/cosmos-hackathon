package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/trinhtan/cosmos-hackathon/x/sunchain/types"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	sunchainCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the sunchain module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	sunchainCmd.AddCommand(flags.GetCommands(
		GetCmdReadOrder(storeKey, cdc),
		GetCmdProduct(storeKey, cdc),
		GetCmdProducts(storeKey, cdc),
		GetCmdSell(storeKey, cdc),
		GetCmdSells(storeKey, cdc),
		GetCmdReservation(storeKey, cdc),
		GetCmdReservations(storeKey, cdc),
	)...)

	return sunchainCmd
}

// GetCmdReadOrder queries order by orderID
func GetCmdReadOrder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "order",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orderID := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/order/%s", queryRoute, orderID),
				nil,
			)
			if err != nil {
				fmt.Printf("read request fail - %s \n", orderID)
				return nil
			}

			var order types.Order
			if err := cdc.UnmarshalJSON(res, &order); err != nil {
				return err
			}
			return cliCtx.PrintOutput(order)
		},
	}
}

// GetCmdProduct queries information about a product
func GetCmdProduct(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "product [productID]",
		Short: "Query info of product",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			productID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/product/%s", queryRoute, productID), nil)
			if err != nil {
				fmt.Printf("could not get product - %s \n", productID)
				return nil
			}

			var out types.Product
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdProducts queries a list of all products
func GetCmdProducts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "products",
		Short: "products",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/products", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query products\n")
				return nil
			}

			var out types.QueryResProducts
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdProduct queries information about a product
func GetCmdSell(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sell [sellID]",
		Short: "Query info of product",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			sellID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/sell/%s", queryRoute, sellID), nil)
			if err != nil {
				fmt.Printf("could not get sell - %s \n", sellID)
				return nil
			}

			var out types.Sell
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdProducts queries a list of all products
func GetCmdSells(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sells",
		Short: "sells",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/sells", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query sells\n")
				return nil
			}

			var out types.QueryResSells
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdReservation queries information about a product
func GetCmdReservation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "reservation [reservationID]",
		Short: "Query info of reservation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			reservationID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/reservation/%s", queryRoute, reservationID), nil)
			if err != nil {
				fmt.Printf("could not get resertvation - %s \n", reservationID)
				return nil
			}

			var out types.Reservation
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdProducts queries a list of all products
func GetCmdReservations(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "reservations",
		Short: "reservations",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/reservations", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query reservations\n")
				return nil
			}

			var out types.QueryResReservations
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
}
