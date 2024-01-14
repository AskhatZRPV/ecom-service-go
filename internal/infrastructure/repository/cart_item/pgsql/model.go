package pgsql

type cartItemRow struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
}
