package createcategory

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/category"
)

type Payload struct {
	Title       string
	Description string
}

type Result = category.Category

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	cRepo category.Repository
}

func New(
	cRepo category.Repository,
) UseCase {
	return &implementation{cRepo}
}
