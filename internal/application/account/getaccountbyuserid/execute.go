package getaccountbyuserid

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.uaRepo.FindByUserId(ctx, p.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist refresh token")
	}
	return res, nil
}
