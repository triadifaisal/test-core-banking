package handler

import (
	trxSvc "core-banking/internal/transaction/service"
	"core-banking/pkg/dto/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MutationHTTPHandler struct {
	svc *trxSvc.TrxService
}

func NewUserHTTPHandler(svc *trxSvc.TrxService) *MutationHTTPHandler {
	return &MutationHTTPHandler{svc: svc}
}

func (h *MutationHTTPHandler) AccountMutation(echoCtx echo.Context) error {
	accountNumber := echoCtx.Param("account_number")

	mutation, err := h.svc.AccountMutation(echoCtx.Request().Context(), accountNumber)
	if err != nil {
		return response.ErrorResponse(echoCtx, err)
	}

	response := response.Response{
		Status: true,
		Data: map[string]interface{}{
			"mutasi": mutation,
		},
		Remark: "Cek mutasi berhasil",
	}

	_ = echoCtx.JSON(http.StatusOK, response)
	return nil
}
