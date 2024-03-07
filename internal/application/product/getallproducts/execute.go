package getallproducts

import (
	"context"
	"ecomsvc/internal/domain/product"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (Result, error) {
	var invMap inventoryMap
	var prodRes []product.Product

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		var err error
		prodRes, err = i.pRepo.GetAll(ctx, p.Limit, p.Offset)
		if err != nil {
			return errors.Wrap(err, "failed to persist refresh token")
		}

		invIds := make([]int, len(prodRes))
		for i, v := range prodRes {
			invMap[i] = v.InventoryId
		}

		invRes, err := i.iRepo.GetInIds(ctx, invIds)
		if err != nil {
			return errors.Wrap(err, "failed to persist refresh token")
		}

		invMap = make(inventoryMap, len(invRes))
		for _, v := range invRes {
			invMap[v.ID] = v.Quantity
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist refresh token")
	}

	return makeProductResult(prodRes, invMap), nil
}

func makeProductResult(p []product.Product, invMap inventoryMap) Result {
	res := make(Result, len(p))
	for _, v := range p {
		res = append(res, *makeSingleProductResult(&v, invMap[v.InventoryId]))
	}

	return res
}

func makeSingleProductResult(p *product.Product, quantity int) *SingleResult {
	return &SingleResult{
		ID:          p.ID,
		CategoryId:  p.CategoryId,
		InventoryId: p.InventoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    quantity,
	}
}
