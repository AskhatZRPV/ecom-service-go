package pgsql

import "ecomsvc/internal/domain/cartitem"

type cartItemRow struct {
	ID        int `db:"id"`
	SessionId int `db:"session_id"`
	ProductId int `db:"product_id"`
	Quantity  int `db:"quantity"`
}

func (c *cartItemRow) ToDomain() *cartitem.CartItem {
	return &cartitem.CartItem{
		ID:        c.ID,
		SessionId: c.SessionId,
		ProductId: c.ProductId,
		Quantity:  c.Quantity,
	}
}

func toDomain(c cartItemRow) cartitem.CartItem {
	return cartitem.CartItem{
		ID:        c.ID,
		SessionId: c.SessionId,
		ProductId: c.ProductId,
		Quantity:  c.Quantity,
	}
}
