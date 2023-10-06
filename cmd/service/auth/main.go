package main

import (
	authHandleRest "corebanking/internal/auth/handler/rest"
	"corebanking/internal/auth/handler/rest/docs"
	"corebanking/pkg/helper/config"
	"corebanking/pkg/helper/logger"
	loggerfactory "corebanking/pkg/helper/logger_factory"
	"fmt"
)

const (
	prefixAppConfig = "AUTH"
)

func main() {
	if err := loggerfactory.LoadLogger(config.GetLogConfig()); err != nil {
		fmt.Printf("%v", err)
		panic("failed to initialize logger")
	}
	logger.Log.Info("Log initialized...")

	docs.SwaggerInfo.Host = config.GetSwaggerServiceHost(prefixAppConfig)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if err := authHandleRest.StartServer(config.GetSwaggerServicePort(prefixAppConfig)); err != nil {
		panic(err)
	}
}
