package order

import "time"

type Order struct {
	ID         int
	CustomerId int
	Status     string
	CreatedAt  time.Time
}
