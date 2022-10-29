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
	Update(ctx context.Context, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	DeleteByOrderIdAndServiceId(ctx context.Context, orderId Id, serviceId Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
	DeleteServiceFromOrderService(ctx context.Context, serviceId Id, orderId Id) error
	OrderExists(ctx context.Context, orderId Id) (bool, error)
	DeleteOrder(ctx context.Context, id Id) error
	ReturnMoneyFromReserve(ctx context.Context, userId Id, amount uint32, whoMade string, reason string) error
	ReadServiceName(ctx context.Context, id Id) (string, error)
}

type ReserveService interface {
	Create(ctx context.Context, reserve Reserve) (Id, error)
	Update(ctx context.Context, reserve Reserve) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (Reserve, error)
	ApproveRevenue(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
	ReturnMoneyFromReserve(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32, whoMade string) error
}
