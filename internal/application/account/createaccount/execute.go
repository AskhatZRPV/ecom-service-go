package createaccount

import (
	"context"
	"ecomsvc/internal/domain/useraccount"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (Result, error) {
	res, err := i.uaRepo.Save(ctx, useraccount.New(p.UserId, p.Balance))
	if err != nil {
		return 0, errors.Wrap(err, "failed to create user account")
	}
	return res, nil
}
