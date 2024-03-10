package updateproduct

import (
	"context"

	"github.com/pkg/errors"
)

// FIXME
func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	var res *Result

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		pres, err := i.pRepo.FindById(ctx, p.ID)
		if err != nil {
			return errors.Wrap(err, "failed to find product in repository")
		}
		ires, err := i.iRepo.FindById(ctx, pres.InventoryId)
		if err != nil {
			return errors.Wrap(err, "failed to find quantity in repository")
		}

		res = pres.ToProductResult(ires.Quantity)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}
	return res, nil
}
