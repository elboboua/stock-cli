package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stock-cli",
	Short: "A CLI tool to get stock quotes",
	Long: `A CLI tool to get stock quotes from different stock services.
	Currently, only Alpha Vantage is supported.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }