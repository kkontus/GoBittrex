package routes

import (
	"fmt"
	gbtHttp "GoBittrex/http"
	gbtConfig "GoBittrex/config"
	gbtValidator "GoBittrex/validator"
	gbtEntity "GoBittrex/entity/cmc"
	gbtError "GoBittrex/error"
	gbtUtil "GoBittrex/util"
)

func SelectCoinmarketcapRoute(cmd string, args interface{}) bool {
	status := false
	switch cmd {
	case "getCoinInfo":
		switch v := args.(type) {
		case gbtValidator.GetCoinInfoParams:
			getCoinInfo(v.Coin)
			status = true
		default:
			status = false
		}
	case "getCoinsInfo":
		switch v := args.(type) {
		case gbtValidator.GetCoinsInfoParams:
			getCoinsInfo(v.Convert, v.Start, v.Limit)
			status = true
		default:
			status = false
		}
	default:
		status = false
	}
	return status
}

func getCoinInfo(coin string) {
	URL := fmt.Sprintf("%s/ticker/%s/", gbtConfig.API_PATH_COINMARKETCAP, coin)
	resp, err := gbtHttp.GetCmcInfo(URL, gbtConfig.NONE)

	if err != nil {
		gbtError.ShowError(err)
	} else {
		switch v := resp.(type) {
		case []gbtEntity.CmcCoinInfo:
			fmt.Printf("%s\n", v[0].Name+" "+v[0].MaxSupply)
		default:
			fmt.Println("Unsupported value in GetCmcInfo")
		}
	}
}

func getCoinsInfo(convert string, start string, limit string) {
	if convert == "" {
		convert = gbtConfig.DEFAULT_CURRENCY
	}

	if start == "" || !gbtUtil.IsNumeric(start) {
		start = "0"
	}

	URL := fmt.Sprintf("%s/ticker/?convert=%s&start=%s", gbtConfig.API_PATH_COINMARKETCAP, convert, start)
	if limit != "" && gbtUtil.IsNumeric(start) {
		URL = fmt.Sprintf("%s&limit=%s", URL, limit)
	}

	resp, err := gbtHttp.GetCmcInfo(URL, gbtConfig.NONE)

	if err != nil {
		gbtError.ShowError(err)
	} else {
		switch v := resp.(type) {
		case []gbtEntity.CmcCoinInfo:
			for i, elem := range v {
				fmt.Printf("%d: %s %s %s %s\n", i, elem.Name, elem.Symbol, elem.AvailableSupply, elem.MaxSupply)
			}
		default:
			fmt.Println("Unsupported value in GetCmcInfo")
		}
	}
}

func GetCoinInfoAPI(coin string) (info interface{}, err error) {
	URL := fmt.Sprintf("%s/ticker/%s/", gbtConfig.API_PATH_COINMARKETCAP, coin)
	return gbtHttp.GetCmcInfo(URL, gbtConfig.NONE)
}
