package deletecategory

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.cRepo.Delete(ctx, p.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete category")
	}
	return res, nil
}
