package main

import (
	"net/http"

	"core-banking/config"
	"core-banking/internal/transaction/handler"
	trx "core-banking/internal/transaction/service"
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
	mutationRepo := repository.NewMutationRepository(pgxPool)

	//init service
	trxSvc := trx.NewTransactionService(mutationRepo)

	// init handler
	trxHandler := handler.NewUserHTTPHandler(trxSvc)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World2!")
	})
	e.Add(http.MethodGet, "/mutasi/:account_number", trxHandler.AccountMutation)
	e.Logger.Fatal(e.Start(":8091"))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
