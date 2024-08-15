package cmd

import (
	"fmt"

	"github.com/elboboua/stock-cli/pkg/config"
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
		config := config.LoadConfig()

		if symbol == "" {
			fmt.Println("Please provide a symbol")
			return
		}
		
		stockService := stockservice.NewStockService(config.StockApiKey)
		res, err := stockService.GetQuoteBySymbol(symbol)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("The price of %s is %s\n", res.Symbol, res.Price)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stock quote",
	Long: `Get a stock quote from a stock service.
	Currently, only Alpha Vantage is supported.`,
}

func init() {
	getPriceCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "The symbol of the stock")
	getCmd.AddCommand(getPriceCmd)
	rootCmd.AddCommand(getCmd)
}