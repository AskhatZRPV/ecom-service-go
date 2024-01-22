package inventory

import "context"

type Repository interface {
	Save(ctx context.Context, i *Inventory) error
	FindById(ctx context.Context, id int) (*Inventory, error)
	UpdateById(ctx context.Context, i *Inventory) (*Inventory, error)
	Delete(ctx context.Context, id int) (*Inventory, error)
}
