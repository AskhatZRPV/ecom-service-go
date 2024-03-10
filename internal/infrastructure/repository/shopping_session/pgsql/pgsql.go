package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/shoppingsession"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) shoppingsession.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, s *shoppingsession.ShoppingSession) (int, error) {
	const createSession = `
		INSERT INTO shopping_session (user_id, total_price)
		VALUES($1, $2);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, createSession, s.UserID, s.TotalPrice)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create session record")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) FindById(ctx context.Context, id int) (*shoppingsession.ShoppingSession, error) {
	const insertSessionQuery = `
		SELECT * FROM shopping_session
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row shoppingSessionRow
	if err := q.GetContext(ctx, &row, insertSessionQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "shopping session not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindByUserId(ctx context.Context, id int) (*shoppingsession.ShoppingSession, error) {
	const insertSessionQuery = `
		SELECT * FROM shopping_session
		WHERE user_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row shoppingSessionRow
	if err := q.GetContext(ctx, &row, insertSessionQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "shopping session not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) UpdateById(ctx context.Context, s *shoppingsession.ShoppingSession) (*shoppingsession.ShoppingSession, error) {
	const updateById = `
		UPDATE shopping_session SET total_price = $2 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row shoppingSessionRow
	if err := q.GetContext(ctx, &row, updateById, s.ID, s.TotalPrice); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "shopping session not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*shoppingsession.ShoppingSession, error) {
	const deleteById = `
		DELETE FROM shopping_session WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row shoppingSessionRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "shopping session not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
