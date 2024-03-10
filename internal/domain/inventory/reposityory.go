package inventory

import "context"

type Repository interface {
	Save(ctx context.Context, i *Inventory) (int, error)
	FindById(ctx context.Context, id int) (*Inventory, error)
	GetInIds(ctx context.Context, ids []int) ([]Inventory, error)
	UpdateById(ctx context.Context, i *Inventory) (*Inventory, error)
	Delete(ctx context.Context, id int) (*Inventory, error)
}
