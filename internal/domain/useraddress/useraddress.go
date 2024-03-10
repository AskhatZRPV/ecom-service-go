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

func New(userId int, firstName string, lastName string, address string, city string, postalCode string, country string, phoneNumber string) *UserAddress {
	now := time.Now()
	return &UserAddress{
		UserId:      userId,
		FirstName:   firstName,
		LastName:    lastName,
		Address:     address,
		City:        city,
		PostalCode:  postalCode,
		Country:     country,
		PhoneNumber: phoneNumber,
		CreatedAt:   &now,
	}
}

func Update(firstName string, lastName string, address string, city string, postalCode string, country string, phoneNumber string) *UserAddress {
	now := time.Now()
	return &UserAddress{
		FirstName:   firstName,
		LastName:    lastName,
		Address:     address,
		City:        city,
		PostalCode:  postalCode,
		Country:     country,
		PhoneNumber: phoneNumber,
		UpdatedAt:   &now,
	}
}
