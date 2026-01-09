package payment

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zerogate/payment/internal/redis"
)

type Service struct {
	repo   *Repository
	locker *redis.Locker
}

func NewService(repo *Repository, locker *redis.Locker) *Service {
	return &Service{repo: repo, locker: locker}
}

func (s *Service) ProcessPayment(ctx context.Context, idempotencyKey string, amount float64) (*Transaction, error) {
	if idempotencyKey == "" {
		return nil, errors.New("idempotency key is required")
	}

	locked, err := s.locker.AcquireLock(ctx, idempotencyKey, 30*time.Second)
	if err != nil {
		return nil, err
	}
	if !locked {
		return nil, errors.New("request is currently being processed")
	}
	defer s.locker.ReleaseLock(context.Background(), idempotencyKey)

	existingTx, err := s.repo.GetByIdempotencyKey(ctx, idempotencyKey)
	if err != nil {
		return nil, err
	}
	if existingTx != nil {
		return existingTx, nil
	}

	newTx := &Transaction{
		ID:             uuid.New().String(),
		IdempotencyKey: idempotencyKey,
		Amount:         amount,
		Status:         "completed",
	}

	if err := s.repo.Save(ctx, newTx); err != nil {
		return nil, err
	}

	return newTx, nil
}
