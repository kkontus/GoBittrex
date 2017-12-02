package tasks

import (
	"fmt"
	"time"
	gbtHttp "GoBittrex/http"
	gbtUtil "GoBittrex/util"
	gbtConfig "GoBittrex/config"
	gbtValidator "GoBittrex/validator"
	gbtScheduler "GoBittrex/scheduler"
	gbtFcm "GoBittrex/fcm"
)

func SelectBittrexEndpoint(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "getMarkets":
		getMarkets()
		status = true
	case "getCurrencies":
		getCurrencies()
		status = true
	case "getTicks":
		switch v := args.(type) {
		case gbtValidator.GetTicksParams:
			getTickes(v.Coin)
			status = true
		default:
			status = false
		}
	case "getOpenOrders":
		switch v := args.(type) {
		case gbtValidator.GetOpenOrdersParams:
			getOpenOrders(v.Coin)
			status = true
		default:
			status = false
		}
	case "startTrailing":
		switch v := args.(type) {
		case gbtValidator.StartTrailingParams:
			startTrailing(v.Coin, v.SL, v.TP, v.TTP)
			status = true
		default:
			status = false
		}
	case "sendPush":
		sendPush()
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

func startTrailing(coin string, SL float32, TP float32, TTP float32) {
	runStartTrailingScheduler(coin, SL, TP, TTP)
}

func runStartTrailingScheduler(coin string, SL float32, TP float32, TTP float32) {
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
	fmt.Println("Exiting trailing")
}

func sendPush() {
	gbtFcm.SendPush(false) // refactor this to send message instead hardcoded data
}
