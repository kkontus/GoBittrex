package routes

import (
	"fmt"
	gbtHttp "GoBittrex/http"
	gbtConfig "GoBittrex/config"
	gbtError "GoBittrex/error"
)

func SelectCoinmarketcalRoute(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "getCmcalCoins":
		getCmcalCoins()
		status = true
	case "getCmcalCategories":
		getCmcalCategories()
		status = true
	case "getCmcalEvents":
		getCmcalEvents()
		status = true
	default:
		status = false
	}
	return status
}

func getCmcalCoins() {
	URL := fmt.Sprintf("%s/coins", gbtConfig.API_PATH_COINMARKETCAL)
	resp, err := gbtHttp.GetCmcalCoins(URL, false)

	if err != nil {
		gbtError.ShowError(err)
	} else {
		switch v := resp.(type) {
		case []string:
			for i, elem := range v {
				fmt.Printf("%d: %s\n", i, elem)
			}
		default:
			fmt.Println("Unsupported value in getCmcalCoins")
		}
	}
}

func getCmcalCategories() {
	URL := fmt.Sprintf("%s/categories", gbtConfig.API_PATH_COINMARKETCAL)
	resp, err := gbtHttp.GetCmcalCategories(URL, false)

	if err != nil {
		gbtError.ShowError(err)
	} else {
		switch v := resp.(type) {
		case []string:
			for i, elem := range v {
				fmt.Printf("%d: %s\n", i, elem)
			}
		default:
			fmt.Println("Unsupported value in getCmcalCategories")
		}
	}
}

func getCmcalEvents() {
	// TODO need to implement
	fmt.Println("getCmcalEvents not implemented yet")
}
