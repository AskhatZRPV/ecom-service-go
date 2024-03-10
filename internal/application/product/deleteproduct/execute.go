package deleteproduct

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		product, err := i.pRepo.FindById(ctx, p.ID)
		if err != nil {
			return errors.Wrap(err, "cannot find product with specified id")
		}
		if _, err := i.iRepo.Delete(ctx, product.InventoryId); err != nil {
			return errors.Wrap(err, "failed to drop entry from inventory repo")
		}
		if _, err := i.pRepo.Delete(ctx, p.ID); err != nil {
			return errors.Wrap(err, "failed to drop entry from inventory repo")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to run transaction")
	}
	return nil
}
