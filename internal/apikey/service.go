package apikey

import (
	"context"
	"time"

	"github.com/cappit/internal/util"
	"github.com/google/uuid"
)

type service struct {
	store Store
}

func NewService(store Store) Service {
	return &service{store: store}
}

func (s *service) ListAPIKeys(ctx context.Context) ([]*Tenant, error) {
	return s.store.GetAll(ctx)
}

func (s *service) CreateAPIKey(ctx context.Context, req CreateAPIKeyRequest) (*Tenant, error) {
	now := time.Now().UTC()

	t := &Tenant{
		ID:        uuid.New(),
		Name:      req.Name,
		APIKey:    uuid.New().String(),
		APISecret: uuid.New().String(),
		BaseURL:   req.BaseURL,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := ValidateTenantInput(t.Name, t.BaseURL, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err = s.store.Insert(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) DeleteAPIKey(ctx context.Context, id string) error {
	err := util.ValidateUUID(id)
	if err != nil {
		return err
	}

	return s.store.DeleteByID(ctx, id)
}
