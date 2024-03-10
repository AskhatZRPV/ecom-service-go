package createorder

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/order"
	"ecomsvc/internal/domain/orderitem"
	"ecomsvc/internal/domain/payment"
	"ecomsvc/internal/domain/shoppingsession"
	"ecomsvc/internal/domain/useraccount"
)

type Payload struct {
	CartID int
	UserID int
}

type UseCase = usecase.UseCase[*Payload, Result]

type Result = int

type SingleCartItemResult struct {
	ID          int
	ProductId   int
	CategoryId  int
	SKU         string
	Title       string
	Description string
	Price       int
	Quantity    int
	TotalPrice  int
}

type implementation struct {
	txManager tx.TransactionManager
	uaRepo    useraccount.Repository
	orRepo    order.Repository
	oiRepo    orderitem.Repository
	paRepo    payment.Repository
	ssRepo    shoppingsession.Repository
	ciRepo    cartitem.Repository
}

func New(
	txManager tx.TransactionManager,
	uaRepo useraccount.Repository,
	orRepo order.Repository,
	oiRepo orderitem.Repository,
	paRepo payment.Repository,
	ssRepo shoppingsession.Repository,
	ciRepo cartitem.Repository,
) UseCase {
	return &implementation{
		txManager,
		uaRepo,
		orRepo,
		oiRepo,
		paRepo,
		ssRepo,
		ciRepo,
	}
}
