package createproduct

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"
)

type Payload struct {
	CategoryId  int
	Title       string
	Description string
	Quantity    int
	Price       int
}
type Result = int

type UseCase = usecase.UseCase[*Payload, Result]

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
