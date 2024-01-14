package register

import (
	"ecomsvc/internal/core/hasher"
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/user"
)

type Payload struct {
	Username, Password string
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	ph        hasher.Hasher
	userRepo  user.Repository
}

func New(
	txManager tx.TransactionManager,
	ph hasher.Hasher,
	userRepo user.Repository,
) UseCase {
	return &implementation{txManager, ph, userRepo}
}
