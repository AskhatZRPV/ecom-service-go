package payment

import "context"

type Repository interface {
	Save(ctx context.Context, p *Payment) (int, error)
	FindById(ctx context.Context, id int) (*Payment, error)
	UpdateById(ctx context.Context, p *Payment) (*Payment, error)
	Delete(ctx context.Context, id int) (*Payment, error)
}
