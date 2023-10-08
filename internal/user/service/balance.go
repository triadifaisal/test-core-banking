package service

import (
	"context"
	"core-banking/pkg/reference"
)

func (svc *UserService) CheckBalance(ctx context.Context, an string) (*float64, error) {
	data, err := svc.repo.GetByAccountNumber(ctx, an)
	if err != nil {
		return nil, reference.ErrAccountNumberNotFound
	}

	return &data.Balance, nil
}
