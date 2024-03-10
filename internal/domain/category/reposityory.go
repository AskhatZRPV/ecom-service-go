package category

import "context"

type Repository interface {
	Save(ctx context.Context, c *Category) (int, error)
	FindById(ctx context.Context, id int) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	UpdateById(ctx context.Context, c *Category) (*Category, error)
	Delete(ctx context.Context, id int) (*Category, error)
}
