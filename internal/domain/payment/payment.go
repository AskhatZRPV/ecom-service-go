package payment

import "time"

type Payment struct {
	ID        int
	OrderId   int
	Amount    int
	CreatedAt *time.Time
}
