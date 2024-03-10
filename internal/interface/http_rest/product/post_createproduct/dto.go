package post_createproduct

import (
	"ecomsvc/internal/application/product/createproduct"
)

type requestBody struct {
	CategoryId  int    `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

func (r *requestBody) toUsecasePayload() *createproduct.Payload {
	return &createproduct.Payload{
		CategoryId:  r.CategoryId,
		Title:       r.Title,
		Description: r.Description,
		Quantity:    r.Quantity,
		Price:       r.Price,
	}
}

type responseBody struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ProductId int    `json:"product_id"`
}

func responseFromResult(id int) *responseBody {
	return &responseBody{
		Status:    "Success",
		Message:   "Successfully created product",
		ProductId: id,
	}
}
