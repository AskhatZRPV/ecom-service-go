package createcart

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/shoppingsession"
)

type Payload struct {
	UserID int
}

type Result = int

type UseCase = usecase.UseCase[*Payload, Result]

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
