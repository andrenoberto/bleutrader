package menu

import (
	"fmt"
	"Bleu/services/account"
	"Bleu/packages"
	"Bleu/wallets"
	"Bleu/services/public"
)

func MainMenu() {
	packages.ClearScreen()
	fmt.Println("########################################################################")
	fmt.Println("# 1 - Get All Balances")
	fmt.Println("# 2 - Get Balance by Currency")
	fmt.Println("# 3 - Make Withdraw")
	fmt.Println("# 4 - See Markets")
	fmt.Println("# Any other key to exit")
	printMessage("Input bellow your action code")
	var option uint8
	fmt.Scanf("%d", &option)
	switchMenu(option)
}

func switchMenu(option uint8) {
	packages.ClearScreen()
	switch option {
	case 1:
		balances := account.GetBalances()
		for index := range balances {
			balance := balances[index]
			fmt.Printf("Currency: %s - Available: %f - Address: %s\n", balance.Currency, balance.Available, balance.CryptoAddress)
		}
		backToMenu()
	case 2:
		var currencyName string
		printMessage("Input the currency code")
		fmt.Scanln(&currencyName)
		fmt.Scanf("%s", &currencyName)
		balance := account.GetBalanceByCurrency(currencyName)
		packages.ClearScreen()
		fmt.Printf("Currency: %s\nAvailable: %f\nAddress: %s\n", balance.Currency, balance.Available, balance.CryptoAddress)
		backToMenu()
	case 3:
		var amountValue float64
		printMessage("Input the amount of DOGE to transfer")
		fmt.Scanln(&amountValue)
		fmt.Scanf("%f", &amountValue)
		if success, message := account.Withdraw("DOGE", amountValue, wallets.Doge); success {
			printMessage(message)
		} else {
			printMessage(message)
		}
		backToMenu()
	case 4:
		markets := public.GetMarkets()
		for index := range markets {
			market := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				market.MarketCurrency, market.BaseCurrency, market.MinTradeSize, market.MarketName)
		}
		backToMenu()
	}
}

func printMessage(message string) {
	fmt.Println("########################################################################")
	fmt.Println("#", message)
	fmt.Println("########################################################################")
}

func backToMenu() {
	var input string
	printMessage("Press any key to go back to Main Menu")
	fmt.Scanln(&input)
	fmt.Scanf("%s", &input)
	MainMenu()
}
