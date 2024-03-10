package removeproductfromcart

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/shoppingsession"
)

type Payload struct {
	ID         int
	CartItemId int
}

type UseCase = usecase.UseCase[*Payload, *Result]

type Result struct {
	ID int
}

type implementation struct {
	txManager tx.TransactionManager
	pRepo     product.Repository
	ssRepo    shoppingsession.Repository
	ciRepo    cartitem.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	ssRepo shoppingsession.Repository,
	ciRepo cartitem.Repository,
) UseCase {
	return &implementation{txManager, pRepo, ssRepo, ciRepo}
}
