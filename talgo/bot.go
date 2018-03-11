package talgo

import (
	"fmt"
	"time"
	gbtScheduler "GoBittrex/scheduler"
)

const (
	S_FIBB  string = "fibb"
	S_SMA   string = "sma"
	S_EMA5  string = "ema5"
	S_EMA20 string = "ema20"
)

func RunTradingBot(args interface{}) {
	fmt.Println("Trading bot started")

	switch v := args.(type) {
	case Fibb:
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

func runFibb(fibb Fibb) {
	fn := TradeFibb(fibb)
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

	if ema.Crossover == 5 {
		fn = TradeEMA5(ema)
		startTrading(fn)
	} else if ema.Crossover == 20 {
		fn = TradeEMA20(ema)
	}

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
