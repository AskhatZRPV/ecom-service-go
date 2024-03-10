package deleteaddress

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/useraddress"
)

type Payload struct {
	ID int
}

type Result = useraddress.UserAddress

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	uaRepo useraddress.Repository
}

func New(
	uaRepo useraddress.Repository,
) UseCase {
	return &implementation{uaRepo}
}
