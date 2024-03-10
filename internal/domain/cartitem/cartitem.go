package cartitem

type CartItem struct {
	ID        int
	SessionId int
	ProductId int
	Quantity  int
}

func New(sessionId, productId, quantity int) *CartItem {
	return &CartItem{
		SessionId: sessionId,
		ProductId: productId,
		Quantity:  quantity,
	}
}

func (c *CartItem) UpdateQuantity(quantity int) *CartItem {
	return &CartItem{
		c.ID,
		c.SessionId,
		c.ProductId,
		quantity,
	}
}
