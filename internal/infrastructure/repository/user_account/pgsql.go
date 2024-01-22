package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/domain/useraccount"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) useraccount.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, u *useraccount.UserAccount) error {
	const createSession = `
		INSERT INTO user_account (user_id, balance)
		VALUES($1, $2);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, createSession, u.UserId, u.Balance)
	if err != nil {
		return errors.Wrap(err, "failed to create user account record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*useraccount.UserAccount, error) {
	const selectUserAccountByIdQuery = `
		SELECT * FROM user_account
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAccountRow
	if err := q.GetContext(ctx, &row, selectUserAccountByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user account not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindByUserId(ctx context.Context, id int) (*useraccount.UserAccount, error) {
	const selectUserAccountByUserIdQuery = `
		SELECT * FROM user_account
		WHERE user_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAccountRow
	if err := q.GetContext(ctx, &row, selectUserAccountByUserIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user account not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) UpdateById(ctx context.Context, s *useraccount.UserAccount) (*useraccount.UserAccount, error) {
	const updateUserAccountByIdQuery = `
		UPDATE user_account SET user_id = $2, balance = $3 
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAccountRow
	if err := q.GetContext(ctx, &row, updateUserAccountByIdQuery, s.ID, s.UserId, s.Balance); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user account not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*useraccount.UserAccount, error) {
	const deleteById = `
		DELETE FROM user_account WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAccountRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user account not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
