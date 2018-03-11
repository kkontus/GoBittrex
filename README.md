# GoBittrex
Bittrex API checker in Go. 
This wrapper uses both v1.1 and v2.0 to fetch data from Bittrex.

#### Available commands:

```bash
Download or clone project from Github
run:
$ go get -u golang.org/x/oauth2/google
$ go get firebase.google.com/go
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
$ ./GoBittrex --cmd=getCmcalEvents <start> <limit> <date_start> <date_end> <'coin_name (coin_symbol),coin_name (coin_symbol)'> <'category1,category2'> <sort_by> <show_only>
$ ./GoBittrex --cmd=runServer
$ ./GoBittrex --cmd=runTradingBot --strategy=<strategy> --exch=<exchange> <coin_symbol> <timespan> <pips>
$ ./GoBittrex --cmd=sendPush // CryptocurrencyManager mobile app FCM token has to be supplied
$ ./GoBittrex --cmd=startRealtimeDatabase
```

```bash
Examples with args
run:
$ ./GoBittrex --cmd=getTicks NEO
$ ./GoBittrex --cmd=getOpenOrders NEO
$ ./GoBittrex --cmd=getOrderBook NEO
$ ./GoBittrex --cmd=getCoinInfo bitcoin
$ ./GoBittrex --cmd=getCoinsInfo USD 0 5
$ ./GoBittrex --cmd=runTradingBot --strategy=EMA20 --exch=Bittrex NEO 14400 40
$ ./GoBittrex --cmd=getCmcalEvents 1 1 23/02/2018 25/02/2018 'Phore (PHR)' Rebranding hot_events hot_events
$ ./GoBittrex --cmd=getCmcalEvents 1 3 23/02/2018 25/02/2018 'Phore (PHR),Worldcore (WRC)' 'Exchange,Rebranding' hot_events hot_events
```

```bash
Clean dependencies
run:
$ go clean -i -x cloud.google.com/...
$ go clean -i -x golang.org/...
$ go clean -i -x firebase.google.com/...
```