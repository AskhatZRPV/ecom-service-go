package post_createorder

import (
	"ecomsvc/internal/application/order/createorder"
)

type requestBody struct {
	CartID int `json:"cart_id"`
	UserID int `json:"user_id"`
}

func (r *requestBody) toUsecasePayload() *createorder.Payload {
	return &createorder.Payload{
		CartID: r.CartID,
		UserID: r.UserID,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	OrderId int    `json:"order_id"`
}

func responseFromResult(id int) *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: "Successfully created order",
		OrderId: id,
	}
}
