package product

import "context"

type Repository interface {
	Save(ctx context.Context, u *Product) error
	FindById(ctx context.Context, id int) (*Product, error)
}
