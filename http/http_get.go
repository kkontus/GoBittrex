package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"bytes"
	gbtEntity "GoBittrex/entity"
)

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer // or out := bytes.Buffer{}
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func showError(err error) {
	fmt.Printf("%s", "Error message: ")
	fmt.Printf("%s\n", err)
	os.Exit(1)
}

func jsonDecode(url string) (error, *json.Decoder, *http.Response) {
	response, err := http.Get(url)
	if err != nil {
		showError(err)
	}

	decoder := json.NewDecoder(response.Body)
	return err, decoder, response
}

func GetData(url string) {
	response, err := http.Get(url)
	if err != nil {
		showError(err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			showError(err)
		}

		fmt.Printf("%s\n", JsonPrettyPrint(string(contents)))
	}
}

func GetMarketsDataJson(url string) {
	err, decoder, response := jsonDecode(url)
	defer response.Body.Close()
	getMarkets(err, decoder)
}

func GetCurrenciesDataJson(url string) {
	err, decoder, response := jsonDecode(url)
	defer response.Body.Close()
	getCurrencies(err, decoder)
}

func getMarkets(err error, decoder *json.Decoder) {
	responseData := gbtEntity.MarketsResponse{} // or var responseData = gbtEntity.MarketsResponse{}
	err = decoder.Decode(&responseData)
	if err != nil {
		showError(err)
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
		showError(err)
	} else {
		//fmt.Printf("Result: %v\n", responseData.Result)
		result := responseData.Result
		for i, elem := range result {
			fmt.Printf("%d: %s %f\n", i, elem.Currency, elem.TxFee)
		}
	}
}
