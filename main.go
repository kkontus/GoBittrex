package main

import (
	"fmt"
	"os"
	"flag"
	"errors"
	gbtRoutes "GoBittrex/routes"
	gbtValidator "GoBittrex/validator"
	gbtTalgoValidator "GoBittrex/talgo"
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
		pingRoutesBittrex(*cmdPtr)
	} else if gbtUtil.Contains(routesCmc, *cmdPtr) {
		pingRoutesCmc(*cmdPtr)
	} else if gbtUtil.Contains(routesCmcal, *cmdPtr) {
		pingRoutesCmcal(*cmdPtr)
	} else if gbtUtil.Contains(routesFirebase, *cmdPtr) {
		pingRoutesFirebase(*cmdPtr)
	} else if gbtUtil.Contains(routesGeneral, *cmdPtr) {
		pingRoutesGeneral(*cmdPtr)
	} else {
		fmt.Println("Route Slice does not contain the element, check if route is added to config", flag.Arg(0))
	}
}

func parseFlags() *string {
	// define flags
	cmdPtr := flag.String(gbtConfig.CMD, "", "command")
	flag.String(gbtConfig.STRATEGY, "", "strategy")
	flag.String(gbtConfig.EXCH, "", "exchange")

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

func pingRoutesBittrex(cmd string) {
	cmd, params, err := gbtValidator.ValidateParams(cmd, flag.Args())
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
		fmt.Println("./GoBittrex --cmd=getMarketSummaries")
		fmt.Println("./GoBittrex --cmd=getMarketSummary <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=getTicks <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=getOpenOrders <coin_symbol>")
		fmt.Println("./GoBittrex --cmd=startTrailing <coin_symbol> SL <sell_loss> TP <take_profit> TTP <trailing_take_profit>")
		fmt.Println("./GoBittrex --cmd=getOrderBook <coin_symbol>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesCmc(cmd string) {
	cmd, params, err := gbtValidator.ValidateParams(cmd, flag.Args())
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

func pingRoutesCmcal(cmd string) {
	cmd, params, err := gbtValidator.ValidateParams(cmd, flag.Args())
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

func pingRoutesFirebase(cmd string) {
	cmd, params, err := gbtValidator.ValidateParams(cmd, flag.Args())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectFirebaseRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=sendPush")
		fmt.Println("./GoBittrex --cmd=startRealtimeDatabase")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}

func pingRoutesGeneral(cmd string) {
	var params interface{}
	var err error
	if cmd == "runTradingBot" {
		strategy := flag.Lookup(gbtConfig.STRATEGY).Value.(flag.Getter).Get().(string)
		exch := flag.Lookup(gbtConfig.EXCH).Value.(flag.Getter).Get().(string)
		cmd, params, err = gbtTalgoValidator.ValidateParams(cmd, strategy, exch, flag.Args())
	} else {
		cmd, params, err = gbtValidator.ValidateParams(cmd, flag.Args())
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := gbtRoutes.SelectGeneralRoute(cmd, params)
	if !status {
		fmt.Println("Unable to run command, check if command is valid:")
		fmt.Println("")
		fmt.Println("./GoBittrex --cmd=runServer")
		fmt.Println("./GoBittrex --cmd=runTradingBot --strategy=<strategy> --exch=<exchange> <coin_symbol> <timespan> <pips>")
		fmt.Println("")
	} else {
		fmt.Println("Status: OK'")
	}
}
