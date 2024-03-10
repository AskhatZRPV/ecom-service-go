package createcart

import (
	"context"
	"ecomsvc/internal/domain/shoppingsession"
	"ecomsvc/internal/domain/user"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (int, error) {
	_, err := i.ssRepo.FindByUserId(ctx, p.UserID)
	if err == nil || !errors.Is(err, user.ErrUserNotFound) {
		return 0, errors.Wrap(ErrAccountAlreadyExists, "shopping cart already exists")
	}

	res, err := i.ssRepo.Save(ctx, shoppingsession.New(p.UserID))
	if err != nil {
		return 0, errors.Wrap(err, "unexpected query error")
	}
	return res, nil
}
