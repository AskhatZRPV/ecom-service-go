package getcart

import (
	"context"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/shoppingsession"

	"github.com/pkg/errors"
)

// FIXME:
func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	var pMap productMap
	var ciRes []cartitem.CartItem
	var ssRes *shoppingsession.ShoppingSession

	err := i.txManager.Do(ctx, func(ctx context.Context) error {
		var err error
		ssRes, err = i.ssRepo.FindById(ctx, p.ID)
		if err != nil {
			return errors.Wrap(err, "cart not found")
		}

		ciRes, err = i.ciRepo.FindAllBySessionId(ctx, ssRes.ID)
		if err != nil {
			return errors.Wrap(err, "cant add products to cart")
		}

		ciSl := make([]int, len(ciRes))
		for i, v := range ciRes {
			ciSl[i] = v.ProductId
		}

		pRes, err := i.pRepo.GetInIds(ctx, ciSl)
		if err != nil {
			return err
		}

		pMap = make(map[int]product.Product, len(pRes))
		for _, v := range pRes {
			pMap[v.ID] = v
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}

	return makeResult(p.ID, ssRes.UserID, ciRes, pMap), nil
}

func makeResult(id int, userId int, ciSl []cartitem.CartItem, pMap productMap) *Result {
	totalPrice := 0
	sciSl := make([]SingleCartItemResult, len(ciSl))

	for i, v := range ciSl {
		pr := pMap[v.ProductId]
		sciSl[i] = *makeSingleCartItemResult(&v, &pr)
		totalPrice += sciSl[i].TotalPrice
	}
	return &Result{
		ID:     id,
		UserId: userId,
		Total:  totalPrice,
		Items:  sciSl,
	}
}

func makeSingleCartItemResult(ci *cartitem.CartItem, p *product.Product) *SingleCartItemResult {
	return &SingleCartItemResult{
		ID:          ci.ID,
		ProductId:   p.ID,
		CategoryId:  p.CategoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    ci.Quantity,
		TotalPrice:  ci.Quantity * p.Price,
	}
}
