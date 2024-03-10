package get_singleproduct

import (
	"ecomsvc/internal/application/product/getproductbyid"
)

type requestParams struct {
	ID int `params:"id"`
}

func (r *requestParams) toUsecasePayload() *getproductbyid.Payload {
	return &getproductbyid.Payload{
		ID: r.ID,
	}
}

type responseBody struct {
	ID          int    `json:"id,omitempty"`
	CategoryId  int    `json:"category_id,omitempty"`
	InventoryId int    `json:"inventory_id,omitempty"`
	SKU         string `json:"sku,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
}

func responseFromResult(r *getproductbyid.Result) *responseBody {
	return &responseBody{
		ID:          r.ID,
		CategoryId:  r.CategoryId,
		InventoryId: r.InventoryId,
		SKU:         r.SKU,
		Title:       r.Title,
		Description: r.Description,
		Price:       r.Price,
	}
}
