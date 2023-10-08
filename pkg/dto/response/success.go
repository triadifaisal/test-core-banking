package response

type Response struct {
	Status bool                   `json:"status"`
	Data   map[string]interface{} `json:"data"`
	Remark string
}
