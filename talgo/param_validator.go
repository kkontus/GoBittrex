package talgo

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
	gbtError "GoBittrex/error"
	gbtString "GoBittrex/c_string"
)

func ValidateParams(cmd string, strategy string, exchange string, args []string) (command string, params interface{}, err error) {
	// cmd := flag.Lookup("cmd").Value.(flag.Getter).Get().(string)

	switch cmd {
	case "runTradingBot":
		return validateRunTradingBotParams(cmd, strategy, exchange, args)
	default:
		return cmd, nil, errors.New("command not recognized")
	}

	return cmd, params, err
}

func validateRunTradingBotParams(cmd string, strategy string, exchange string, args []string) (command string, params interface{}, err error) {
	var s interface{}
	switch strings.ToUpper(strategy) {
	case strings.ToUpper(I_FIB):
		if len(args) == 3 {
			s = Fib{Coin: args[0], Exchange: exchange}
		} else {
			return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
		}
	case strings.ToUpper(I_SMA):
		if len(args) == 3 {
			s = SMA{Coin: args[0], Exchange: exchange}
		} else {
			return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
		}
	case strings.ToUpper(I_EMA):
		if len(args) == 4 {
			period, err := strconv.Atoi(args[3])
			if err != nil {
				gbtError.ShowError(errors.New(gbtString.ERR_MSG_EMA_PERIOD_INT))
			}
			s = EMA{Coin: args[0], Exchange: exchange, Crossover: period}
		} else {
			return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
		}
	default:
		return cmd, nil, errors.New(gbtString.ERR_MSG_TRADING_STRATEGY_NOT_FOUND)
	}

	return cmd, s, nil
}
