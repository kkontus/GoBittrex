package config

var RoutesBittrex = []string{"getMarkets", "getCurrencies", "getTicks", "getOpenOrders", "getOrderBook", "startTrailing", "getMarketSummaries", "getMarketSummary"}
var RoutesCmc = []string{"getCoinInfo", "getCoinsInfo"}
var RoutesCmcal = []string{"getCmcalCoins", "getCmcalCategories", "getCmcalEvents"}
var RoutesFirebase = []string{"sendPush", "startRealtimeDatabase"}
var RoutesGeneral = []string{"runServer", "runTradingBot"}
