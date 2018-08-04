package market

import (
	"Bleu/packages"
	"net/http"
	"strconv"
	"encoding/json"
	"Bleu/services"
	"strings"
)

type CancelOrderResponse struct {
	services.Response
	Result string `json:"result"`
}

type PlaceOrderResponse struct {
	services.Response
	Result PlaceOrder `json:"result"`
}

type PlaceOrder struct {
	OrderId uint64 `json:"orderid,string"`
}

type OpenOrderResponse struct {
	services.Response
	Result []OpenOrder `json:"result"`
}

type OpenOrder struct {
	OrderId            uint64 `json:",string"`
	Exchange           string
	Type               string
	Quantity           float64 `json:",string"`
	QuantityRemaining  float64 `json:",string"`
	QuantityBaseTraded float64 `json:",string"`
	Price              float64 `json:",string"`
	Status             string
	Created            string
	Comments           string
}

const baseURI = "market"

func PlaceSellOrder(marketName string, rate float64, quantity float64) (bool, PlaceOrderResponse) {
	marketName = strings.ToUpper(marketName)
	sellURI := "/selllimit"
	params := "&market=" + marketName
	params += "&rate=" + strconv.FormatFloat(rate, 'f', -1, 64)
	params += "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64)
	signature, uri := packages.GetAPISign(baseURI+sellURI, params)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson PlaceOrderResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Success, responseJson
}

func PlaceBuyOrder(marketName string, rate float64, quantity float64) (bool, PlaceOrderResponse) {
	marketName = strings.ToUpper(marketName)
	sellURI := "/buylimit"
	params := "&market=" + marketName
	params += "&rate=" + strconv.FormatFloat(rate, 'f', -1, 64)
	params += "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64)
	signature, uri := packages.GetAPISign(baseURI+sellURI, params)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson PlaceOrderResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Success, responseJson
}

func CancelOrder(orderId uint64) (bool, string) {
	cancelURI := "/cancel"
	params := "&orderid=" + strconv.FormatUint(orderId, 10)
	signature, uri := packages.GetAPISign(baseURI+cancelURI, params)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson CancelOrderResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Success, responseJson.Message
}

func GetOpenOrders() []OpenOrder {
	ordersURI := "/getopenorders"
	signature, uri := packages.GetAPISign(baseURI + ordersURI)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson OpenOrderResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Result
}
