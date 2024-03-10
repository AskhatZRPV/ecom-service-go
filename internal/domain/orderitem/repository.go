package orderitem

import "context"

type Repository interface {
	Save(ctx context.Context, o *OrderItem) (int, error)
	FindById(ctx context.Context, id int) (*OrderItem, error)
	FindAllByOrderId(ctx context.Context, id int) ([]OrderItem, error)
	UpdateById(ctx context.Context, o *OrderItem) (*OrderItem, error)
	Delete(ctx context.Context, id int) (*OrderItem, error)
}
