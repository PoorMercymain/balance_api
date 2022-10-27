package domain

import (
	"context"
)

type Order struct {
	OrderId    Id     `json:"order_id"`
	UserId     string `json:"user_id"`
	ServiceIds []Id   `json:"service_ids"`
	Total      uint32 `json:"total"`
}

type OrderRepository interface {
	Create(ctx context.Context, order Order) (Id, error)
	Update(ctx context.Context, id Id, order Order) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Order, error)
	AddService(ctx context.Context, orderId Id, serviceId Id) error
}

type OrderService interface {
	Create(ctx context.Context, order Order) (Id, error)
	Update(ctx context.Context, id Id, order Order) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Order, error)
	AddService(ctx context.Context, orderId Id, serviceId Id) error
}
