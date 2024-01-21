package pgsql

import "ecomsvc/internal/domain/inventory"

type inventoryRow struct {
	ID       int `db:"id"`
	Quantity int `db:"quantity"`
}

func (i *inventoryRow) ToDomain() *inventory.Inventory {
	return &inventory.Inventory{
		ID:       i.ID,
		Quantity: i.Quantity,
	}
}
