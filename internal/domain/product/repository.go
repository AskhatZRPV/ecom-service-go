package product

import "context"

type Repository interface {
	Save(ctx context.Context, p *Product) (int, error)
	FindById(ctx context.Context, id int) (*Product, error)
	GetInIds(ctx context.Context, ids []int) ([]Product, error)
	GetAll(ctx context.Context, limit int, offset int) ([]Product, error)
	UpdateById(ctx context.Context, p *Product) (*Product, error)
	Delete(ctx context.Context, id int) (*Product, error)
}
