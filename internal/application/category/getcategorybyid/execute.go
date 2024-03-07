package getcategorybyid

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.cRepo.FindById(ctx, p.ID)
	if err != nil {
		return nil, errors.Wrap(err, "cant get category by id")
	}
	return res, nil
}
