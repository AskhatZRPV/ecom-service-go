package delete_deletecart

import (
	"ecomsvc/internal/application/cart/deletecart"
)

type requestParams struct {
	ID int `params:"id"`
}

func (r *requestParams) toUsecasePayload() *deletecart.Payload {
	return &deletecart.Payload{
		ID: r.ID,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func makeResponse() *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully deleted the cart",
	}
}
