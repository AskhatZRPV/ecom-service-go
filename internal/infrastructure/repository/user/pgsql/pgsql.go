package pgsql

import (
	"context"
	"database/sql"
	"ecomsvc/internal/domain/user"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) user.Repository {
	return &repo{db}
}

func (i *repo) Save(ctx context.Context, u *user.User) (int, error) {
	const insertUserQuery = `
		INSERT INTO user (username, password, role, created_at) VALUES($1, $2, $3, $4);
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	qRes, err := q.ExecContext(ctx, insertUserQuery, u.Username, u.Password, u.Role, u.CreatedAt)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert new user recotd")
	}

	lastId, err := qRes.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "cannot get last inserted id")
	}

	return int(lastId), nil
}

func (i *repo) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	const selectUserByUsernameQuery = `
		SELECT * FROM user WHERE username = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	var row userRow
	err := q.GetContext(ctx, &row, selectUserByUsernameQuery, username)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (i *repo) UpdateLastLoginAttempt(ctx context.Context, id user.ID, updatedDate time.Time) error {
	const updateUserLastLoginAttempForIDQuery = `
		UPDATE user SET last_login_attempt = $1 WHERE id = $2
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	r, err := q.ExecContext(ctx, updateUserLastLoginAttempForIDQuery, updatedDate, id)
	if err != nil {
		return errors.Wrap(err, "failed to update users table row")
	}

	if rowsAffected, err := r.RowsAffected(); err != nil {
		return errors.Wrap(err, "failed to get rows affected number")
	} else if rowsAffected == 0 {
		return errors.Wrap(user.ErrUserNotFound, "user not found. no records updated")
	}

	return nil
}
