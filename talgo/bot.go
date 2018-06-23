package talgo

import (
	"fmt"
	"time"
	gbtScheduler "GoBittrex/scheduler"
)

func RunTradingBot(args interface{}) {
	fmt.Println("Trading bot started")

	switch v := args.(type) {
	case Fib:
		runFibb(v)
	case SMA:
		runSma(v)
	case EMA:
		runEma(v)
	default:
		fmt.Println("Trading strategy not implemented")
	}

	fmt.Println("Trading bot stopped")
}

func runFibb(fib Fib) {
	fn := TradeFib(fib)
	startTrading(fn)
}

func runSma(sma SMA) {
	fn := TradeSMA(sma)
	startTrading(fn)
}

func runEma(ema EMA) {
	fn := func() {
		fmt.Println("Trading function should have replaced me")
	}

	fn = TradeEMA(ema)
	startTrading(fn)
}

func startTrading(fn func()) {
	// run scheduler every 60 seconds
	scheduler := gbtScheduler.Schedule(fn, 5*time.Second)
	// stop by pressing enter
	var input string
	fmt.Scanln(&input)
	scheduler <- true
	// stop by using time.Sleep, after an hour
	//time.Sleep(3600 * time.Second)
	//scheduler <- true
	fmt.Println("Exiting trailing")
}
