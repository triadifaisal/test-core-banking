package reference

import "errors"

const (
	ErrUniqueViolationCode = "23505"
)

var (
	ErrUniqueHpNik                   = errors.New("user already registered")
	ErrGenerateAccountNumber         = errors.New("error generate account number")
	ErrAccountNumberNotFound         = errors.New("account number not found")
	ErrBalanceZero                   = errors.New("account number has zero balance")
	ErrBalanceLessThanWithrawNominal = errors.New("cannot withdraw more than balance")
	ErrInvalidRequestPayload         = errors.New("invalid request")
)
