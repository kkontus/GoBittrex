package routes

import (
	"fmt"
	"net/http"
	gbtEntity "GoBittrex/entity"
	gbtConfig "GoBittrex/config"
)

func RunServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/coin-info", handlerExample)

	http.ListenAndServe(":"+gbtConfig.CRYPTO_API_PORT, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from GoBittrex server!")
}

func handlerExample(w http.ResponseWriter, r *http.Request) {
	params, exists := r.URL.Query()["coin"]
	if !exists || len(params) < 1 {
		fmt.Println("Url Param 'coin' is missing")
		return
	}
	// Query()["coin"] will return an array of items, we only want the single item.
	coin := params[0]

	fmt.Println("Url Param 'coin' is: " + string(coin))

	resp, err := GetCoinInfoAPI(coin)
	if err != nil {
		fmt.Fprintf(w, "%v", "Error: "+err.Error())
	} else {
		switch v := resp.(type) {
		case []gbtEntity.CmcCoinInfo:
			fmt.Fprintf(w, v[0].Name+" "+v[0].MaxSupply)
		default:
			fmt.Fprintf(w, "%v", "Unsupported value in GetCmcInfo")
		}
	}
}
