package addproducttocart

import (
	"context"
	"ecomsvc/internal/domain/cartitem"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		ss, err := i.ssRepo.FindByUserId(ctx, p.UserID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		product, err := i.pRepo.FindById(ctx, p.ID)
		if err != nil {
			return errors.Wrap(err, "can't find product")
		}

		inv, err := i.iRepo.FindById(ctx, product.InventoryId)
		if err != nil {
			return errors.Wrap(err, "product not found in inventory")
		}

		if inv.Quantity < p.Quantity {
			return errors.New("not enough products in inventory")
		}
		_, err = i.ciRepo.Save(ctx, cartitem.New(ss.ID, p.ID, p.Quantity))
		if err != nil {
			return errors.Wrap(err, "cant add products to cart")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to run transaction")
	}
	return nil
}
