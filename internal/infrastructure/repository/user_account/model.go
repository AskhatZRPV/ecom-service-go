package pgsql

import (
	"ecomsvc/internal/domain/useraccount"
)

type userAccountRow struct {
	ID      int `db:"id"`
	UserId  int `db:"user_id"`
	Balance int `db:"balance"`
}

func (u *userAccountRow) ToDomain() *useraccount.UserAccount {
	return &useraccount.UserAccount{
		ID:      u.ID,
		UserId:  u.UserId,
		Balance: u.Balance,
	}
}
