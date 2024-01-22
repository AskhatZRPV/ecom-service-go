package product

type Product struct {
	ID          int
	CategoryId  int
	InventoryId int
	SKU         string
	Title       string
	Description string
	Price       int
}

func New(id int, title string, categoryId int, inventoryId int, sku string, description string, price int) *Product {
	return &Product{
		ID:          id,
		CategoryId:  categoryId,
		InventoryId: inventoryId,
		SKU:         sku,
		Title:       title,
		Description: description,
		Price:       price,
	}
}
