package pgsql

import (
	"context"
	"database/sql"
	"ecomsvc/internal/domain/orderitem"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	"ecomsvc/pkg/utils/mapper"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) orderitem.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, o *orderitem.OrderItem) error {
	const createOrderItemQuery = `
		INSERT INTO order (order_id, product_id, quantity)
		VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, createOrderItemQuery, o.OrderId, o.ProductId, o.Quantity)
	if err != nil {
		return errors.Wrap(err, "failed to create order item record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*orderitem.OrderItem, error) {
	const insertOrderQuery = `
		SELECT * FROM order_item
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderItemRow
	if err := q.GetContext(ctx, &row, insertOrderQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindAllByOrderId(ctx context.Context, id int) ([]orderitem.OrderItem, error) {
	const insertOrderItemQuery = `
		SELECT * FROM order_item
		WHERE order_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []orderItemRow
	if err := q.GetContext(ctx, &rows, insertOrderItemQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return mapper.Map[orderItemRow, orderitem.OrderItem](rows, toDomain), nil
}

func (r *repo) UpdateById(ctx context.Context, o *orderitem.OrderItem) (*orderitem.OrderItem, error) {
	const updateById = `
		UPDATE order_item SET order_id = $2, product_id = $3, quantity = $4 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderItemRow
	if err := q.GetContext(ctx, &row, updateById, o.ID, o.OrderId, o.ProductId, o.Quantity); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*orderitem.OrderItem, error) {
	const deleteById = `
		DELETE FROM order_item WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row orderItemRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "order item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
