package handler

import (
	userSvc "core-banking/internal/user/service"
	"core-banking/pkg/dto/request"
	"core-banking/pkg/dto/response"
	"core-banking/pkg/reference"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHTTPHandler struct {
	svc *userSvc.UserService
}

func NewUserHTTPHandler(svc *userSvc.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{svc: svc}
}

func (h *UserHTTPHandler) Create(echoCtx echo.Context) error {
	req := request.CreateRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return response.ErrorResponse(echoCtx, reference.ErrInvalidRequestPayload)
	}

	accountNumber, err := h.svc.Create(echoCtx.Request().Context(), req)
	if err != nil {
		return response.ErrorResponse(echoCtx, err)
	}

	response := response.Response{
		Status: true,
		Data: map[string]interface{}{
			"no_rekening": accountNumber,
		},
		Remark: "Pendaftaran user berhasil",
	}

	_ = echoCtx.JSON(http.StatusCreated, response)
	return nil
}

func (h *UserHTTPHandler) Deposit(echoCtx echo.Context) error {
	req := request.DepositRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return response.ErrorResponse(echoCtx, reference.ErrInvalidRequestPayload)
	}

	balance, err := h.svc.Deposit(echoCtx.Request().Context(), req)
	if err != nil {
		return response.ErrorResponse(echoCtx, err)
	}

	response := response.Response{
		Status: true,
		Data: map[string]interface{}{
			"saldo": balance,
		},
		Remark: fmt.Sprintf("Deposit berhasil ke no rekening: %v", req.AccountNumber),
	}

	_ = echoCtx.JSON(http.StatusOK, response)
	return nil
}

func (h *UserHTTPHandler) Withdraw(echoCtx echo.Context) error {
	req := request.WithdrawRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return response.ErrorResponse(echoCtx, reference.ErrInvalidRequestPayload)
	}

	balance, err := h.svc.Withdraw(echoCtx.Request().Context(), req)
	if err != nil {
		return response.ErrorResponse(echoCtx, err)
	}

	response := response.Response{
		Status: true,
		Data: map[string]interface{}{
			"saldo": balance,
		},
		Remark: "Penarikan berhasil",
	}

	_ = echoCtx.JSON(http.StatusOK, response)
	return nil
}
