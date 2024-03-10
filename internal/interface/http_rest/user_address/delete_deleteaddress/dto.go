package delete_deleteaddress

import (
	"ecomsvc/internal/application/address/deleteaddress"
)

type requestBody struct {
	Id int `json:"id"`
}

func (r *requestBody) toUsecasePayload() *deleteaddress.Payload {
	return &deleteaddress.Payload{
		ID: r.Id,
	}
}

type responseBody struct {
	ID          int
	UserId      int
	FirstName   string
	LastName    string
	Address     string
	City        string
	PostalCode  string
	Country     string
	PhoneNumber string
}

func responseFromResult(r *deleteaddress.Result) *responseBody {
	return &responseBody{
		ID:          r.ID,
		UserId:      r.UserId,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Address:     r.Address,
		City:        r.City,
		PostalCode:  r.PostalCode,
		Country:     r.Country,
		PhoneNumber: r.PhoneNumber,
	}
}
