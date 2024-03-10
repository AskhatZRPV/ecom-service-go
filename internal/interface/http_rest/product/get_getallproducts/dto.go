package get_getallproducts

import (
	"ecomsvc/internal/application/product/getallproducts"
	"ecomsvc/pkg/utils/mapper"
)

type queryParams struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func (r *queryParams) toUsecasePayload() *getallproducts.Payload {
	return &getallproducts.Payload{
		Limit:  r.Limit,
		Offset: r.Offset,
	}
}

type responseBody struct {
	Products []product `json:"products"`
}

type product struct {
	ID          int    `json:"id"`
	CategoryId  int    `json:"category_id"`
	InventoryId int    `json:"inventory_id"`
	SKU         string `json:"sku"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

func responseFromResult(r getallproducts.Result) *responseBody {
	products := mapper.Map[getallproducts.SingleResult, product](r, toDomain)
	return &responseBody{
		Products: products,
	}
}

func toDomain(p getallproducts.SingleResult) product {
	return product{
		ID:          p.ID,
		CategoryId:  p.CategoryId,
		InventoryId: p.InventoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}
}
