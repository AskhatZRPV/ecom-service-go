package order

import "context"

type Repository interface {
	Save(ctx context.Context, o *Order) (int, error)
	FindById(ctx context.Context, id int) (*Order, error)
	FindAllByUserId(ctx context.Context, id int) ([]Order, error)
	UpdateById(ctx context.Context, o *Order) (*Order, error)
	Delete(ctx context.Context, id int) (*Order, error)
}
