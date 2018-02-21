package main

import (
	"fmt"
	"os"
	"flag"
	"errors"
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

	cmdPtr := parseFlags()

	if gbtUtil.Contains(routesBittrex, *cmdPtr) {
		pingRoutesBittrex()
	} else if gbtUtil.Contains(routesCmc, *cmdPtr) {
		pingRoutesCmc()
	} else if gbtUtil.Contains(routesCmcal, *cmdPtr) {
		pingRoutesCmcal()
	} else if gbtUtil.Contains(routesFirebase, *cmdPtr) {
		pingRoutesFirebase()
	} else if gbtUtil.Contains(routesGeneral, *cmdPtr) {
		pingRoutesGeneral()
	} else {
		fmt.Println("Route Slice does not contain the element, check if route is added to config", flag.Arg(0))
	}
}

func parseFlags() *string {
	cmdPtr := flag.String("cmd", "", "command")
	flag.Parse()

	fmt.Println("cmd: ", *cmdPtr)
	fmt.Println(flag.Args())

	if (*cmdPtr == "") {
		err := errors.New("command parameter must be set")
		fmt.Println(err)
		os.Exit(1)
	}

	return cmdPtr
}

func pingRoutesBittrex() {
	cmd, params, err := gbtValidator.ValidateParams(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectBittrexRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=getMarkets")
		fmt.Println("./GoBittrex --cmd=geCurrencies")
		fmt.Println("./GoBittrex --cmd=getTicks <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=getOpenOrders <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=startTrailing <coin_symbol> SL <sell_loss> TP <take_profit> TTP <trailing_take_profit>")
		fmt.Println("./GoBittrex --cmd=getOrderBook <coin_symbol>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesCmc() {
	cmd, params, err := gbtValidator.ValidateParams(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectCoinmarketcapRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=getCoinInfo <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=getCoinsInfo <convert> <start> <limit>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesCmcal() {
	cmd, params, err := gbtValidator.ValidateParams(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectCoinmarketcalRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=getCmcalCoins")
		fmt.Println("./GoBittrex --cmd=getCmcalCategories")
		fmt.Println("./GoBittrex --cmd=getCmcalEvents")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesFirebase() {
	cmd, params, err := gbtValidator.ValidateParams(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectFirebaseRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=sendPush")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesGeneral() {
	cmd, params, err := gbtValidator.ValidateParams(flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectGeneralRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=runServer")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}
