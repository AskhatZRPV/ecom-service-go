package updateaddress

import (
	"ecomsvc/internal/core/tx"
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

type UseCase = usecase.Interactor[*Payload]

type implementation struct {
	txManager tx.TransactionManager
	uaRepo    useraddress.Repository
}

func New(
	txManager tx.TransactionManager,
	uaRepo useraddress.Repository,
) UseCase {
	return &implementation{txManager, uaRepo}
}
