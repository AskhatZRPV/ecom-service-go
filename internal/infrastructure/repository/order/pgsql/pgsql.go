package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/order"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	"ecomsvc/pkg/utils/mapper"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) order.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, o *order.Order) (int, error) {
	const createOrder = `
		INSERT INTO order (payment_id, user_id, status, created_at)
		VALUES($1, $2, $3, $4);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, createOrder, o.PaymentId, o.UserId, o.Status, o.CreatedAt)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create order record")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) FindById(ctx context.Context, id int) (*order.Order, error) {
	const insertOrderQuery = `
		SELECT * FROM order
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderRow
	if err := q.GetContext(ctx, &row, insertOrderQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindAllByUserId(ctx context.Context, id int) ([]order.Order, error) {
	const insertOrderQuery = `
		SELECT * FROM order
		WHERE user_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []orderRow
	if err := q.GetContext(ctx, &rows, insertOrderQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return mapper.Map[orderRow, order.Order](rows, toDomain), nil
}

func (r *repo) UpdateById(ctx context.Context, o *order.Order) (*order.Order, error) {
	const updateById = `
		UPDATE order SET payment_id = $2, user_id = $3, status = $4 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderRow
	if err := q.GetContext(ctx, &row, updateById, o.ID, o.PaymentId, o.UserId, o.Status); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*order.Order, error) {
	const deleteById = `
		DELETE FROM order WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
