package validator

import (
	"errors"
	"strconv"
	"fmt"
)

func ValidateParams(args []string) (command string, params interface{}, err error) {
	cmd := ""
	if len(args) > 1 {
		cmd = args[1]
	} else {
		return cmd, nil, errors.New("command parameter must be set")
	}

	switch cmd {
	case "getMarkets":
		return validateGetMarketsParams(args)
	case "getCurrencies":
		return validateGetCurrenciesParams(args)
	case "getTicks":
		return validateGetTicksParams(args)
	case "getOpenOrders":
		return validateGetOpenOrdersParams(args)
	case "startTrailing":
		return validateStartTrailingParams(args)
	case "sendPush":
		return validateSendPushParams(args)
	default:
		return cmd, nil, errors.New("command not recognized")
	}

	return cmd, params, err
}

func validateGetMarketsParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]
	if len(args) == 2 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetCurrenciesParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]
	if len(args) == 2 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetTicksParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]
	if len(args) == 3 {
		s := GetTicksParams{Coin: args[2]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateGetOpenOrdersParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]
	if len(args) == 3 {
		s := GetOpenOrdersParams{Coin: args[2]}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateStartTrailingParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]

	if len(args) == 6 {
		valueSL, errSL := strconv.ParseFloat(args[3], 32)
		if errSL != nil {
			return cmd, nil, errSL
		}
		SL := float32(valueSL)

		valueTP, errTP := strconv.ParseFloat(args[4], 32)
		if errTP != nil {
			return cmd, nil, errTP
		}
		TP := float32(valueTP)

		valueTTP, errTTP := strconv.ParseFloat(args[5], 32)
		if errTTP != nil {
			return cmd, nil, errTTP
		}
		TTP := float32(valueTTP)

		s := StartTrailingParams{Coin: args[2], SL: SL, TP: TP, TTP: TTP}
		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}

func validateSendPushParams(args []string) (command string, params interface{}, err error) {
	cmd := args[1]
	if len(args) == 2 {
		return cmd, nil, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}
