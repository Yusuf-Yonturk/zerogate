package payment

import (
	"context"
	"database/sql"
	"errors"
)

type Transaction struct {
	ID             string
	IdempotencyKey string
	Amount         float64
	Status         string
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByIdempotencyKey(ctx context.Context, key string) (*Transaction, error) {
	query := `SELECT id, idempotency_key, amount, status FROM transactions WHERE idempotency_key = $1`
	var t Transaction
	err := r.db.QueryRowContext(ctx, query, key).Scan(&t.ID, &t.IdempotencyKey, &t.Amount, &t.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *Repository) Save(ctx context.Context, t *Transaction) error {
	query := `INSERT INTO transactions (id, idempotency_key, amount, status) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, t.ID, t.IdempotencyKey, t.Amount, t.Status)
	return err
}
