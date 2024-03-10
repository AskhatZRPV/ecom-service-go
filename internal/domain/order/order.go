package order

import "time"

type Order struct {
	ID        int
	PaymentId int
	UserId    int
	Status    string
	CreatedAt *time.Time
}

func New(paymentId int, userId int) *Order {
	now := time.Now()
	return &Order{
		PaymentId: paymentId,
		UserId:    userId,
		Status:    "CREATED",
		CreatedAt: &now,
	}
}
