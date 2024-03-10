package pgsql

import (
	"ecomsvc/internal/domain/payment"
	"time"
)

type paymentRow struct {
	ID        int        `db:"id"`
	Amount    int        `db:"amount"`
	CreatedAt *time.Time `db:"created_at"`
}

func (p *paymentRow) ToDomain() *payment.Payment {
	return &payment.Payment{
		ID:        p.ID,
		Amount:    p.Amount,
		CreatedAt: p.CreatedAt,
	}
}
