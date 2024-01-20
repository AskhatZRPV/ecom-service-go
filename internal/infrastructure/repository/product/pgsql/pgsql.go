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

func (r *repo) Save(ctx context.Context, p *product.Product) error {
	const insertUserQuery = `
		INSERT INTO product (id, title, description, price, category_id) VALUES($1, $2, $3, $4, $5);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	if _, err := q.ExecContext(ctx, insertUserQuery, p.ID, p.Title, p.Description, p.Price, p.CategoryId); err != nil {
		return errors.Wrap(err, "failed to insert new product record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*product.Product, error) {
	const selectProductByIdQuery = `
		SELECT * FROM product WHERE id = $1;
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

func (r *repo) GetAll(ctx context.Context) ([]product.Product, error) {
	const getAllProducts = `
		SELECT * FROM product;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []productRow
	if err := q.GetContext(ctx, &rows, getAllProducts); err != nil {
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
		UPDATE product SET category_id = $1, inventory_id = $2, SKU = $3, name = $4, description = $5, price = $6;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row productRow
	if err := q.GetContext(ctx, &row, updateById, p.CategoryId, p.InventoryId, p.SKU, p.Title, p.Description, p.Price); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, p *product.Product) (*product.Product, error) {
	const deleteById = `
		DELETE FROM product WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row productRow
	if err := q.GetContext(ctx, &row, deleteById, p.ID); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
