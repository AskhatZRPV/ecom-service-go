package pgsql

import (
	"ecomsvc/internal/domain/order"
	"time"
)

type orderRow struct {
	ID        int        `db:"id"`
	PaymentId int        `db:"payment_id"`
	UserId    int        `db:"user_id"`
	Status    string     `db:"status"`
	CreatedAt *time.Time `db:"created_at"`
}

func (o *orderRow) ToDomain() *order.Order {
	return &order.Order{
		ID:        o.ID,
		PaymentId: o.PaymentId,
		UserId:    o.UserId,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
	}
}

func toDomain(o orderRow) order.Order {
	return order.Order{
		ID:        o.ID,
		PaymentId: o.PaymentId,
		UserId:    o.UserId,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
	}
}
