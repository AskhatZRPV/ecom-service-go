package useraccount

import "context"

type Repository interface {
	Save(ctx context.Context, u *UserAccount) (int, error)
	FindById(ctx context.Context, id int) (*UserAccount, error)
	FindByUserId(ctx context.Context, id int) (*UserAccount, error)
	UpdateById(ctx context.Context, u *UserAccount) (*UserAccount, error)
	Delete(ctx context.Context, id int) (*UserAccount, error)
}
