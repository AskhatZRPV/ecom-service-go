package pgsql

import (
	"context"
	"database/sql"
	"ecomsvc/internal/domain/inventory"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	"ecomsvc/pkg/utils/mapper"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) inventory.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, i *inventory.Inventory) (int, error) {
	const insertInventoryQuery = `
		INSERT INTO inventory (quantity) VALUES($1);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, insertInventoryQuery, i.Quantity)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert new inventory record")
	}
	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) GetInIds(ctx context.Context, ids []int) ([]inventory.Inventory, error) {
	const selectInventoryByIdQuery = `
		SELECT * FROM inventory 
		WHERE id IN (?);
	`

	query, args, err := sqlx.In(selectInventoryByIdQuery, ids)
	if err != nil {
		return nil, err
	}

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []inventoryRow
	if err := q.SelectContext(ctx, &rows, query, args...); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "inventory not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return mapper.Map[inventoryRow, inventory.Inventory](rows, toDomain), nil
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
