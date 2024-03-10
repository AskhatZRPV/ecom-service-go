package deleteaddress

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.uaRepo.Delete(ctx, p.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user address")
	}
	return res, nil
}
