package createproduct

import (
	"context"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (Result, error) {
	sku := uuid.New().String()
	var invId int
	var res int

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		_, err := i.iRepo.Save(ctx, inventory.New(p.Quantity))
		if err != nil {
			return errors.Wrap(err, "failed to save product in inventory repository")
		}
		invId = res

		res, err = i.pRepo.Save(ctx, product.New(p.CategoryId, invId, sku, p.Title, p.Description, p.Price))
		if err != nil {
			return errors.Wrap(err, "failed to save product in repository")
		}
		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to run transaction")
	}

	return res, nil
}
