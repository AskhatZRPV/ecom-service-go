package createcategory

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/category"
)

type Payload struct {
	Title       string
	Description string
}

type Result = int

type UseCase = usecase.UseCase[*Payload, Result]

type implementation struct {
	cRepo category.Repository
}

func New(
	cRepo category.Repository,
) UseCase {
	return &implementation{cRepo}
}
