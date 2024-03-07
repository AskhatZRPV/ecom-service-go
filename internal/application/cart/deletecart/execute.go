package deletecart

import (
	"context"
	"ecomsvc/internal/domain/shoppingsession"
	"ecomsvc/internal/domain/user"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	_, err := i.ssRepo.FindByUserId(ctx, p.UserID)
	if err != nil || errors.Is(err, user.ErrUserNotFound) {
		return errors.Wrap(ErrAccountAlreadyExists, "shopping cart does not exists")
	}

	if err := i.ssRepo.Save(ctx, shoppingsession.New(p.UserID)); err != nil {
		return errors.Wrap(err, "unexpected query error")
	}
	return nil
}
