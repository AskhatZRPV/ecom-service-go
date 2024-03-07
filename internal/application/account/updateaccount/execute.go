package updateaccount

import (
	"context"
	"ecomsvc/internal/domain/useraccount"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.uaRepo.UpdateById(ctx, useraccount.Update(p.Balance))
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist refresh token")
	}
	return res, nil
}
