package pgsql

import (
	"ecomsvc/internal/domain/product"
)

type productRow struct {
	ID          int    `db:"id"`
	CategoryId  int    `db:"category_id"`
	InventoryId int    `db:"inventory_id"`
	SKU         string `db:"SKU"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
}

func (p *productRow) ToDomain() *product.Product {
	return &product.Product{
		ID:          p.ID,
		CategoryId:  p.CategoryId,
		InventoryId: p.InventoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
	}
}

func toDomain(p productRow) product.Product {
	return product.Product{
		ID:          p.ID,
		CategoryId:  p.CategoryId,
		InventoryId: p.InventoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
	}
}
