package deletecart

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/shoppingsession"
)

type Payload struct {
	ID int
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	ssRepo    shoppingsession.Repository
}

func New(
	txManager tx.TransactionManager,
	ssRepo shoppingsession.Repository,
) UseCase {
	return &implementation{txManager, ssRepo}
}
