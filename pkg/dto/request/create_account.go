package request

type CreateRequest struct {
	Name        string `form:"name" json:"name" xml:"name"`
	NIK         string `form:"nik" json:"nik" xml:"nik"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" xml:"phonenumber"`
}
