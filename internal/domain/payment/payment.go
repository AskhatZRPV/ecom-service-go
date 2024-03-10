package payment

import "time"

type Payment struct {
	ID        int
	Amount    int
	CreatedAt *time.Time
}

func New(amount int) *Payment {
	now := time.Now()
	return &Payment{
		Amount:    amount,
		CreatedAt: &now,
	}
}
