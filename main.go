package main

import (
	"Bleu/services/account"
	"fmt"
)

func main() {
	btc := account.GetBalanceByCurrency("BTC")
	fmt.Println(btc)
	account.Withdraw("BTC", 0.001, "1AqABNPPPurq8Sy6RuxWkjRHdNLNzaR1bZ")
}
