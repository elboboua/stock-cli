package stockservice

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const ALPHA_ADVANTAGE_BASE_URL = "https://www.alphavantage.co/query"

type AlphaVantageStockService struct {
	ApiKey string
}

func (service *AlphaVantageStockService) GetQuoteBySymbol(symbol string) (StockQuote, error) {
	url, err := url.Parse(ALPHA_ADVANTAGE_BASE_URL)
	if err != nil {
		return StockQuote{}, err
	}
	queries := url.Query()
	queries.Add("function", "GLOBAL_QUOTE")
	queries.Add("symbol", symbol)
	queries.Add("apikey", service.ApiKey)
	url.RawQuery = queries.Encode()
	res, err := http.Get(url.String())
	if err != nil {
		return StockQuote{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return StockQuote{}, err
	}

	var globalQuoteResponse GlobalQuoteResponse
	err = json.Unmarshal(body, &globalQuoteResponse)
	if err != nil {
		return StockQuote{}, err
	}

	return StockQuote{
		Symbol: globalQuoteResponse.GlobalQuote.Symbol,
		Open: globalQuoteResponse.GlobalQuote.Open,
		High: globalQuoteResponse.GlobalQuote.High,
		Low: globalQuoteResponse.GlobalQuote.Low,
		Price: globalQuoteResponse.GlobalQuote.Price,
		Volume: globalQuoteResponse.GlobalQuote.Volume,
	}, nil
}


/* Types */

type GlobalQuote struct {
	Symbol string `json:"01. symbol"`
	Open string `json:"02. open"`
	High string `json:"03. high"`
	Low string `json:"04. low"`
	Price string `json:"05. price"`
	Volume string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose string `json:"08. previous close"`
	Change string `json:"09. change"`
	ChangePercent string `json:"10. change percent"`
}

type GlobalQuoteResponse struct {
	GlobalQuote GlobalQuote `json:"Global Quote"`
}