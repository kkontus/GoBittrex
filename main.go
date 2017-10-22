package main

import (
	"fmt"
	gbtHttp "GoBittrex/http"
	"os"
)

func main() {
	fmt.Printf("Hello Bittrex\n")

	argsWithCommand := os.Args
	argsWithoutCommand := os.Args[1:]
	arg := os.Args[1]
	fmt.Println(argsWithCommand)
	fmt.Println(argsWithoutCommand)
	fmt.Println(arg)

	status := selectBittrexEndpoint(arg)
	if !status {
		fmt.Println("Ëndpoint doesn't exist, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getMarkets")
		fmt.Println("./GoBittrex geCurrencies")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func selectBittrexEndpoint(param string) bool {
	status := false
	switch param {
	case "getMarkets":
		//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getmarkets")
		gbtHttp.GetMarketsDataJson("https://bittrex.com/api/v1.1/public/getmarkets")
		status = true
	case "getCurrencies":
		//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getcurrencies")
		gbtHttp.GetCurrenciesDataJson("https://bittrex.com/api/v1.1/public/getcurrencies")
		status = true
	default:
		status = false
	}
	return status
}
