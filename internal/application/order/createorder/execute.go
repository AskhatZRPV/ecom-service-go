package getcart

import (
	"context"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/order"
	"ecomsvc/internal/domain/orderitem"
	"ecomsvc/internal/domain/payment"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	var pMap productMap
	var ciRes []cartitem.CartItem

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		ss, err := i.ssRepo.FindById(ctx, p.CartID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		ciRes, err = i.ciRepo.FindAllBySessionId(ctx, ss.ID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		ua, err := i.uaRepo.FindByUserId(ctx, p.UserID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}
		if ua.Balance < ss.TotalPrice {
			errors.New("not enough money on account")
		}

		paId, err := i.paRepo.Save(ctx, payment.New(ss.TotalPrice))
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		orId, err := i.orRepo.Save(ctx, order.New(paId, p.UserID))
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		for _, v := range ciRes {
			i.ciRepo.Delete(ctx, v.ID)
			i.oiRepo.Save(ctx, orderitem.New(orId, v.ProductId, v.Quantity))
		}

		_, err = i.ssRepo.Delete(ctx, p.CartID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		_, err = i.uaRepo.UpdateById(ctx, ua.Update(ua.Balance-ss.TotalPrice))
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}

	return makeResult(p.ID, p.UserID, ciRes, pMap), nil
}
