package deleteproduct

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"
)

type Payload struct {
	ID int
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	pRepo     product.Repository
	iRepo     inventory.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	iRepo inventory.Repository,
) UseCase {
	return &implementation{txManager, pRepo, iRepo}
}
