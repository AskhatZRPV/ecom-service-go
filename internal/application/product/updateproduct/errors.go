package updateproduct

import "ecomsvc/internal/core/domainerr"

var (
	ErrAccountDoesNotExist = domainerr.New("account does not exist")
	ErrIncorrectPassword   = domainerr.New("err incorrect password")
)
