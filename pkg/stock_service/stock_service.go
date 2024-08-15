package stockservice

type StockQuote struct {
	Symbol string
	Open string
	High string
	Low string
	Price string
	Volume string
}

type StockService interface {
	GetQuoteBySymbol(symbol string) (StockQuote, error)
}

func NewStockService(apiKey string) StockService {


	return &AlphaVantageStockService{
		ApiKey: apiKey,
	}
}