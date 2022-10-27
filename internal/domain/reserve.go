package domain

import (
	"context"
)

type Reserve struct {
	ReserveId Id     `json:"reserve_id"`
	UserId    Id     `json:"user_id"`
	ServiceId Id     `json:"service_id"`
	OrderId   Id     `json:"order_id"`
	Money     uint32 `json:"money"`
}

type ReserveRepository interface {
	Create(ctx context.Context, reserve Reserve) (Id, error)
	Update(ctx context.Context, id Id, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, id Id) error
}

type ReserveService interface {
	Create(ctx context.Context, reserve Reserve) (Id, error)
	Update(ctx context.Context, id Id, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, id Id) error
}
