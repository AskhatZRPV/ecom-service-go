package post_additem

import (
	addproducttocart "ecomsvc/internal/application/cart/addproductocart"
)

type requestBody struct {
	Id       int `json:"id"`
	UserId   int `json:"user_id"`
	Quantity int `json:"quantity"`
}

func (r *requestBody) toUsecasePayload() *addproducttocart.Payload {
	return &addproducttocart.Payload{
		ID:       r.Id,
		UserID:   r.UserId,
		Quantity: r.Quantity,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func responseFromResult() *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully added product to cart",
	}
}
