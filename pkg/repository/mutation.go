package repository

import (
	"context"
	"core-banking/pkg/model"
	"core-banking/pkg/reference"
	"errors"

	"gorm.io/gorm"
)

type MutationRepository struct {
	db *gorm.DB
}

func NewMutationRepository(db *gorm.DB) *MutationRepository {
	return &MutationRepository{db: db}
}

func (q *MutationRepository) Insert(ctx context.Context, request *model.Mutation) error {
	err := q.db.Create(request).Error
	return err
}

func (q *MutationRepository) GetAccountMutationByAccountNumber(ctx context.Context, an string) ([]*model.Mutation, error) {
	var result []*model.Mutation
	if res := q.db.Model(&model.Mutation{}).
		Joins("left join users.users as u on u.uuid = transactions.mutations.user_uuid").
		Where("u.account_number = ?", an).Find(&result); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, reference.ErrAccountNumberNotFound
	} else if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}
