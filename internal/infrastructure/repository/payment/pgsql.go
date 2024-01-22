package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/payment"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) payment.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, p *payment.Payment) error {
	const insertPaymentQuery = `
		INSERT INTO payment (order_id, amount, created_at)
		VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	if _, err := q.ExecContext(ctx, insertPaymentQuery, p.ID, p.OrderId, p.CreatedAt); err != nil {
		return errors.Wrap(err, "failed to insert new payment record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*payment.Payment, error) {
	const selectPaymentByIdQuery = `
		SELECT * FROM payment 
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row paymentRow
	if err := q.GetContext(ctx, &row, selectPaymentByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "payment not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) UpdateById(ctx context.Context, p *payment.Payment) (*payment.Payment, error) {
	const updateByIdQuery = `
		UPDATE payment
		SET order_id = $2, amount = $3
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row paymentRow
	if err := q.GetContext(ctx, &row, updateByIdQuery, p.ID, p.OrderId, p.Amount); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "payment not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*payment.Payment, error) {
	const deleteById = `
		DELETE FROM payment WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row paymentRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "payment not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
