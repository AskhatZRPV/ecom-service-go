package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/cartitem"
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

func New(db *sqlx.DB) cartitem.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, c *cartitem.CartItem) (int, error) {
	const insertCartItemQuery = `
		INSERT INTO cart_item (session_id, product_id, quantity) VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, insertCartItemQuery, c.SessionId, c.ProductId, c.Quantity)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert new cart_item record")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) FindById(ctx context.Context, id int) (*cartitem.CartItem, error) {
	const selectCartItemByIdQuery = `
		SELECT * FROM cart_item WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row cartItemRow
	if err := q.GetContext(ctx, &row, selectCartItemByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "cart_item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindAllBySessionId(ctx context.Context, id int) ([]cartitem.CartItem, error) {
	const selectCartItemByIdQuery = `
		SELECT * FROM cart_item WHERE session_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []cartItemRow
	if err := q.SelectContext(ctx, &rows, selectCartItemByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "cart_item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return mapper.Map[cartItemRow, cartitem.CartItem](rows, toDomain), nil
}

func (r *repo) UpdateById(ctx context.Context, c *cartitem.CartItem) (*cartitem.CartItem, error) {
	const updateById = `
		UPDATE cart_item SET session_id = $2, product_id = $3, quantity = $4 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row cartItemRow
	if err := q.GetContext(ctx, &row, updateById, c.ID, c.SessionId, c.ProductId, c.Quantity); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "cart_item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*cartitem.CartItem, error) {
	const deleteById = `
		DELETE FROM cart_item WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row cartItemRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "cart_item not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
