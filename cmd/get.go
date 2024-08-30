package cmd

import (
	"fmt"
	"strconv"

	stockservice "github.com/elboboua/stock-cli/pkg/stock_service"
	"github.com/spf13/cobra"
)

var symbol string
var getPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "Get the price of a stock",
	Long: `Get the price of a stock from a stock service.
	Currently, only Alpha Vantage is supported.`,
	Run: func(cmd *cobra.Command, args []string) {

		if symbol == "" {
			fmt.Println("Please provide a symbol")
			return
		}
		ss := cmd.Context().Value("stockService").(stockservice.StockService)
		res, err := ss.GetQuoteBySymbol(symbol)
		if err != nil {
			fmt.Println(err)
			return
		}
		priceF, err := strconv.ParseFloat(res.Price, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("The price of %s is $%s\n", res.Symbol, strconv.FormatFloat(priceF, 'f', 2, 64))
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a variety of stock information",
	Long:  `Get a variety of stock information from a stock service.`,
}

func init() {
	getPriceCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "The symbol of the stock")
	getCmd.AddCommand(getPriceCmd)
	rootCmd.AddCommand(getCmd)
}
