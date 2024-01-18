package main

import (
	"context"
	"ecomsvc/internal/application/auth/login"
	"ecomsvc/internal/application/auth/register"
	"ecomsvc/internal/infrastructure/crosscutting/bcrypt"
	"ecomsvc/internal/infrastructure/crosscutting/pgsql_client"
	viperconfig "ecomsvc/internal/infrastructure/crosscutting/viper_config"
	"ecomsvc/internal/infrastructure/paseto_tokens"
	sessionRepo "ecomsvc/internal/infrastructure/repository/session/pgsql"
	userRepo "ecomsvc/internal/infrastructure/repository/user/pgsql"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	httprest "ecomsvc/internal/interface/http_rest"
	"ecomsvc/internal/interface/http_rest/auth/post_login"
	"ecomsvc/internal/interface/http_rest/auth/post_register"
	"ecomsvc/internal/interface/http_rest/common"
	"os/signal"
	"syscall"
)

func main() {
	var handlers []common.Handler

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := viperconfig.New()
	if err != nil {
		panic(err)
	}

	tp, err := paseto_tokens.New(cfg)
	if err != nil {
		panic(err)
	}

	bh := bcrypt.New()

	pgc, err := pgsql_client.New(cfg)
	if err != nil {
		panic(err)
	}

	tx := pgsqltx.New(pgc)

	ur := userRepo.New(pgc)
	sr := sessionRepo.New(pgc)

	luc := login.New(tx, bh, ur, tp, sr)
	ruc := register.New(tx, bh, ur)

	handlers = append(
		handlers,
		post_login.New(luc),
		post_register.New(ruc),
	)

	httprest.New(ctx, cfg, handlers)
}
