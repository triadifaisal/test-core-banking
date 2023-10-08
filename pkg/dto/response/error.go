package response

import (
	"core-banking/pkg/reference"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorResponse function to handle error on handler.
func ErrorResponse(ctx echo.Context, err error) error {
	code := http.StatusInternalServerError
	switch err {
	case reference.ErrUniqueHpNik,
		reference.ErrInvalidRequestPayload,
		reference.ErrBalanceLessThanWithrawNominal,
		reference.ErrBalanceZero:
		code = http.StatusBadRequest
	}

	errResp := CommonErrorResponse{
		Status: false,
		Remark: err.Error(),
	}
	_ = ctx.JSON(code, errResp)
	return err
}
