package logging

import (
	"g-case-study/consts"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func SetupLogger(environment consts.Environment, logLevel string) {
	if logLevel == "" {
		logLevel = "default"
	}

	level, err := setLogLevel(environment, logLevel)
	if err != nil {
		//Log.WithFields(logrus.Fields{"logLevel": logLevel}).Error(errors.Wrap(err, "SetupLogger: "))
		panic("Log level set error, paniced!")
	}

	Log = getLogger(*level)
	Log.Debug("Here is loaded configuration")

	Log.Infof("Running on %v environment", environment)
}
func setLogLevel(environment consts.Environment, logLevel string) (*logrus.Level, error) {

	if logLevel == "default" {
		defaultLogLevel := consts.LogTypeEnvMap[environment]
		return &defaultLogLevel, nil
	}

	level, err := logrus.ParseLevel(logLevel)

	if err != nil {
		return nil, errors.Wrap(err, "setLogLevel: ")
	}

	return &level, nil
}
func getLogger(consoleLogLevel logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(consoleLogLevel)
	logger.SetOutput(os.Stdout)

	return logger
}
