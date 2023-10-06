package loggerfactory

import (
	"corebanking/pkg/helper/config"

	"github.com/pkg/errors"
)

// package loggerfactory handles creating concrete logger with factory method pattern.

// logfactoryBuilderMap logger mapp to map logger code to logger builder.
//
//nolint:gochecknoglobals //Logger singleton builder instance
var logfactoryBuilderMap = map[string]LogFbInterface{
	config.ZAP: &ZapFactory{},
}

// LogFbInterface interface for logger factory.
type LogFbInterface interface {
	Build(*config.LogConfig) error
}

// GetLogFactoryBuilder accessors for factoryBuilderMap.
func GetLogFactoryBuilder(key string) LogFbInterface {
	return logfactoryBuilderMap[key]
}

// LoadLogger the logger.
func LoadLogger(lc config.LogConfig) error {
	loggerType := lc.Code
	err := GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
