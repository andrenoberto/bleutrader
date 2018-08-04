package account

import (
	"Bleu/packages"
	"fmt"
	"encoding/json"
)

type BalanceResponse struct {
	Success string    `json:"success"`
	Message string    `json:"message"`
	Result  []Balance `json:"result"`
}

type Balance struct {
	Currency      string
	Balance       string
	Available     string
	Pending       string
	CryptoAddress string
	IsActive      string
	AllowDeposit  string
	AllowWithdraw string
}

const baseURI = "account"

func GetBalances() []Balance {
	balanceURI := "/getbalances"
	signature, uri := packages.GetAPISign(baseURI + balanceURI)
	response := packages.RequestHandler("GET", uri, nil, signature)
	var responseJson BalanceResponse
	err := json.Unmarshal(response, &responseJson)
	packages.ErrorHandler(err)
	fmt.Println(responseJson.Result[0])
	return responseJson.Result
}
