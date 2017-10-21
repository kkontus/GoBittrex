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

func GetGolangData(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", "Error message: ")
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", JsonPrettyPrint(string(contents)))
	}
}

func GetGolangDataJson(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", "Error message: ")
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()

		responseData := gbtEntity.Response{} // or var responseData = Response{}

		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			fmt.Printf("%s", "Error message: ")
			fmt.Printf("%s\n", err)
			os.Exit(1)
		} else {
			// fmt.Printf("Results: %v\n", responseData)
			// fmt.Printf("Results: %v\n", responseData.Success)
			// fmt.Printf("Results: %v\n", responseData.Message)
			// fmt.Printf("Results: %v\n", responseData.Result)

			result := responseData.Result
			// fmt.Printf("Results: %v\n", result)

			for i, elem := range result {
				fmt.Printf("%d: %s %s\n", i, elem.MarketCurrency, elem.MarketName)
			}
		}
	}
}
