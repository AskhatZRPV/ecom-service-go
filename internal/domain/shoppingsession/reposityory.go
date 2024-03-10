package shoppingsession

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, s *ShoppingSession) (int, error)
	FindById(ctx context.Context, id int) (*ShoppingSession, error)
	FindByUserId(ctx context.Context, id int) (*ShoppingSession, error)
	UpdateById(ctx context.Context, s *ShoppingSession) (*ShoppingSession, error)
	Delete(ctx context.Context, id int) (*ShoppingSession, error)
}
