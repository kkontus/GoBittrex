package main

import (
	"fmt"
	gbtHttp "GoBittrex/http"
)

func main() {
	fmt.Printf("Hello Bittrex\n")

	//gbtHttp.GetGolangData("https://bittrex.com/api/v1.1/public/getmarkets")
	gbtHttp.GetGolangDataJson("https://bittrex.com/api/v1.1/public/getmarkets")
}
