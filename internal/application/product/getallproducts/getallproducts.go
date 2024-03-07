package getallproducts

import (
	"ecomsvc/internal/core/tx"
	"ecomsvc/internal/core/usecase"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/product"
)

type Payload struct {
	Limit  int
	Offset int
}

type Result = []SingleResult

type SingleResult struct {
	ID          int
	CategoryId  int
	InventoryId int
	SKU         string
	Title       string
	Description string
	Price       int
	Quantity    int
}

type inventoryMap = map[int]int

type UseCase = usecase.UseCase[*Payload, Result]

type implementation struct {
	txManager tx.TransactionManager
	pRepo     product.Repository
	iRepo     inventory.Repository
}

func New(
	txManager tx.TransactionManager,
	pRepo product.Repository,
	iRepo inventory.Repository,
) UseCase {
	return &implementation{txManager, pRepo, iRepo}
}
