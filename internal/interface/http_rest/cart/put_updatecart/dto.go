package put_updatecart

import (
	"ecomsvc/internal/application/cart/updatecartitemquantity"
	"fmt"
)

type requestBody struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (r *requestBody) toUsecasePayload() *updatecartitemquantity.Payload {
	return &updatecartitemquantity.Payload{
		CartID:      r.Id,
		CartItemID:  r.ProductId,
		NewQuantity: r.Quantity,
	}
}

type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"string"`
}

func responseFromResult(r *updatecartitemquantity.Result) *responseBody {
	return &responseBody{
		Status:  "Success",
		Message: fmt.Sprintf("Successfully updated cart with ID - %d", r.ID),
	}
}
