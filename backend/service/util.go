package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
)

func (s *Service) withTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed starting transaction: %w", err)
	}

	if err := fn(tx); err != nil {
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
