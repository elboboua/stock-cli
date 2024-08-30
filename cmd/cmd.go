package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/elboboua/stock-cli/pkg/config"
	stockservice "github.com/elboboua/stock-cli/pkg/stock_service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stock-cli",
	Short: "A CLI tool to get stock quotes",
	Long: `A CLI tool to get stock quotes from different stock services.
	Currently, only Alpha Vantage is supported.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		ctx := context.Background()
		stockService := stockservice.NewStockService(config.StockApiKey)
		ctx = context.WithValue(ctx, "stockService", stockService)
		cmd.SetContext(ctx)
	},
}

func Execute() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
