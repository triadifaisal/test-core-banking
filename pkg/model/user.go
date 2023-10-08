package model

import (
	"core-banking/pkg/reference"
	"database/sql"
)

// User struct based on table
type User struct {
	Base
	Name          string `gorm:"not null"`
	NIK           string `gorm:"not null;column:nik"`
	Phonenumber   sql.NullString
	AccountNumber string  `gorm:"not null"`
	Balance       float64 `gorm:"not null;column:latest_balance;default:0"`
}

func (User) TableName() string {
	return "users.users"
}

func (u *User) AddBalance(nominal float64) {
	if u.Balance == 0 {
		u.Balance = nominal
		return
	}

	newBalance := u.Balance + nominal
	u.Balance = newBalance
}

func (u *User) SubtractBalance(nominal float64) error {
	if u.Balance == 0 {
		return reference.ErrBalanceZero
	}

	if u.Balance < nominal {
		return reference.ErrBalanceLessThanWithrawNominal
	}

	newBalance := u.Balance - nominal
	u.Balance = newBalance

	return nil
}
