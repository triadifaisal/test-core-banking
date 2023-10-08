package service

import (
	"context"
	"core-banking/pkg/model"
	"core-banking/pkg/reference"
)

func (svc *TrxService) AccountMutation(ctx context.Context, an string) ([]*model.Mutation, error) {
	data, err := svc.repo.GetAccountMutationByAccountNumber(ctx, an)
	if err != nil {
		return nil, reference.ErrAccountNumberNotFound
	}

	return data, nil
}
