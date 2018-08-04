package menu

import (
	"fmt"
	"Bleu/services/account"
	"Bleu/packages"
	"Bleu/wallets"
)

func MainMenu() {
	packages.ClearScreen()
	fmt.Println("##################################")
	fmt.Println("# 1 - Get All Balances")
	fmt.Println("# 2 - Get Balance by Currency")
	fmt.Println("# 3 - Make Withdraw")
	printMessage("Input bellow your action")
	var option uint8
	fmt.Scanf("%d", &option)
	switchMenu(option)
}

func switchMenu(option uint8) {
	packages.ClearScreen()
	switch option {
	default:
		printMessage("Invalid option selected")
		printMessageAndWait("Press Enter key to continue")
		//MainMenu()
	case 1:
		balances := account.GetBalances()
		fmt.Println(balances)
	case 2:
		printMessage("Input the currency code")
		var currencyName string
		fmt.Scanf("%s", &currencyName)
		balance := account.GetBalanceByCurrency("DOGE")
		fmt.Println(balance)
	case 3:
		/*withdraw := */account.Withdraw("DOGE", 11, wallets.Doge)
		//fmt.Println(withdraw)
	}
}

func printMessage(message string) {
	fmt.Println("##################################")
	fmt.Println("#", message)
	fmt.Println("##################################")
}

func printMessageAndWait(message string) {
	printMessage(message)
	fmt.Scanln()
}
