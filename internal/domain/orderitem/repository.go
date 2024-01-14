package orderitem

import "context"

type Repository interface {
	Save(context.Context, *OrderItem) error
}
