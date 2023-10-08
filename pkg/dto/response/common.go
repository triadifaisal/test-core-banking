package response

type CommonErrorResponse struct {
	Status bool `json:"status"`
	Remark string
}
