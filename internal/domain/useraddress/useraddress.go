package useraddress

import "time"

type UserAddress struct {
	ID          int
	UserId      int
	FirstName   string
	LastName    string
	Address     string
	City        string
	PostalCode  string
	Country     string
	PhoneNumber string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
