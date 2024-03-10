package post_createcategory

import (
	"ecomsvc/internal/application/category/createcategory"
)

type requestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *requestBody) toUsecasePayload() *createcategory.Payload {
	return &createcategory.Payload{
		Title:       r.Title,
		Description: r.Description,
	}
}

type responseBody struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	CategoryId int    `json:"category_id"`
}

func responseFromResult(id int) *responseBody {
	return &responseBody{
		Status:     "Success",
		Message:    "Successfully added product to cart",
		CategoryId: id,
	}
}
