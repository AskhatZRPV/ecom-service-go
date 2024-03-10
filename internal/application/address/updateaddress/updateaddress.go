package updateaddress

import (
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/useraddress"
)

type Payload struct {
	UserId      int
	FirstName   string
	LastName    string
	Address     string
	City        string
	PostalCode  string
	Country     string
	PhoneNumber string
}

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	uaRepo useraddress.Repository
}

func New(
	uaRepo useraddress.Repository,
) UseCase {
	return &implementation{uaRepo}
}
