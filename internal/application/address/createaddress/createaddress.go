package createaddress

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
