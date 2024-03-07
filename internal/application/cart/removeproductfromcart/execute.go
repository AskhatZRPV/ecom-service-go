package removeproductfromcart

import (
	"context"
	"ecomsvc/internal/domain/cartitem"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	var ciRes *cartitem.CartItem

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		ss, err := i.ssRepo.FindById(ctx, p.ID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		ci, err := i.ciRepo.FindById(ctx, p.CartItemId)
		if err != nil {
			return errors.Wrap(err, "cannot delete cart item")
		}

		pr, err := i.pRepo.FindById(ctx, ci.ProductId)
		if err != nil {
			return errors.Wrap(err, "cannot delete cart item")
		}

		ciRes, err = i.ciRepo.Delete(ctx, p.CartItemId)
		if err != nil {
			return errors.Wrap(err, "cannot delete cart item")
		}

		if _, err := i.ssRepo.UpdateById(ctx, ss.UpdatePrice(ss.TotalPrice-ci.Quantity*pr.Price)); err != nil {
			return errors.Wrap(err, "cannot delete cart item")
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}

	return makeResult(ciRes), nil
}

func makeResult(p *cartitem.CartItem) *Result {
	return &Result{
		ID: p.ID,
	}
}
