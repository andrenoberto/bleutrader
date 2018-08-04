package public

import (
	"Bleu/services"
	"Bleu/packages"
	"net/http"
	"encoding/json"
	"strings"
)

type MarketResponse struct {
	services.Response
	Result []Market `json:"result"`
}

type Market struct {
	MarketCurrency     string
	BaseCurrency       string
	MarketCurrencyLong string
	BaseCurrencyLong   string
	MinTradeSize       float64
	MarketName         string
	IsActive           bool `json:",string"`
}

const baseURI = "public"

func GetMarkets() []Market {
	balanceURI := "/getmarkets"
	signature, uri := packages.GetAPISign(baseURI + balanceURI)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson MarketResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Result
}

func GetMarketsByMarketCurrency(marketCurrency string) []Market {
	filteredMarkets := make([]Market, 0)
	marketCurrency = strings.ToUpper(marketCurrency)
	markets := GetMarkets()
	for index := range markets {
		if markets[index].MarketCurrency == marketCurrency {
			filteredMarkets = append(filteredMarkets, markets[index])
		}
	}
	return filteredMarkets
}

func GetMarketsByBaseCurrency(baseCurrency string) []Market {
	filteredMarkets := make([]Market, 0)
	baseCurrency = strings.ToUpper(baseCurrency)
	markets := GetMarkets()
	for index := range markets {
		if markets[index].BaseCurrency == baseCurrency {
			filteredMarkets = append(filteredMarkets, markets[index])
		}
	}
	return filteredMarkets
}

func GetMarketByMarketName(marketName string) Market {
	marketName = strings.ToUpper(marketName)
	markets := GetMarkets()
	for index := range markets {
		if markets[index].MarketName == marketName {
			return markets[index]
		}
	}
	return Market{}
}