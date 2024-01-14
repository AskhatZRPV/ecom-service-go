package category

import "time"

type OrderItem struct {
	ID         int
	CustomerId int
	Status     string
	CreatedAt  time.Time
}
