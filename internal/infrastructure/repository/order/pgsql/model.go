package pgsql

import (
	"time"
)

type orderRow struct {
	Id         int
	CustomerId int
	Status     string
	CreatedAt  time.Time
}
