package domain

import (
	"context"
)

type Reserve struct {
	ReserveId Id     `json:"reserve_id"`
	UserId    Id     `json:"user_id"`
	OrderId   Id     `json:"order_id"`
	Money     uint32 `json:"money"`
}

type ReserveRepository interface {
	Create(ctx context.Context, reserve Reserve) (Id, error)
	Update(ctx context.Context, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	DeleteByOrderId(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
}

type ReserveService interface {
	Create(ctx context.Context, reserve Reserve) (Id, error)
	Update(ctx context.Context, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
}
