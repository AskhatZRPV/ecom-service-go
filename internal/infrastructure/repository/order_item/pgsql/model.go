package pgsql

type orderItemRow struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	CategoryId  int    `db:"category_id"`
}
