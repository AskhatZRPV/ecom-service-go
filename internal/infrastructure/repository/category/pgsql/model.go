package pgsql

import (
	"ecomsvc/internal/domain/category"
)

type categoryRow struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

func (c *categoryRow) ToDomain() *category.Category {
	return &category.Category{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
	}
}

func toDomain(c categoryRow) category.Category {
	return category.Category{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
	}
}
