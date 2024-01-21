package cartitem

import "context"

type Repository interface {
	FindById(ctx context.Context, id int) (*CartItem, error)
	FindAllBySessionId(ctx context.Context, id int) ([]CartItem, error)
	UpdateById(ctx context.Context, c *CartItem) (*CartItem, error)
	Delete(ctx context.Context, id int) (*CartItem, error)
}
