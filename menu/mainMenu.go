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
	fmt.Println("# 5 - Get Market by Market Currency")
	fmt.Println("# 6 - Get Market by Market Base")
	fmt.Println("# 7 - Get Market by Market Name")
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
		printMessage("Input the currency CODE")
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
	case 5:
		var marketCurrency string
		printMessage("Input the Market Currency")
		fmt.Scanln(&marketCurrency)
		fmt.Scanf("%s", &marketCurrency)
		markets := public.GetMarketsByMarketCurrency(marketCurrency)
		packages.ClearScreen()
		for index := range markets {
			market := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				market.MarketCurrency, market.BaseCurrency, market.MinTradeSize, market.MarketName)
		}
		backToMenu()
	case 6:
		var marketBase string
		printMessage("Input the Market Base")
		fmt.Scanln(&marketBase)
		fmt.Scanf("%s", &marketBase)
		markets := public.GetMarketsByBaseCurrency(marketBase)
		packages.ClearScreen()
		for index := range markets {
			market := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				market.MarketCurrency, market.BaseCurrency, market.MinTradeSize, market.MarketName)
		}
		backToMenu()
	case 7:
		var marketName string
		printMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		market := public.GetMarketByMarketName(marketName)
		packages.ClearScreen()
		fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
			market.MarketCurrency, market.BaseCurrency, market.MinTradeSize, market.MarketName)
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
