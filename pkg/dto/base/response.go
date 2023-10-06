package dtobase

import (
	"time"

	"github.com/google/uuid"
)

// BaseRes envelope struct for non paginated response
type BaseRes struct {
	Code       int     `json:"code" validate:"required"`
	Message    string  `json:"message" validate:"required"`
	Stacktrace *string `json:"stacktrace,omitempty"`
}

func (b *BaseRes) Error() string {
	return b.Message
}

// BasePagination base page struct
type BasePagination struct {
	Offset  int    `json:"offset" validate:"required"`
	Limit   int    `json:"limit" validate:"required"`
	Count   int    `json:"count" validate:"required"`
	OrderBy string `json:"order_by" validate:"required"`
}

// BaseResPagination envelope struct for paginated response
type BaseResPagination struct {
	BaseRes
	Page BasePagination `json:"page"`
}

// BaseResTime envelope struct for entity timestamps
type BaseResTime struct {
	CreatedAt time.Time  `json:"created_at" validate:"required"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by"`
}
