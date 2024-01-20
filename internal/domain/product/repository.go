package product

import "context"

type Repository interface {
	Save(ctx context.Context, p *Product) error
	FindById(ctx context.Context, id int) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	UpdateById(ctx context.Context, p *Product) (*Product, error)
	Delete(ctx context.Context, p *Product) (*Product, error)
}
