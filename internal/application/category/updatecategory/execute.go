package updatecategory

import (
	"context"
	"ecomsvc/internal/domain/category"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	res, err := i.cRepo.UpdateById(ctx, category.Update(p.ID, p.Title, p.Description))
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist refresh token")
	}
	return res, nil
}
