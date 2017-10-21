package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "encoding/json"
    "bytes"
)


type Response struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Result []BittrexResponse `json:"result"`
}

type BittrexResponse struct {
    MarketCurrency string `json:"MarketCurrency"`
    BaseCurrency string `json:"BaseCurrency"`
    MarketCurrencyLong string `json:"MarketCurrencyLong"`
    BaseCurrencyLong string `json:"BaseCurrencyLong"`
    MinTradeSize float32 `json:"MinTradeSize"`
    MarketName string `json:"MarketName"`
    IsActive bool `json:"IsActive"`
    Created string `json:"Created"`
    Notice string `json:"Notice"`
    IsSponsored bool `json:"IsSponsored"`
    LogoUrl string `json:"LogoUrl"`
}

func JsonPrettyPrint( in string) string {
    var out bytes.Buffer // or out := bytes.Buffer{}
    err := json.Indent(&out, []byte(in), "", "\t")
    if err != nil {
        return in
    }
    return out.String()
}

func ExportedMethod() {
    fmt.Println("Bittrex")
}

func GetGolangData() {
    response, err := http.Get("https://bittrex.com/api/v1.1/public/getmarkets")
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

func GetGolangDataJson() {
    response, err := http.Get("https://bittrex.com/api/v1.1/public/getmarkets")
    if err != nil {
        fmt.Printf("%s", "Error message: ")
        fmt.Printf("%s\n", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()

        responseData := Response{} // or var responseData = Response{}

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
