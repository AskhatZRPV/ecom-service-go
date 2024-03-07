package deletecategory

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/category"
)

type Payload struct {
	ID int
}

type Result = category.Category

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	cRepo category.Repository
}

func New(
	cRepo category.Repository,
) UseCase {
	return &implementation{cRepo}
}
