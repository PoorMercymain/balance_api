package domain

import (
	"context"
)

type Order struct {
	OrderId Id `json:"order_id"`
	UserId  Id `json:"user_id"`
}

type OrderRepository interface {
	Create(ctx context.Context, order Order) (Id, error)
	Update(ctx context.Context, order Order) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Order, error)
	AddService(ctx context.Context, orderId Id, serviceId Id) error
}

type OrderService interface {
	Create(ctx context.Context, order Order) (Id, error)
	Update(ctx context.Context, order Order) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Order, error)
	AddService(ctx context.Context, orderId Id, serviceId Id) error
}
