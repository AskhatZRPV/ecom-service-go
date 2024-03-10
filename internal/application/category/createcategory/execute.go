package createcategory

import (
	"context"
	"ecomsvc/internal/domain/category"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (Result, error) {
	res, err := i.cRepo.Save(ctx, category.New(p.Title, p.Description))
	if err != nil {
		return 0, errors.Wrap(err, "failed to persist refresh token")
	}
	return res, nil
}
