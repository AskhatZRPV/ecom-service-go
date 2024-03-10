package inventory

type Inventory struct {
	ID       int
	Quantity int
}

func New(quantity int) *Inventory {
	return &Inventory{
		Quantity: quantity,
	}
}
