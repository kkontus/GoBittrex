package validator

import (
	"errors"
	"strconv"
	"fmt"
)

func ValidateParams(cmd string, args []string) (command string, params interface{}, err error) {
	// cmd := flag.Lookup("cmd").Value.(flag.Getter).Get().(string)

	switch cmd {
	case "getMarkets":
		return validateGetMarketsParams(cmd, args)
	case "getCurrencies":
		return validateGetCurrenciesParams(cmd, args)
	case "getMarketSummaries":
		return validateGetMarketSummariesParams(cmd, args)
	case "getMarketSummary":
		return validateGetMarketSummaryParams(cmd, args)
	case "getTicks":
		return validateGetTicksParams(cmd, args)
	case "getOpenOrders":
		return validateGetOpenOrdersParams(cmd, args)
	case "getOrderBook":
		return validateGetOrderBookParams(cmd, args)
	case "startTrailing":
		return validateStartTrailingParams(cmd, args)
	case "runServer":
		return validateRunServerParams(cmd, args)
	case "sendPush":
		return validateSendPushParams(cmd, args)
	case "startRealtimeDatabase":
		return validateStartRealtimeDatabaseParams(cmd, args)
	case "getCoinInfo":
		return validateGetCoinInfoParams(cmd, args)
	case "getCoinsInfo":
		return validateGetCoinsInfoParams(cmd, args)
	case "getCmcalCoins":
		return validateGetCmcalCoinsParams(cmd, args)
	case "getCmcalCategories":
		return validateGetCmcalCategoriesParams(cmd, args)
	case "getCmcalEvents":
		return validateGetCmcalEventsParams(cmd, args)

	default:
		return cmd, nil, errors.New("command not recognized")
	}

	return cmd, params, err
}

func validateGetMarketsParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCurrenciesParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetMarketSummariesParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetMarketSummaryParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 1 {
		s := GetMarketSummaryParams{Coin: args[0]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetTicksParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 1 {
		s := GetTicksParams{Coin: args[0]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetOpenOrdersParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 1 {
		s := GetOpenOrdersParams{Coin: args[0]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetOrderBookParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 1 {
		s := GetOrderBookParams{Coin: args[0]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateStartTrailingParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 7 {
		valueSL, errSL := strconv.ParseFloat(args[2], 32)
		if errSL != nil {
			return cmd, nil, errSL
		}
		SL := float32(valueSL)

		valueTP, errTP := strconv.ParseFloat(args[4], 32)
		if errTP != nil {
			return cmd, nil, errTP
		}
		TP := float32(valueTP)

		valueTTP, errTTP := strconv.ParseFloat(args[6], 32)
		if errTTP != nil {
			return cmd, nil, errTTP
		}
		TTP := float32(valueTTP)

		s := StartTrailingParams{Coin: args[0], SL: SL, TP: TP, TTP: TTP}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateRunServerParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateSendPushParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateStartRealtimeDatabaseParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCoinInfoParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 1 {
		s := GetCoinInfoParams{Coin: args[0]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCoinsInfoParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 3 {
		s := GetCoinsInfoParams{Convert: args[0], Start: args[1], Limit: args[2]}
		return cmd, s, nil
	} else if len(args) == 2 {
		s := GetCoinsInfoParams{Convert: args[0], Start: args[1], Limit: ""}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCmcalCoinsParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCmcalCategoriesParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 0 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCmcalEventsParams(cmd string, args []string) (command string, params interface{}, err error) {
	if len(args) == 8 {
		s := GetEventsParams{Start: args[0], Limit: args[1], DateStart: args[2], DateEnd: args[3], Coins: args[4], Categories: args[5], SortBy: args[6], ShowOnly: args[7]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}
