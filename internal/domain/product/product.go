package product

type Product struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	CategoryId  int    `db:"category_id"`
}

func New(id int, title string, description string, price int, categoryId int) *Product {
	return &Product{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       price,
		CategoryId:  categoryId,
	}
}
