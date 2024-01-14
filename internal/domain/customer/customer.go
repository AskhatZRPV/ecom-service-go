package customer

import "time"

type Customer struct {
	ID         int
	CustomerId int
	Status     string
	CreatedAt  time.Time
}
