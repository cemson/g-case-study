package utilities

import (
	"fmt"
	"g-case-study/consts"
	log "g-case-study/logging"
	"g-case-study/settings"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

func LoadConfig(incomingConfigIns *settings.AppSettings) {
	configFileEnv := consts.Dev
	argEnv, err := GetEnvironment()
	if err != nil {
		panic(errors.New(fmt.Sprintf("Invalid environment argument: %v\n%v", argEnv, err.Error())))
	}

	configFileEnv = argEnv

	setViperConfig(configFileEnv)
	setFields(incomingConfigIns)

	log.SetupLogger(configFileEnv, incomingConfigIns.LogLevel)
}

func GetEnvironment() (consts.Environment, error) {
	var env string
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		env = argsWithoutProg[0]
	}

	return consts.ParseEnv(env)
}
func setViperConfig(configFileEnv consts.Environment) {
	viper.SetConfigName("conf." + string(configFileEnv))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".") // look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrap(err, "setViperConfig: "))
	}
}
func setFields(settings interface{}) {
	valueDef := reflect.ValueOf(settings)
	typeDef := reflect.TypeOf(settings)

	if typeDef.Kind() == reflect.Ptr {
		typeDef = typeDef.Elem()
		valueDef = valueDef.Elem()
	}

	for i := 0; i < typeDef.NumField(); i++ {
		v := valueDef.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			setFields(v.Addr().Interface())
		default:
			fieldMeta := valueDef.Type().Field(i)
			fieldNameLowerCase := strings.ToLower(fieldMeta.Name)
			configValue := viper.GetString(fieldNameLowerCase)
			if configValue == "" {
				v.SetString(fieldMeta.Tag.Get("default"))
			} else {
				v.SetString(configValue)
			}
		}
	}
}
