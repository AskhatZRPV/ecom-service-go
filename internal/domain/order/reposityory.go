package order

import "context"

type Repository interface {
	Save(context.Context, *Order) error
}
