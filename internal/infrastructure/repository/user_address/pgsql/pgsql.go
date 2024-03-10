package pgsql

import (
	"database/sql"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/domain/useraddress"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) useraddress.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, c *useraddress.UserAddress) (int, error) {
	const insertUserAddressQuery = `
		INSERT INTO user_address (user_id, first_name, last_name, address_line1, city, postal_code, country, phone_number, created_at, updated_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	qRes, err := q.ExecContext(ctx, insertUserAddressQuery,
		c.UserId,
		c.FirstName,
		c.LastName,
		c.Address, c.City,
		c.PostalCode, c.Country,
		c.PhoneNumber,
		c.CreatedAt,
		c.UpdatedAt,
	)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert new user_address record")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (r *repo) FindById(ctx context.Context, id int) (*useraddress.UserAddress, error) {
	const selectUserAddressByIdQuery = `
		SELECT * FROM user_address WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAddressRow
	if err := q.GetContext(ctx, &row, selectUserAddressByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user_address not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) FindByUserId(ctx context.Context, id int) (*useraddress.UserAddress, error) {
	const selectUserAddressByIdQuery = `
		SELECT * FROM user_address WHERE user_id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAddressRow
	if err := q.GetContext(ctx, &row, selectUserAddressByIdQuery, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user_address not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) UpdateById(ctx context.Context, i *useraddress.UserAddress) (*useraddress.UserAddress, error) {
	const updateById = `
		UPDATE user_address 
		SET first_name = $2, last_name = $3, address_line1 = $4, city = $5, postal_code = $6, country = $7, phone_number $8, updated_at = $9 
		WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAddressRow
	if err := q.GetContext(ctx, &row, updateById,
		i.ID,
		i.FirstName,
		i.LastName,
		i.Address,
		i.City,
		i.PostalCode,
		i.Country,
		i.PhoneNumber,
		i.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user_address not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (r *repo) Delete(ctx context.Context, id int) (*useraddress.UserAddress, error) {
	const deleteById = `
		DELETE FROM user_address WHERE id = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	var row userAddressRow
	if err := q.GetContext(ctx, &row, deleteById, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user_address not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}
