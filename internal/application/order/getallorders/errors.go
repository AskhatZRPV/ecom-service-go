package getallorders

import "ecomsvc/internal/core/domainerr"

var (
	ErrAccountAlreadyExists = domainerr.New("account already exists")
)
