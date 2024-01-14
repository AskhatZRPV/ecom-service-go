package token

import "ecomsvc/internal/core/domainerr"

var (
	ErrTokenExpired = domainerr.New("token expired")
	ErrInvalidSign  = domainerr.New("invalid sign")
)
