package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StockApiKey string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return Config{
		StockApiKey: os.Getenv("STOCK_API_KEY"),
	}
}