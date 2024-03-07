package createcategory

import (
	"context"
	"ecomsvc/internal/domain/category"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	if err := i.cRepo.Save(ctx, category.New(p.Title, p.Description)); err != nil {
		return errors.Wrap(err, "failed to persist refresh token")
	}
	return nil
}
