package user

import "ecomsvc/internal/core/domainerr"

var (
	ErrUserExitst   = domainerr.New("user exists")
	ErrUserNotFound = domainerr.New("user not found")
)
