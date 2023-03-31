package consts

import "github.com/sirupsen/logrus"

var LogTypeEnvMap = map[Environment]logrus.Level{
	Stg:    logrus.DebugLevel,
	Local:  logrus.DebugLevel,
	Dev:    logrus.WarnLevel,
	ProdUs: logrus.ErrorLevel,
	ProdEu: logrus.ErrorLevel,
}
