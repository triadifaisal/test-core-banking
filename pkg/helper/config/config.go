package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/ini.v1"
)

var (
	config *ini.File

	onceLogConfig sync.Once
	logConfig     LogConfig
)

const (
	// EnvPrefix ...
	EnvPrefix      = "COREBANKING_"
	SwaggerSection = "swagger"
)

func init() {
	var err error

	configPath := os.Getenv("CORE_CONFIG")
	path := os.Getenv("CORE_DIR")
	if len(configPath) > 0 {
		path = configPath
	} else if len(path) > 0 {
		path += "/config.ini"
	} else {
		path = "config.ini"
	}

	config, err = ini.LoadSources(ini.LoadOptions{
		Insensitive:                true,
		AllowPythonMultilineValues: true,
	}, path)
	if err != nil {
		log.Panic("Configuration file could not be loaded")
	}

	populateEnv(config)
}

func populateEnv(config *ini.File) {
	path := os.Getenv("CORE_DIR")
	if len(path) > 0 {
		path += "/.env"
	} else {
		path = ".env"
	}
	_ = godotenv.Load(path)

	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			key := e[:i]
			value := e[i+1:]
			if strings.HasPrefix(key, EnvPrefix) {
				key = strings.TrimPrefix(key, EnvPrefix)
				config.Section("").Key(key).SetValue(value)
			}
		}
	}
}

// GetLoggerPath Get logger path
func GetLoggerPath() string {
	sec := config.Section("")
	if sec.HasKey("log_path") {
		return sec.Key("log_path").String()
	}

	return "/tmp/corebanking.log"
}

// GetLogConfig ...
func GetLogConfig() LogConfig {
	onceLogConfig.Do(func() {
		logConfig = LogConfig{
			Code:         config.Section("log").Key("CODE").String(),
			Level:        config.Section("log").Key("LEVEL").String(),
			Encoding:     config.Section("log").Key("ENCODING").String(),
			EnableCaller: config.Section("log").Key("ENABLE_CALLER").MustBool(false),
			FilePath:     config.Section("log").Key("FILE_PATH").String(),
		}
	})

	return logConfig
}

// IsProductionMode Get app mode.
func IsProductionMode() bool {
	if mode := config.Section("").Key("mode").String(); mode == "production" {
		return true
	}
	return false
}

// GetSwaggerServicePort Get Port Service config
func GetSwaggerServicePort(service string) int {
	return config.Section(SwaggerSection).Key(fmt.Sprintf("%v_SERVICE_PORT", strings.ToUpper(service))).MustInt(0)
}

// GetSwaggerServiceHost Get Host Service config
func GetSwaggerServiceHost(service string) string {
	host := config.Section(SwaggerSection).Key(fmt.Sprintf("%v_SERVICE_HOST", strings.ToUpper(service))).String()
	port := GetSwaggerServicePort(service)
	if GetProductionMode() == "local" {
		return fmt.Sprintf("%s:%d", host, port)
	}

	return host
}

// GetProductionMode Get app mode.
func GetProductionMode() string {
	return config.Section("").Key("mode").String()
}
