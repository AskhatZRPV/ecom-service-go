package paseto_tokens

import (
	"context"
	"ecomsvc/internal/core/config"
	"ecomsvc/internal/domain/token"
	"ecomsvc/internal/domain/user"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/chacha20poly1305"
)

type provider struct {
	authCfg config.Auth
	pst     *paseto.V2
}

func New(config *config.Config) (token.Provider, error) {
	if len(config.AccessTokenSecret) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("access token secret key size is incorrect")
	}
	if len(config.RefreshTokenSecret) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("refresh token secret key size is incorrect")
	}

	return &provider{config.Auth, paseto.NewV2()}, nil
}

func (p *provider) Grant(ctx context.Context, userID user.ID) (r *token.GrantResult, err error) {
	sessID := uuid.New().String()
	cf := token.NewClaimsFactory(sessID, userID)
	at, rt := token.New(time.Now().Add(p.authCfg.AccessTokenTTL)), token.New(time.Now().Add(p.authCfg.RefreshTokenTTL))
	ac, rc := cf(at.ExpiresAt), cf(rt.ExpiresAt)

	at.Value, err = p.signToken(p.authCfg.AccessTokenSecret, ac)
	if err != nil {
		return nil, err
	}

	rt.Value, err = p.signToken(p.authCfg.RefreshTokenSecret, rc)
	if err != nil {
		return nil, err
	}

	r = &token.GrantResult{}
	r.Access = *at
	r.Refresh = *rt

	return
}

func (p *provider) signToken(secret string, c *token.Claims) (string, error) {
	token, err := p.pst.Encrypt([]byte(secret), c, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")
	}

	return token, nil

}
