package pgsql

import "ecomsvc/internal/domain/product"

type productRow struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	CategoryId  int    `db:"category_id"`
}

func (p *productRow) ToDomain() *product.Product {
	return &product.Product{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		CategoryId:  p.CategoryId,
	}
}
