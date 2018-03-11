package routes

import (
	gbtTrading "GoBittrex/talgo"
)

func SelectGeneralRoute(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "runServer":
		RunServer()
		status = true
	case "runTradingBot":
		gbtTrading.RunTradingBot(args)
		status = true
	default:
		status = false
	}
	return status
}
