package delete_deletecategory

import (
	"ecomsvc/internal/application/category/deletecategory"
)

type requestParams struct {
	ID int `params:"id"`
}

func (r *requestParams) toUsecasePayload() *deletecategory.Payload {
	return &deletecategory.Payload{
		ID: r.ID,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func responseFromResult(r *deletecategory.Result) *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully delete a category",
	}
}
