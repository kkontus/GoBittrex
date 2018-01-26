package routes

import (
	"fmt"
	gbtHttp "GoBittrex/http"
	gbtConfig "GoBittrex/config"
	gbtValidator "GoBittrex/validator"
	gbtEntity "GoBittrex/entity"
	gbtError "GoBittrex/error"
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
		status = true
	default:
		status = false
	}
	return status
}

func getCoinInfo(coin string) {
	URL := fmt.Sprintf("%s/ticker/%s/", gbtConfig.API_PATCH_COINMARKETCAP, coin)
	resp, err := gbtHttp.GetCmcInfo(URL, false)

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

func GetCoinInfoAPI(coin string) (info interface{}, err error) {
	URL := fmt.Sprintf("%s/ticker/%s/", gbtConfig.API_PATCH_COINMARKETCAP, coin)
	return gbtHttp.GetCmcInfo(URL, false)
}
