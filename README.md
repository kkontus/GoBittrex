# GoBittrex
Bittrex API checker in Go. 
This wrapper uses both v1.1 and v2.0 to fetch data from Bittrex.

#### Available commands:

```bash
Download or clone project from Github
run:
$ go get -u golang.org/x/oauth2/google
$ go build
$ ./GoBittrex --cmd=getMarkets
$ ./GoBittrex --cmd=getCurrencies
$ ./GoBittrex --cmd=getTicks <coin_symbol>
$ ./GoBittrex --cmd=getOpenOrders <coin_symbol>
$ ./GoBittrex --cmd=getOrderBook <coin_symbol>
$ ./GoBittrex --cmd=getCoinInfo <coin_symbol>
$ ./GoBittrex --cmd=getCoinsInfo <convert> <start> <limit>
$ ./GoBittrex --cmd=getCmcalCoins
$ ./GoBittrex --cmd=getCmcalCategories
$ ./GoBittrex --cmd=runServer
$ ./GoBittrex --cmd=sendPush // CryptocurrencyManager mobile app FCM token has to be supplied
```

```bash
Examples with args
run:
$ ./GoBittrex --cmd=getTicks NEO
$ ./GoBittrex --cmd=getOpenOrders NEO
$ ./GoBittrex --cmd=getOrderBook NEO
$ ./GoBittrex --cmd=getCoinInfo bitcoin
$ ./GoBittrex --cmd=getCoinsInfo USD 0 5
```

```bash
Clean dependencies
run:
$ go clean -i -x cloud.google.com/...
$ go clean -i -x golang.org/...
```