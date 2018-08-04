package main

import (
	"Bleu/services/account"
	"fmt"
)

func main() {
	balances := account.GetBalances()
	fmt.Println(balances)
}
