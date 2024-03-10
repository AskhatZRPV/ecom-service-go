package delete_deleteproduct

import (
	"ecomsvc/internal/application/product/deleteproduct"
)

type requestParams struct {
	Id int `json:"id"`
}

func (r *requestParams) toUsecasePayload() *deleteproduct.Payload {
	return &deleteproduct.Payload{
		ID: r.Id,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func responseFromResult() *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully delete a category",
	}
}
