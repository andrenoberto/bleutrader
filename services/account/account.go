package account

import (
	"Bleu/packages"
		"encoding/json"
		"net/http"
)

type BalanceResponse struct {
	Success string    `json:"success"`
	Message string    `json:"message"`
	Result  []Balance `json:"result"`
}

type Balance struct {
	Currency      string
	Balance       float64
	Available     float64
	Pending       float64
	CryptoAddress string
	IsActive      bool
	AllowDeposit  bool
	AllowWithdraw bool
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

/*func Withdraw() {
	withdrawURI := "/withdraw"
	signature, uri := paca
}*/
