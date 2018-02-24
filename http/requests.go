package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	gbtUtil "GoBittrex/util"
	gbtError "GoBittrex/error"
	gbtEntity "GoBittrex/entity"
	. "GoBittrex/config"
)

func jsonDecode(url string, authRequired AuthType) (error, *json.Decoder, *http.Response) {
	//resp, err := http.Get(url)
	resp, err := NewRequest(http.MethodGet, url, nil, authRequired)
	if err != nil {
		gbtError.ShowError(err)
	}

	decoder := json.NewDecoder(resp.Body)
	return err, decoder, resp
}

func GetData(url string, authRequired AuthType) {
	// resp, err := http.Get(url)
	resp, err := NewRequest(http.MethodGet, url, nil, authRequired)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			gbtError.ShowError(err)
		}

		fmt.Printf("%s\n", gbtUtil.JsonPrettyPrint(string(contents)))
	}
}

// BITTREX
func GetMarkets(url string, authRequired AuthType) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	responseData := gbtEntity.MarketsResponse{} // or var responseData = gbtEntity.MarketsResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result
		for i, elem := range result {
			fmt.Printf("%d: %s %s\n", i, elem.MarketCurrency, elem.MarketName)
		}
	}
}

// BITTREX
func GetCurrencies(url string, authRequired AuthType) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	responseData := gbtEntity.CurrenciesResponse{} // or var responseData = gbtEntity.CurrenciesResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result
		for i, elem := range result {
			fmt.Printf("%d: %s %f\n", i, elem.Currency, elem.TxFee)
		}
	}
}

// BITTREX
func GetTicks(url string, authRequired AuthType) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	responseData := gbtEntity.TicksResponse{} // or var responseData = gbtEntity.TicksResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result
		for i, elem := range result {
			fmt.Printf("%d: %f %f\n", i, elem.Low, elem.High)
		}
	}
}

// BITTREX
func GetOpenOrders(url string, authRequired AuthType) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	responseData := gbtEntity.OpenOrdersResponse{} // or var responseData = gbtEntity.OpenOrdersResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result
		for i, elem := range result {
			fmt.Printf("%d: %s %s %f %f\n", i, elem.Exchange, elem.OrderType, elem.Quantity, elem.Limit)
		}
	}
}

// BITTREX
func GetOrderBook(url string, authRequired AuthType) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	responseData := gbtEntity.OrderBookResponse{} // or var responseData = gbtEntity.OpenOrdersResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		gbtError.ShowError(err)
	} else {
		fmt.Println("BUY")

		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result

		resultBuy := result.Buy
		for i, elem := range resultBuy {
			fmt.Printf("%d: %f %f\n", i, elem.Rate, elem.Quantity)
		}

		fmt.Println("SELL")

		resultSell := result.Sell
		for i, elem := range resultSell {
			fmt.Printf("%d: %f %f\n", i, elem.Rate, elem.Quantity)
		}
	}
}

// COINMARKETCAP
func GetCmcInfo(url string, authRequired AuthType) (info interface{}, err error) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	var responseData []gbtEntity.CmcCoinInfo
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	} else {
		return responseData, nil
	}
}

// COINMARKETCAL
func GetCmcalCoins(url string, authRequired AuthType) (info interface{}, err error) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	var responseData []string
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	} else {
		return responseData, nil
	}
}

// COINMARKETCAL
func GetCmcalCategories(url string, authRequired AuthType) (info interface{}, err error) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	var responseData []string
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	} else {
		return responseData, nil
	}
}

// COINMARKETCAL
func GetCmcalEvents(url string, authRequired AuthType) (info interface{}, err error) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()

	var responseData []gbtEntity.CmcalEvents
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	} else {
		return responseData, nil
	}
}
