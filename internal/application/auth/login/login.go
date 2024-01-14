package login

import (
	"ecomsvc/internal/core/hasher"
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/session"
	"ecomsvc/internal/domain/token"
	"ecomsvc/internal/domain/user"
)

type Payload struct {
	Username, Password string
}

type Result = token.GrantResult

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	txManager     tx.TransactionManager
	ph            hasher.Hasher
	userRepo      user.Repository
	tokenProvider token.Provider
	refreshRepo   session.Repository
}

func New(
	txManager tx.TransactionManager,
	ph hasher.Hasher,
	userRepo user.Repository,
	tokenProvider token.Provider,
	refreshRepo session.Repository,
) UseCase {
	return &implementation{txManager, ph, userRepo, tokenProvider, refreshRepo}
}
