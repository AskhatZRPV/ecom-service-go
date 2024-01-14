package pgsql

import (
	"context"
	"ecomsvc/internal/domain/orderitem"
	"ecomsvc/internal/infrastructure/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) orderitem.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, s *orderitem.OrderItem) error {
	const insertSessionQuery = `
		INSERT INTO order_items (id, title, description, price, category_id)
		VALUES($1, $2, $3, $4, $5);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, insertSessionQuery, s.ID, s.Title, s.Description, s.Price, s.CategoryId)
	if err != nil {
		return errors.Wrap(err, "failed to create order_item record")
	}

	return nil
}
