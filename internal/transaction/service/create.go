package service

import (
	"context"
	"core-banking/pkg/dto/request"
	"core-banking/pkg/model"
)

func (svc *TrxService) Create(ctx context.Context, req request.MutationRequest) (bool, error) {
	err := svc.repo.Insert(ctx, &model.Mutation{
		UserUUID: req.UserUUID,
		TrxCode:  req.TrxCode,
		TrxTime:  req.TrxTime,
		Nominal:  req.Nominal,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
