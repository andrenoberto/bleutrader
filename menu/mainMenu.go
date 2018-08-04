package menu

import (
	"fmt"
	"Bleu/services/account"
	"Bleu/packages"
	"Bleu/wallets"
	"Bleu/services/public"
	"Bleu/services/market"
	"strconv"
)

func MainMenu() {
	packages.ClearScreen()
	fmt.Println("########################################################################")
	fmt.Println("# 01 - Get All Balances")
	fmt.Println("# 02 - Get Balance by Currency")
	fmt.Println("# 03 - Make Withdraw")
	fmt.Println("# 04 - See Markets")
	fmt.Println("# 05 - Get Market by Market Currency")
	fmt.Println("# 06 - Get Market by Market Base")
	fmt.Println("# 07 - Get Market by Market Name")
	fmt.Println("# 08 - Place Sell Order")
	fmt.Println("# 09 - Place Buy Order")
	fmt.Println("# 10 - Cancel an Order")
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
			marketItem := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				marketItem.MarketCurrency, marketItem.BaseCurrency, marketItem.MinTradeSize, marketItem.MarketName)
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
			marketItem := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				marketItem.MarketCurrency, marketItem.BaseCurrency, marketItem.MinTradeSize, marketItem.MarketName)
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
			marketItem := markets[index]
			fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
				marketItem.MarketCurrency, marketItem.BaseCurrency, marketItem.MinTradeSize, marketItem.MarketName)
		}
		backToMenu()
	case 7:
		var marketName string
		printMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		marketItem := public.GetMarketByMarketName(marketName)
		packages.ClearScreen()
		fmt.Printf("Market: %s - Base: %s - Min. Trade: %f - Market Name: %s\n",
			marketItem.MarketCurrency, marketItem.BaseCurrency, marketItem.MinTradeSize, marketItem.MarketName)
		backToMenu()
	case 8:
		var marketName string
		var rate float64
		var quantity float64
		printMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		packages.ClearScreen()
		printMessage("Input the rate")
		fmt.Scanln(&rate)
		fmt.Scanf("%f", &rate)
		packages.ClearScreen()
		printMessage("Input the quantity")
		fmt.Scanln(&quantity)
		fmt.Scanf("%f", &quantity)
		if success, order := market.PlaceSellOrder(marketName, rate, quantity); success {
			printMessage("Placed sell order with ID: " + strconv.FormatUint(order.Result.OrderId, 10))
		} else {
			printMessage(order.Message)
		}
		backToMenu()
	case 9:
		var marketName string
		var rate float64
		var quantity float64
		printMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		packages.ClearScreen()
		printMessage("Input the rate")
		fmt.Scanln(&rate)
		fmt.Scanf("%f", &rate)
		packages.ClearScreen()
		printMessage("Input the quantity")
		fmt.Scanln(&quantity)
		fmt.Scanf("%f", &quantity)
		if success, order := market.PlaceBuyOrder(marketName, rate, quantity); success {
			printMessage("Placed buy order with ID: " + strconv.FormatUint(order.Result.OrderId, 10))
		} else {
			printMessage(order.Message)
		}
		backToMenu()
	case 10:
		var orderId uint64
		printMessage("Input the ORDER ID")
		fmt.Scanln(&orderId)
		fmt.Scanf("%d", &orderId)
		if success, message := market.CancelOrder(orderId); success {
			printMessage("Order canceled successfully.")
		} else {
			printMessage(message)
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
