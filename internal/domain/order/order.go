package order

import "time"

type Order struct {
	ID        int
	PaymentId int
	UserId    int
	Status    string
	CreatedAt *time.Time
}
