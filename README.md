# GoBittrex
Bittrex API checker in Go. 
This wrapper uses both v1.1 and v2.0 to fetch data from Bittrex.

#### Available commands:

```bash
Download or clone project from Github
run:
$ go get -u golang.org/x/oauth2/google
$ go build
$ ./GoBittrex getMarkets
$ ./GoBittrex getCurrencies
$ ./GoBittrex getTicks <coin_symbol>
$ ./GoBittrex getOpenOrders <coin_symbol>
$ ./GoBittrex getOrderBook <coin_symbol>
$ ./GoBittrex getCoinInfo <coin_symbol>
$ ./GoBittrex getCoinsInfo <convert> <start> <limit>
$ ./GoBittrex getCmcalCoins
$ ./GoBittrex getCmcalCategories
$ ./GoBittrex runServer"
$ ./GoBittrex sendPush
```

```bash
Clean dependencies
run:
$ go clean -i -x cloud.google.com/...
$ go clean -i -x golang.org/...
```