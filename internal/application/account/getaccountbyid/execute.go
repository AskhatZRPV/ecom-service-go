package getaccountbyid

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.uaRepo.FindById(ctx, p.ID)
	if err != nil {
		return nil, errors.Wrap(err, "cant get user account by id")
	}
	return res, nil
}
