package main

import (
    "fmt"
    btrexhttp "GoBittrex/http"
)

func main() {
    fmt.Printf("Hello Bittrex\n")

    // btrexhttp.ExportedMethod()
    // btrexhttp.GetGolangData()
    btrexhttp.GetGolangDataJson()
}
