package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/category"
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

func New(db *sqlx.DB) category.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, c *category.Category) error {
	const insertCategoryQuery = `
		INSERT INTO category (title, description) VALUES($1, $2);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	if _, err := q.ExecContext(ctx, insertCategoryQuery, c.Title, c.Description); err != nil {
		return errors.Wrap(err, "failed to insert new product record")
	}

	return nil
}

func (r *repo) FindById(ctx context.Context, id int) (*category.Category, error) {
	const selectCategoryByIdQuery = `
		SELECT * FROM category WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row categoryRow
	if err := q.GetContext(ctx, &row, selectCategoryByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "category not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) GetAll(ctx context.Context) ([]category.Category, error) {
	const getAllCategories = `
		SELECT * FROM category;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var rows []categoryRow
	if err := q.GetContext(ctx, &rows, getAllCategories); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "category not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}
	return mapper.Map[categoryRow, category.Category](rows, toDomain), nil
}

func (r *repo) UpdateById(ctx context.Context, c *category.Category) (*category.Category, error) {
	const updateById = `
		UPDATE category SET title = $2, description = $3 WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row categoryRow
	if err := q.GetContext(ctx, &row, updateById, c.ID, c.Title, c.Description); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "category not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*category.Category, error) {
	const deleteById = `
		DELETE FROM category WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row categoryRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "category not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
