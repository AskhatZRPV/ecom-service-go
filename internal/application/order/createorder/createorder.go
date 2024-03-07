package getcart

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/cartitem"
	"ecomsvc/internal/domain/order"
	"ecomsvc/internal/domain/orderitem"
	"ecomsvc/internal/domain/payment"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/shoppingsession"
	"ecomsvc/internal/domain/useraccount"
)

type Payload struct {
	CartID int
	UserID int
}

type UseCase = usecase.UseCase[*Payload, *Result]

type Result struct {
	ID int
}

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

type productMap map[int]product.Product

type implementation struct {
	txManager tx.TransactionManager
	uaRepo    useraccount.Repository
	prRepo    product.Repository
	orRepo    order.Repository
	oiRepo    orderitem.Repository
	paRepo    payment.Repository
	ssRepo    shoppingsession.Repository
	ciRepo    cartitem.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	ssRepo shoppingsession.Repository,
	ciRepo cartitem.Repository,
) UseCase {
	return &implementation{txManager, pRepo, ssRepo, ciRepo}
}
