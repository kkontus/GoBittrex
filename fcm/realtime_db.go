package fcm

import (
	"fmt"
	"net/http"
	"io/ioutil"
	gbtConfig "GoBittrex/config"
	gbtError "GoBittrex/error"
	gbtUtil "GoBittrex/util"
)

func RealtimeDB() {
	URL := fmt.Sprintf("%s", gbtConfig.FCM_REALTIME_CREDS)
	resp, err := NewRequestFcm(http.MethodGet, URL, nil)

	handleResponseRealtimeDB(err, resp)

	fetchCurrencies(URL, resp, err)
}

func handleResponseRealtimeDB(err error, resp *http.Response) {
	if err != nil {
		gbtError.ShowError(err)
	} else {
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			gbtError.ShowError(err)
		}

		fmt.Printf("%s\n", gbtUtil.JsonPrettyPrint(string(contents)))
	}
}

func fetchCurrencies(URL string, resp *http.Response, err error) {
	URL = "https://crypto-curr-exch.firebaseio.com/currencies.json"
	resp, err = NewRequestFcm(http.MethodGet, URL, nil)
	handleResponseRealtimeDB(err, resp)
}
