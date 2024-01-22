package useraddress

import "context"

type Repository interface {
	Save(ctx context.Context, u *UserAddress) error
	FindById(ctx context.Context, id int) (*UserAddress, error)
	FindByUserId(ctx context.Context, id int) (*UserAddress, error)
	UpdateById(ctx context.Context, u *UserAddress) (*UserAddress, error)
	Delete(ctx context.Context, id int) (*UserAddress, error)
}
