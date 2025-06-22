package apikey

import (
	"context"
)

type Service interface {
	ListAPIKeys(ctx context.Context) ([]*Tenant, error)
	CreateAPIKey(ctx context.Context, req CreateAPIKeyRequest) (*Tenant, error)
	DeleteAPIKey(ctx context.Context, id string) error
}

type Store interface {
	GetAll(ctx context.Context) ([]*Tenant, error)
	Insert(ctx context.Context, t *Tenant) error
	DeleteByID(ctx context.Context, id string) error
}
