package put_updatecategory

import (
	"ecomsvc/internal/application/category/updatecategory"
)

type requestBody struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *requestBody) toUsecasePayload() *updatecategory.Payload {
	return &updatecategory.Payload{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
	}
}

type responseBody struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func responseFromResult(r *updatecategory.Result) *responseBody {
	return &responseBody{}
}
