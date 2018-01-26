package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	gbtUtil "GoBittrex/util"
	gbtError "GoBittrex/error"
	gbtEntity "GoBittrex/entity"
)

func jsonDecode(url string, authRequired bool) (error, *json.Decoder, *http.Response) {
	//resp, err := http.Get(url)
	resp, err := NewRequest(http.MethodGet, url, nil, authRequired)
	if err != nil {
		gbtError.ShowError(err)
	}

	decoder := json.NewDecoder(resp.Body)
	return err, decoder, resp
}

func GetData(url string, authRequired bool) {
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

func GetMarkets(url string, authRequired bool) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	getMarkets(err, decoder)
}

func GetCurrencies(url string, authRequired bool) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	getCurrencies(err, decoder)
}

func GetTicks(url string, authRequired bool) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	getTicksData(err, decoder)
}

func GetOpenOrders(url string, authRequired bool) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	getOpenOrdersData(err, decoder)
}

func GetOrderBook(url string, authRequired bool) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	getOrderBookData(err, decoder)
}

func GetCmcInfo(url string, authRequired bool) (info interface{}, err error) {
	err, decoder, resp := jsonDecode(url, authRequired)
	defer resp.Body.Close()
	return getCmcInfo(err, decoder)
}

func getCmcInfo(err error, decoder *json.Decoder) (info interface{}, error error) {
	var responseData []gbtEntity.CmcCoinInfo // or var responseData = gbtEntity.CmcCoinInfo{}
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	} else {
		return responseData, nil
	}
}

func getMarkets(err error, decoder *json.Decoder) {
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

func getCurrencies(err error, decoder *json.Decoder) {
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

func getTicksData(err error, decoder *json.Decoder) {
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

func getOpenOrdersData(err error, decoder *json.Decoder) {
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

func getOrderBookData(err error, decoder *json.Decoder) {
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
