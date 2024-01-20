package pgsql

import (
	"ecomsvc/internal/domain/shoppingsession"
)

type shoppingSessionRow struct {
	Id         int `db:"id"`
	UserID     int `db:"user_id"`
	TotalPrice int `db:"total_price"`
}

func (s *shoppingSessionRow) ToDomain() *shoppingsession.ShoppingSession {
	return &shoppingsession.ShoppingSession{
		ID:         s.Id,
		UserID:     s.UserID,
		TotalPrice: s.TotalPrice,
	}
}
