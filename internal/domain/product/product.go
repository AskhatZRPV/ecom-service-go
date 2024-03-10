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

func New(categoryId int, inventoryId int, sku string, title string, description string, price int) *Product {
	return &Product{
		CategoryId:  categoryId,
		InventoryId: inventoryId,
		SKU:         sku,
		Title:       title,
		Description: description,
		Price:       price,
	}
}

func (p *Product) ToProductResult(quantity int) *ProductResult {
	return &ProductResult{
		CategoryId:  p.CategoryId,
		InventoryId: p.InventoryId,
		SKU:         p.SKU,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    quantity,
	}
}

type ProductResult struct {
	ID          int
	CategoryId  int
	InventoryId int
	SKU         string
	Title       string
	Description string
	Price       int
	Quantity    int
}

func NewResult(categoryId int, inventoryId int, sku string, title string, description string, price int, quantity int) *ProductResult {
	return &ProductResult{
		CategoryId:  categoryId,
		InventoryId: inventoryId,
		SKU:         sku,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}
