package public

import "Bleu/services"

type MarketResponse struct {
	services.Response
	Result []Market `json:"result"`
}

type Market struct {
	MarketCurrency     string
	BaseCurrency       string
	MarketCurrencyLong string
	BaseCurrencyLong   string
	MinTradeSize       float64 `json:",string"`
	MarketName         string
	IsActive           bool `json:",string"`
}