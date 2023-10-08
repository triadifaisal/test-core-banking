package repository

import (
	"context"
	"core-banking/pkg/helper"
	"core-banking/pkg/model"
	"core-banking/pkg/reference"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (q *UserRepository) Insert(ctx context.Context, request *model.User) error {
	err := q.db.Create(request).Error
	return err
}

func (q *UserRepository) CheckNikAndPhoneNumberExist(ctx context.Context, nik, phoneNumber string) (*model.User, error) {
	var result = &model.User{}
	if res := q.db.Where(model.User{
		NIK: nik,
		Phonenumber: sql.NullString{
			String: phoneNumber,
			Valid:  true,
		},
	}).Take(result); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}

func (q *UserRepository) CheckAccountNumberExist(ctx context.Context, an string) (*bool, error) {
	var result *model.User
	if res := q.db.Where(model.User{
		AccountNumber: an,
	}).Take(&result); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return helper.ToPointer(false), nil
	} else if res.Error != nil {
		return nil, res.Error
	}

	return helper.ToPointer(true), nil
}

func (q *UserRepository) GetByAccountNumber(ctx context.Context, an string) (*model.User, error) {
	var result *model.User
	if res := q.db.Where(model.User{
		AccountNumber: an,
	}).Take(&result); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, reference.ErrAccountNumberNotFound
	} else if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}

func (q *UserRepository) UpdateBalance(ctx context.Context, an string, balance float64) error {
	err := q.db.Where(model.User{
		AccountNumber: an,
	}).Updates(model.User{
		Balance: balance,
	}).Error

	return err
}
