package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/product"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

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

func (i *repo) Save(ctx context.Context, u *product.Product) error {
	const insertUserQuery = `
		INSERT INTO products (id, title, description, price, category_id) VALUES($1, $2, $3, $4, $5);
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	if _, err := q.ExecContext(ctx, insertUserQuery, u.ID, u.Title, u.Description, u.Price, u.CategoryId); err != nil {
		return errors.Wrap(err, "failed to insert new product record")
	}

	return nil
}

func (i *repo) FindById(ctx context.Context, id int) (*product.Product, error) {
	const selectUserByUsernameQuery = `
		SELECT * FROM products WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	var row productRow
	err := q.GetContext(ctx, &row, selectUserByUsernameQuery, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "product not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
