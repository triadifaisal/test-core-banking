package main

import (
	"net/http"

	"core-banking/config"
	"core-banking/internal/user/handler"
	user "core-banking/internal/user/service"
	"core-banking/pkg/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// init config
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// init db
	pgxPool, err := config.NewPgx(*cfg)
	checkError(err)

	// init repo
	userRepo := repository.NewUserRepository(pgxPool)

	//init service
	userSvc := user.NewUserService(userRepo)

	// init handler
	userHandler := handler.NewUserHTTPHandler(userSvc)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World2!")
	})
	e.Add(http.MethodPost, "/daftar", userHandler.Create)
	e.Add(http.MethodPost, "/tabung", userHandler.Deposit)
	e.Add(http.MethodPost, "/tarik", userHandler.Withdraw)
	e.Add(http.MethodGet, "/saldo/:account_number", userHandler.CheckBalance)
	e.Logger.Fatal(e.Start(":8090"))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
