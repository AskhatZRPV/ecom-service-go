package updatecartitemquantity

import (
	"context"
	"ecomsvc/internal/domain/cartitem"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	var ciRes *cartitem.CartItem

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		ss, err := i.ssRepo.FindById(ctx, p.CartID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		ci, err := i.ciRepo.FindById(ctx, p.CartItemID)
		if err != nil {
			return errors.Wrap(err, "cart item not found")
		}
		if ss.ID != ci.SessionId {
			return errors.New("cart item not found in repo")
		}

		ir, err := i.iRepo.FindById(ctx, ci.ProductId)
		if err != nil {
			return errors.Wrap(err, "product not found")
		}
		if ir.Quantity < p.NewQuantity {
			return errors.New("not enough products in inventory")
		}

		pr, err := i.pRepo.FindById(ctx, ci.ProductId)
		if err != nil {
			return errors.Wrap(err, "product not found")
		}

		if _, err := i.ssRepo.UpdateById(ctx, ss.UpdatePrice(ss.TotalPrice-(ci.Quantity-p.NewQuantity)*pr.Price)); err != nil {
			return errors.Wrap(err, "cannot delete cart item")
		}

		ciRes, err = i.ciRepo.UpdateById(ctx, ci.UpdateQuantity(p.NewQuantity))
		if err != nil {
			return errors.Wrap(err, "cant update cart quantity")
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
