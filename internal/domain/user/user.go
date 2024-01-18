package user

import (
	"time"
)

type ID int

type User struct {
	ID

	Username, Password, Role string

	CreatedAt, UpdatedAt, LastLoginAttempt *time.Time
}

func New(username, password string) *User {
	now := time.Now()
	return &User{
		Username:  username,
		Password:  password,
		Role:      "user",
		CreatedAt: &now,
	}
}

func FromData(id ID, username, password, role string, created, updated, lastLogin *time.Time) *User {
	return &User{
		id,
		username,
		password,
		role,
		created,
		updated,
		lastLogin,
	}
}
