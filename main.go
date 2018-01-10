package main

import (
	"fmt"
	"os"
	gbtTasks "GoBittrex/endpoints"
	gbtValidator "GoBittrex/validator"
)

func main() {
	fmt.Printf("Hello Bittrex\n")

	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	status := gbtTasks.SelectBittrexEndpoint(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getMarkets")
		fmt.Println("./GoBittrex geCurrencies")
		fmt.Println("./GoBittrex getTicks <coin_symbol>")
		fmt.Println("./GoBittrex getOpenOrders <coin_symbol>")
		fmt.Println("./GoBittrex startTrailing <coin_symbol> SL <sell_loss> TP <take_profit> TTP <trailing_take_profit>")
		fmt.Println("./GoBittrex getOrderBook <coin_symbol>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}
