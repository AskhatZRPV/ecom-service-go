package updatecartitemquantity

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/shoppingsession"
)

type Payload struct {
	CartID      int
	CartItemID  int
	NewQuantity int
}

type UseCase = usecase.UseCase[*Payload, *Result]

type Result struct {
	ID int
}
type implementation struct {
	txManager tx.TransactionManager
	pRepo     product.Repository
	iRepo     inventory.Repository
	ssRepo    shoppingsession.Repository
	ciRepo    cartitem.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	iRepo inventory.Repository,
	ssRepo shoppingsession.Repository,
	ciRepo cartitem.Repository,
) UseCase {
	return &implementation{txManager, pRepo, iRepo, ssRepo, ciRepo}
}
