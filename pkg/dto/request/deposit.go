package request

type DepositRequest struct {
	Nominal       float64 `form:"nominal" json:"nominal" xml:"nominal"`
	AccountNumber string  `form:"account_number" json:"account_number" xml:"account_number"`
}
