package menu

import (
	"fmt"
	"Bleu/services/account"
	"Bleu/packages"
	"Bleu/wallets"
	"Bleu/services/public"
	"Bleu/services/market"
	"strconv"
	"Bleu/reader"
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
	fmt.Println("# 11 - Check Open Orders")
	fmt.Println("# 12 - Read Ethereum Transaction")
	fmt.Println("# Any other key to exit")
	PrintMessage("Input bellow your action code")
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
		PrintMessage("Input the currency CODE")
		fmt.Scanln(&currencyName)
		fmt.Scanf("%s", &currencyName)
		balance := account.GetBalanceByCurrency(currencyName)
		packages.ClearScreen()
		fmt.Printf("Currency: %s\nAvailable: %f\nAddress: %s\n", balance.Currency, balance.Available, balance.CryptoAddress)
		backToMenu()
	case 3:
		var amountValue float64
		PrintMessage("Input the amount of DOGE to transfer")
		fmt.Scanln(&amountValue)
		fmt.Scanf("%f", &amountValue)
		if success, message := account.Withdraw("DOGE", amountValue, wallets.Doge); success {
			PrintMessage(message)
		} else {
			PrintMessage(message)
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
		PrintMessage("Input the Market Currency")
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
		PrintMessage("Input the Market Base")
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
		PrintMessage("Input the Market Name CODE")
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
		PrintMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		packages.ClearScreen()
		PrintMessage("Input the rate")
		fmt.Scanln(&rate)
		fmt.Scanf("%f", &rate)
		packages.ClearScreen()
		PrintMessage("Input the quantity")
		fmt.Scanln(&quantity)
		fmt.Scanf("%f", &quantity)
		if success, order := market.PlaceSellOrder(marketName, rate, quantity); success {
			PrintMessage("Placed sell order with ID: " + strconv.FormatUint(order.Result.OrderId, 10))
		} else {
			PrintMessage(order.Message)
		}
		backToMenu()
	case 9:
		var marketName string
		var rate float64
		var quantity float64
		PrintMessage("Input the Market Name CODE")
		fmt.Scanln(&marketName)
		fmt.Scanf("%s", &marketName)
		packages.ClearScreen()
		PrintMessage("Input the rate")
		fmt.Scanln(&rate)
		fmt.Scanf("%f", &rate)
		packages.ClearScreen()
		PrintMessage("Input the quantity")
		fmt.Scanln(&quantity)
		fmt.Scanf("%f", &quantity)
		if success, order := market.PlaceBuyOrder(marketName, rate, quantity); success {
			PrintMessage("Placed buy order with ID: " + strconv.FormatUint(order.Result.OrderId, 10))
		} else {
			PrintMessage(order.Message)
		}
		backToMenu()
	case 10:
		var orderId uint64
		PrintMessage("Input the ORDER ID")
		fmt.Scanln(&orderId)
		fmt.Scanf("%d", &orderId)
		if success, message := market.CancelOrder(orderId); success {
			PrintMessage("Order canceled successfully.")
		} else {
			PrintMessage(message)
		}
		backToMenu()
	case 11:
		orders := market.GetOpenOrders()
		for index := range orders {
			order := orders[index]
			fmt.Printf("(%s - %s) ID: %d - Exchange: %s\nQty: %f - QtyRemaining: %f - QtyBaseTraded: %f - Price: %f\n\n",
				order.Type, order.Status, order.OrderId, order.Exchange, order.Quantity, order.QuantityRemaining, order.QuantityBaseTraded,
				order.Price)
		}
		backToMenu()
	case 12:
		var hash string
		PrintMessage("Input the Hash")
		fmt.Scanln(&hash)
		fmt.Scanf("%s", &hash)
		reader.GetMessageByHash(hash)
		backToMenu()
	}
}

func PrintMessage(message string) {
	fmt.Println("########################################################################")
	fmt.Println("#", message)
	fmt.Println("########################################################################")
}

func backToMenu() {
	var input string
	PrintMessage("Press any key to go back to Main Menu")
	fmt.Scanln(&input)
	fmt.Scanf("%s", &input)
	MainMenu()
}
