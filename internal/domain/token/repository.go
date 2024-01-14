package token

import (
	"context"
	"ecomsvc/internal/domain/user"
)

type GrantResult struct {
	Access, Refresh Token
}

type Provider interface {
	Grant(context.Context, user.ID) (*GrantResult, error)
}
