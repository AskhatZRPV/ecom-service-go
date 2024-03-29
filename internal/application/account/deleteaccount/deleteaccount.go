package deleteaccount

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/useraccount"
)

type Payload struct {
	ID int
}

type Result = useraccount.UserAccount

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	uaRepo useraccount.Repository
}

func New(
	uaRepo useraccount.Repository,
) UseCase {
	return &implementation{uaRepo}
}
