package main

import (
	"fmt"
	"os"
	gbtRoutes "GoBittrex/routes"
	gbtValidator "GoBittrex/validator"
	gbtUtil "GoBittrex/util"
	gbtConfig "GoBittrex/config"
)

func main() {
	fmt.Printf("Hello Bittrex\n")

	routesBittrex := gbtConfig.RoutesBittrex
	routesCmc := gbtConfig.RoutesCmc
	routesCmcal := gbtConfig.RoutesCmcal
	routesFirebase := gbtConfig.RoutesFirebase
	routesGeneral := gbtConfig.RoutesGeneral

	if gbtUtil.Contains(routesBittrex, os.Args[1]) {
		pingRoutesBittrex()
	} else if gbtUtil.Contains(routesCmc, os.Args[1]) {
		pingRoutesCmc()
	} else if gbtUtil.Contains(routesCmcal, os.Args[1]) {
		pingRoutesCmcal()
	} else if gbtUtil.Contains(routesFirebase, os.Args[1]) {
		pingRoutesFirebase()
	} else if gbtUtil.Contains(routesGeneral, os.Args[1]) {
		pingRoutesGeneral()
	} else {
		fmt.Println("Route Slice does not contain the element, check if route is added to config", os.Args[1])
	}
}

func pingRoutesBittrex() {
	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectBittrexRoute(cmd, params)
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

func pingRoutesCmc() {
	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectCoinmarketcapRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getCoinInfo <coin>")
		fmt.Println("./GoBittrex getCoinsInfo <convert> <start> <limit>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesCmcal() {
	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectCoinmarketcalRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex getCmcalCoins")
		fmt.Println("./GoBittrex getCmcalCategories")
		fmt.Println("./GoBittrex getCmcalEvents")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesFirebase() {
	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectFirebaseRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex sendPush")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesGeneral() {
	cmd, params, err := gbtValidator.ValidateParams(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectGeneralRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex runServer")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}
