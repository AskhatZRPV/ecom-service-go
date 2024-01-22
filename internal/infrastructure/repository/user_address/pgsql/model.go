package pgsql

import (
	"ecomsvc/internal/domain/useraddress"
	"time"
)

type userAddressRow struct {
	ID          int        `db:"id"`
	UserId      int        `db:"user_id"`
	FirstName   string     `db:"first_name"`
	LastName    string     `db:"last_name"`
	Address     string     `db:"address"`
	City        string     `db:"city"`
	PostalCode  string     `db:"postal_code"`
	Country     string     `db:"country"`
	PhoneNumber string     `db:"phone_number"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

func (u *userAddressRow) ToDomain() *useraddress.UserAddress {
	return &useraddress.UserAddress{
		ID:          u.ID,
		UserId:      u.UserId,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Address:     u.Address,
		City:        u.City,
		PostalCode:  u.PostalCode,
		Country:     u.Country,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
