package put_updateaddress

import (
	"ecomsvc/internal/application/address/updateaddress"
)

type requestBody struct {
	UserId      int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
}

func (r *requestBody) toUsecasePayload() *updateaddress.Payload {
	return &updateaddress.Payload{
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

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func responseFromResult() *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully updated address",
	}
}
