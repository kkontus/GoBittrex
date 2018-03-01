package routes

import (
	"fmt"
	"net/url"
	gbtHttp "GoBittrex/http"
	gbtConfig "GoBittrex/config"
	gbtError "GoBittrex/error"
	gbtValidator "GoBittrex/validator"
	gbtEntity "GoBittrex/entity/cmcal"
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
		switch v := args.(type) {
		case gbtValidator.GetEventsParams:
			getCmcalEvents(v)
			status = true
		default:
			status = false
		}
	default:
		status = false
	}
	return status
}

func getCmcalCoins() {
	URL := fmt.Sprintf("%s/coins", gbtConfig.API_PATH_COINMARKETCAL)
	resp, err := gbtHttp.GetCmcalCoins(URL, gbtConfig.NONE)

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
	resp, err := gbtHttp.GetCmcalCategories(URL, gbtConfig.NONE)

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

func getCmcalEvents(v gbtValidator.GetEventsParams) {
	URL := fmt.Sprintf("%s/events?page=%s&max=%s&dateRangeStart=%s&dateRangeEnd=%s&coins=%s&categories=%s&sortBy=%s&showOnly=%s",
		gbtConfig.API_PATH_COINMARKETCAL,
		url.QueryEscape(v.Start),
		url.QueryEscape(v.Limit),
		url.QueryEscape(v.DateStart),
		url.QueryEscape(v.DateEnd),
		url.QueryEscape(v.Coins),
		url.QueryEscape(v.Categories),
		url.QueryEscape(v.SortBy),
		url.QueryEscape(v.ShowOnly))

	resp, err := gbtHttp.GetCmcalEvents(URL, gbtConfig.NONE)

	if err != nil {
		gbtError.ShowError(err)
	} else {
		switch v := resp.(type) {
		case []gbtEntity.CmcalEvents:
			for i, elem := range v {
				fmt.Printf("%d: %s %s %s %s\n", i, elem.DateEvent, elem.Title, elem.Coins[0].Name, elem.Coins[0].Symbol)
			}
		default:
			fmt.Println("Unsupported value in getCmcalEvents")
		}
	}
}
