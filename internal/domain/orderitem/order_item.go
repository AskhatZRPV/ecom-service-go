package orderitem

type OrderItem struct {
	ID        int
	OrderId   int
	ProductId int
	Quantity  int
}

func New(orderId, productId, quantity int) *OrderItem {
	return &OrderItem{
		OrderId:   orderId,
		ProductId: productId,
		Quantity:  quantity,
	}
}
