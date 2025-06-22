package apikey

import (
	"context"
	"errors"
	"fmt"

	"github.com/cappit/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	getAllQuery = `SELECT id, name, api_key, api_secret, base_url, created_at, updated_at FROM tenants`
	insertQuery = `INSERT INTO tenants (id, name, api_key, api_secret, base_url, created_at, updated_at)` +
		` VALUES ($1, $2, $3, $4, $5, $6, $7)`
	deleteQuery = `DELETE FROM tenants WHERE id = $1`
)

type store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &store{db: db}
}

func (s *store) GetAll(ctx context.Context) ([]*Tenant, error) {
	rows, err := s.db.Query(ctx, getAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []*Tenant
	for rows.Next() {
		var t Tenant
		if err = rows.Scan(&t.ID, &t.Name, &t.APIKey, &t.APISecret, &t.BaseURL, &t.CreatedAt,
			&t.UpdatedAt); err != nil {
			return nil, err
		}
		tenants = append(tenants, &t)
	}
	return tenants, nil
}

func (s *store) Insert(ctx context.Context, t *Tenant) error {
	_, err := s.db.Exec(ctx, insertQuery,
		t.ID, t.Name, t.APIKey, t.APISecret, t.BaseURL, t.CreatedAt, t.UpdatedAt)

	if err != nil {
		logger.Error("DB error: failed to insert tenant", zap.Error(err))
		return fmt.Errorf("DB error: failed to insert tenant: %w", err)
	}

	return nil
}

func (s *store) DeleteByID(ctx context.Context, id string) error {
	rows, err := s.db.Exec(ctx, deleteQuery, id)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
