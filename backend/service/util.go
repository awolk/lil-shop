package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
)

type txClientKey struct{}

func (s *Service) client(ctx context.Context) *ent.Client {
	txClient := ctx.Value(txClientKey{})
	if txClient != nil {
		return txClient.(*ent.Client)
	}

	return s.globalClient
}

func (s *Service) withTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := s.globalClient.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed starting transaction: %w", err)
	}

	ctxWithTx := context.WithValue(ctx, txClientKey{}, tx.Client())

	if err := fn(ctxWithTx); err != nil {
		err := fmt.Errorf("failed executing transaction: %w", err)

		rerr := tx.Rollback()
		if rerr != nil {
			return fmt.Errorf("failed rolling back transaction: %v\n%w", rerr, err)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed commiting transaction: %w", err)
	}

	return nil
}
