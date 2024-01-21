package pgsql

import "ecomsvc/internal/domain/orderitem"

type orderItemRow struct {
	ID        int `db:"id"`
	OrderId   int `db:"order_id"`
	ProductId int `db:"product_id"`
	Quantity  int `db:"quantity"`
}

func (o *orderItemRow) ToDomain() *orderitem.OrderItem {
	return &orderitem.OrderItem{
		ID:        o.ID,
		OrderId:   o.OrderId,
		ProductId: o.ProductId,
		Quantity:  o.Quantity,
	}
}

func toDomain(o orderItemRow) orderitem.OrderItem {
	return orderitem.OrderItem{
		ID:        o.ID,
		OrderId:   o.OrderId,
		ProductId: o.ProductId,
		Quantity:  o.Quantity,
	}
}
