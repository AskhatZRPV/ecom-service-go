package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/product"
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

func New(db *sqlx.DB) product.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, p *product.Product) (int, error) {
	const insertUserQuery = `
		INSERT INTO product (id, title, description, price, category_id) 
		VALUES($1, $2, $3, $4, $5);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, insertUserQuery, p.ID, p.Title, p.Description, p.Price, p.CategoryId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert new product record")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) FindById(ctx context.Context, id int) (*product.Product, error) {
	const selectProductByIdQuery = `
		SELECT * FROM product 
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row productRow
	if err := q.GetContext(ctx, &row, selectProductByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) GetInIds(ctx context.Context, ids []int) ([]product.Product, error) {
	const selectProductsByIdQuery = `
		SELECT * FROM product 
		WHERE id IN (?);
	`

	query, args, err := sqlx.In(selectProductsByIdQuery, ids)
	if err != nil {
		return nil, err
	}

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []productRow
	if err := q.SelectContext(ctx, &rows, query, args...); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return mapper.Map[productRow, product.Product](rows, toDomain), nil
}

func (r *repo) GetAll(ctx context.Context, limit int, offset int) ([]product.Product, error) {
	const getAllProducts = `
		SELECT * FROM product
		LIMIT $1 OFFSET $2;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []productRow
	if err := q.SelectContext(ctx, &rows, getAllProducts, limit, offset); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}
	return mapper.Map[productRow, product.Product](rows, toDomain), nil
}

func (r *repo) UpdateById(ctx context.Context, p *product.Product) (*product.Product, error) {
	const updateById = `
		UPDATE product SET category_id = $2, inventory_id = $3, SKU = $4, name = $5, description = $6, price = $7 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row productRow
	if err := q.GetContext(ctx, &row, updateById, p.ID, p.CategoryId, p.InventoryId, p.SKU, p.Title, p.Description, p.Price); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*product.Product, error) {
	const deleteById = `
		DELETE FROM product WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row productRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
