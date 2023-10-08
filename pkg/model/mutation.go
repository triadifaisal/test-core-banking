package model

import (
	"time"

	"github.com/google/uuid"
)

type Mutation struct {
	Base
	UserUUID uuid.UUID `gorm:"not null"`
	TrxCode  string    `gorm:"not null"`
	TrxTime  time.Time `gorm:"default:now()"`
	Nominal  float64   `gorm:"not null"`
}

func (Mutation) TableName() string {
	return "transactions.mutations"
}
