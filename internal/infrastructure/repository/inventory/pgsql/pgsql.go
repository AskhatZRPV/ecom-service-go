package pgsql

import (
	"context"
	"database/sql"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) inventory.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, i *inventory.Inventory) error {
	const insertInventoryQuery = `
		INSERT INTO inventory (quantity) VALUES($1);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	if _, err := q.ExecContext(ctx, insertInventoryQuery, i.Quantity); err != nil {
		return errors.Wrap(err, "failed to insert new inventory record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*inventory.Inventory, error) {
	const selectInventoryByIdQuery = `
		SELECT * FROM inventory WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row inventoryRow
	if err := q.GetContext(ctx, &row, selectInventoryByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "inventory not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) UpdateById(ctx context.Context, i *inventory.Inventory) (*inventory.Inventory, error) {
	const updateById = `
		UPDATE inventory SET quantity = $2 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row inventoryRow
	if err := q.GetContext(ctx, &row, updateById, i.ID, i.Quantity); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "inventory not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*inventory.Inventory, error) {
	const deleteById = `
		DELETE FROM inventory WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row inventoryRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "inventory not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
