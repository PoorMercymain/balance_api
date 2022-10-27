package domain

import (
	"context"
)

type Service struct {
	ServiceId Id     `json:"service_id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
}

type ServiceRepository interface {
	Create(ctx context.Context, service Service) (Id, error)
	Update(ctx context.Context, id Id, service Service) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Service, error)
}

type ServiceService interface {
	Create(ctx context.Context, service Service) (Id, error)
	Update(ctx context.Context, id Id, service Service) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Service, error)
}
