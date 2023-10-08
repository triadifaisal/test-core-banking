package service

import "core-banking/pkg/repository"

type TrxService struct {
	repo *repository.MutationRepository
}

func NewTransactionService(repo *repository.MutationRepository) *TrxService {
	return &TrxService{repo: repo}
}
