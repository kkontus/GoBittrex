package main

import (
	"fmt"
	"os"
	gbtHttp "GoBittrex/http"
	gbtUtil "GoBittrex/util"
)

const (
	API_BASE       = "https://bittrex.com/api/" // Bittrex API endpoint
	API_VERSION    = "v1.1"
	API_VERSION_V2 = "v2.0"
	WS_BASE        = "socket.bittrex.com" // Bittrex WS API endpoint
	WS_HUB         = "CoreHub"            // SignalR main hub
	API_PATH       = API_BASE + API_VERSION
	API_PATH_V2    = API_BASE + API_VERSION_V2
)

func main() {
	fmt.Printf("Hello Bittrex\n")

	arg := ""
	coin := ""
	if len(os.Args) == 2 {
		arg = os.Args[1]
		coin = ""
	} else if len(os.Args) == 3 {
		arg = os.Args[1]
		coin = os.Args[2]
	}

	fmt.Println(arg)
	fmt.Println(coin)

	status := selectBittrexEndpoint(arg, coin)
	if !status {
		fmt.Println("Ëndpoint doesn't exist, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getMarkets")
		fmt.Println("./GoBittrex geCurrencies")
		fmt.Println("./GoBittrex getTicks <coin_symbol>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func selectBittrexEndpoint(param string, coin string) bool {
	status := false
	switch param {
	case "getMarkets":
		//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getmarkets")
		//URL := API_PATH + "/public/getmarkets"
		URL := fmt.Sprintf("%s/public/getmarkets", API_PATH)
		gbtHttp.GetMarketsDataJson(URL)
		status = true
	case "getCurrencies":
		//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getcurrencies")
		//URL := API_PATH + "/public/getcurrencies"
		URL := fmt.Sprintf("%s/public/getcurrencies", API_PATH)
		gbtHttp.GetCurrenciesDataJson(URL)
		status = true
	case "getTicks":
		//gbtHttp.GetData(API_PATH_V2+"/pub/market/GetTicks?marketName=BTC-"+coin+"&tickInterval=day&_="+gbtUtil.FormatInt(gbtUtil.GetTimestamp()))
		//URL := API_PATH_V2 + "/pub/market/GetTicks?marketName=BTC-" + coin + "&tickInterval=day&_=" + gbtUtil.FormatInt(gbtUtil.GetTimestamp())
		URL := fmt.Sprintf("%s/pub/market/GetTicks?marketName=BTC-%s&tickInterval=day&_=%s", API_PATH_V2, coin, gbtUtil.FormatInt(gbtUtil.GetTimestamp()))
		gbtHttp.GetTicksDataJson(URL)
		status = true
	default:
		status = false
	}
	return status
}
