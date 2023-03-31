package consts

import (
	"errors"
	"strings"
)

type Environment string

const (
	Dev    Environment = "dev"
	Stg                = "stg"
	ProdEu             = "prod-eu"
	ProdUs             = "prod"
	Local              = "local"
	Test               = "test"
)

func ParseEnv(env string) (Environment, error) {
	switch strings.ToLower(env) {
	case "dev":
		return Dev, nil
	case "stg":
		return Stg, nil
	case "prod-eu":
		return ProdEu, nil
	case "prod":
		return ProdUs, nil
	case "local":
		return Local, nil
	case "-test":
		return Test, nil
	}

	var e Environment
	return e, errors.New("ParseEnv: not a valid environment " + env)
}
