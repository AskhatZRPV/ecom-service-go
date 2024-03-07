package addproducttocart

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/shoppingsession"
)

type Payload struct {
	ID       int
	UserID   int
	Quantity int
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	pRepo     product.Repository
	ssRepo    shoppingsession.Repository
	iRepo     inventory.Repository
	ciRepo    cartitem.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	ssRepo shoppingsession.Repository,
	iRepo inventory.Repository,
	ciRepo cartitem.Repository,
) UseCase {
	return &implementation{txManager, pRepo, ssRepo, iRepo, ciRepo}
}
