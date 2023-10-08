package request

import (
	"time"

	"github.com/google/uuid"
)

type MutationRequest struct {
	UserUUID uuid.UUID `form:"user_uuid" json:"user_uuid" xml:"user_uuid"`
	TrxCode  string    `form:"trx_code" json:"trx_code" xml:"trx_code"`
	TrxTime  time.Time `form:"trx_time" json:"trx_time" xml:"trx_time"`
	Nominal  float64   `form:"nominal" json:"nominal" xml:"nominal"`
}
