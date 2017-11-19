package main

import (
	"fmt"
	"os"
	"time"
	gbtHttp "GoBittrex/http"
	gbtUtil "GoBittrex/util"
	gbtFalgo "GoBittrex/falgo"
	gbtConfig "GoBittrex/config"
	gbtScheduler "GoBittrex/scheduler"
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

	// test data, we will take this from getTicks
	l := []float64{50, 98, 93, 77, 82, 83}
	h := []float64{115, 198, 193, 177, 182, 183}

	avg := gbtFalgo.Average(l, h)
	fmt.Println(avg)

	// scheduler example start
	fn := func() {
		getOpenOrders(coin)
	}

	// run scheduler every 60 seconds
	scheduler := gbtScheduler.Schedule(fn, 60*time.Second)

	// stop by pressing enter
	var input string
	fmt.Scanln(&input)
	scheduler <- true

	// stop by using time.Sleep, after an hour
	//time.Sleep(3600 * time.Second)
	//scheduler <- true
	fmt.Println("Done")
	// scheduler example end

	status := selectBittrexEndpoint(arg, coin)
	if !status {
		fmt.Println("Ëndpoint doesn't exist, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getMarkets")
		fmt.Println("./GoBittrex geCurrencies")
		fmt.Println("./GoBittrex getTicks <coin_symbol>")
		fmt.Println("./GoBittrex getOpenOrders <coin_symbol>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func selectBittrexEndpoint(param string, coin string) bool {
	status := false
	switch param {
	case "getMarkets":
		getMarkets()
		status = true
	case "getCurrencies":
		getCurrencies()
		status = true
	case "getTicks":
		getTickes(coin)
		status = true
	case "getOpenOrders":
		getOpenOrders(coin)
		status = true
	default:
		status = false
	}
	return status
}

func getMarkets() {
	//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getmarkets", false)
	//URL := API_PATH + "/public/getmarkets"
	URL := fmt.Sprintf("%s/public/getmarkets", gbtConfig.API_PATH)
	gbtHttp.GetMarketsDataJson(URL, false)
}

func getCurrencies() {
	//gbtHttp.GetData("https://bittrex.com/api/v1.1/public/getcurrencies", false)
	//URL := API_PATH + "/public/getcurrencies"
	URL := fmt.Sprintf("%s/public/getcurrencies", gbtConfig.API_PATH)
	gbtHttp.GetCurrenciesDataJson(URL, false)
}

func getTickes(coin string) {
	//gbtHttp.GetData(gbtConfig.API_PATH_V2+"/pub/market/GetTicks?marketName=BTC-"+coin+"&tickInterval=day&_="+gbtUtil.FormatInt(gbtUtil.GetTimestamp()), false)
	//URL := gbtConfig.API_PATH_V2 + "/pub/market/GetTicks?marketName=BTC-" + coin + "&tickInterval=day&_=" + gbtUtil.FormatInt(gbtUtil.GetTimestamp())
	URL := fmt.Sprintf("%s/pub/market/GetTicks?marketName=BTC-%s&tickInterval=day&_=%s", gbtConfig.API_PATH_V2, coin, gbtUtil.FormatInt(gbtUtil.GetTimestamp()))
	gbtHttp.GetTicksDataJson(URL, false)
}

func getOpenOrders(coin string) {
	URL := fmt.Sprintf("%s/market/getopenorders?apikey=%s&market=BTC-%s", gbtConfig.API_PATH, gbtConfig.API_KEY, coin)
	gbtHttp.GetData(URL, true)
	gbtHttp.GetOpenOrdersDataJson(URL, true)
}
