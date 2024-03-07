package addproducttocart

import (
	"context"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/user"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		ss, err := i.ssRepo.FindByUserId(ctx, p.UserID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		product, err := i.pRepo.FindById(ctx, p.ID)
		if err == nil || !errors.Is(err, user.ErrUserNotFound) {
			return errors.Wrap(ErrAccountAlreadyExists, "account with such username exists")
		}

		inv, err := i.iRepo.FindById(ctx, product.InventoryId)
		if err == nil || !errors.Is(err, user.ErrUserNotFound) {
			return errors.Wrap(ErrAccountAlreadyExists, "account with such username exists")
		}

		if inv.Quantity < p.Quantity {
			return errors.New("not enough products in inventory")
		}
		i.
		if err := i.ciRepo.Save(ctx, cartitem.New(ss.ID, p.ID, p.Quantity)); err != nil {
			return errors.Wrap(err, "cant add products to cart")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to run transaction")
	}
	return nil
}
