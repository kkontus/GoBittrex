package talgo

import (
	"errors"
	"fmt"
	"strings"
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
	if len(args) == 3 {
		var s interface{}
		switch strings.ToLower(strategy) {
		case S_FIBB:
			s = Fibb{Coin: args[0], Exchange: exchange}
		case S_SMA:
			s = SMA{Coin: args[0], Exchange: exchange}
		case S_EMA5:
			s = EMA{Coin: args[0], Exchange: exchange, Crossover: 5}
		case S_EMA20:
			s = EMA{Coin: args[0], Exchange: exchange, Crossover: 20}
		default:
			return cmd, nil, errors.New("Trading strategy not found")
		}

		return cmd, s, nil
	} else {
		return cmd, nil, errors.New(fmt.Sprintf("%s unsupported parameters", cmd))
	}
}
