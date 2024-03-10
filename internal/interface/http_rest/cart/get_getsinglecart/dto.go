package get_getsinglecart

import (
	"ecomsvc/internal/application/cart/getcart"
	"ecomsvc/pkg/utils/mapper"
)

type requestParams struct {
	ID int `json:"id"`
}

func (r *requestParams) toUsecasePayload() *getcart.Payload {
	return &getcart.Payload{
		ID: r.ID,
	}
}

type responseBody struct {
	ID     int                    `json:"id"`
	UserId int                    `json:"user_id"`
	Total  int                    `json:"total"`
	Items  []singleCartItemResult `json:"items"`
}

type singleCartItemResult struct {
	ID          int    `json:"id"`
	ProductId   int    `json:"product_id"`
	CategoryId  int    `json:"category_id"`
	SKU         string `json:"sku"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}

func responseFromResult(r *getcart.Result) *responseBody {
	products := mapper.Map[getcart.SingleCartItemResult, singleCartItemResult](r.Items, toDomain)
	return &responseBody{
		ID:     r.ID,
		UserId: r.UserId,
		Total:  r.Total,
		Items:  products,
	}
}

func toDomain(p getcart.SingleCartItemResult) singleCartItemResult {
	return singleCartItemResult{
		ID:          p.ID,
		ProductId:   p.ProductId,
		CategoryId:  p.CategoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		TotalPrice:  p.TotalPrice,
	}
}
