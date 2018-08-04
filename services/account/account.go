package account

import (
	"Bleu/packages"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"Bleu/services"
)

type BalanceResponse struct {
	services.Response
	Result []Balance `json:"result"`
}

type WithdrawResponse struct {
	services.Response
	Result []string `json:"result"`
}

type Balance struct {
	Currency      string
	Balance       float64 `json:",string"`
	Available     float64 `json:",string"`
	Pending       float64 `json:",string"`
	CryptoAddress string
	IsActive      bool    `json:",string"`
	AllowDeposit  bool    `json:",string"`
	AllowWithdraw bool    `json:",string"`
}

const baseURI = "account"

func GetBalances() []Balance {
	balanceURI := "/getbalances"
	signature, uri := packages.GetAPISign(baseURI + balanceURI)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson BalanceResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Result
}

func GetBalanceByCurrency(currencyName string) Balance {
	currencyName = strings.ToUpper(currencyName)
	balances := GetBalances()
	for index := range balances {
		if balances[index].Currency == currencyName {
			return balances[index]
		}
	}
	return Balance{}
}

func Withdraw(currency string, quantity float64, address string) (bool, string) {
	withdrawURI := "/withdraw"
	params := "&currency=" + currency
	params += "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64)
	params += "&address=" + address
	signature, uri := packages.GetAPISign(baseURI+withdrawURI, params)
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("apisign", signature)
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson BalanceResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	return responseJson.Success, responseJson.Message
}
