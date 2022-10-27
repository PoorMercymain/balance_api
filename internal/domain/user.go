package domain

import (
	"context"
)

type User struct {
	UserId   Id     `json:"user_id"`
	Username string `json:"username"`
	Balance  uint32 `json:"balance"`
}

type UserRepository interface {
	Create(ctx context.Context, user User) (Id, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (User, error)
	ReadBalance(ctx context.Context, id Id) (uint32, error)
	ReserveMoney(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
	AddMoney(ctx context.Context, id Id, amount uint32) error
	SubtractMoney(ctx context.Context, id Id, amount uint32) error
}

type UserService interface {
	Create(ctx context.Context, user User) (Id, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (User, error)
	ReadBalance(ctx context.Context, id Id) (uint32, error)
	ReserveMoney(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
	AddMoney(ctx context.Context, id Id, amount uint32) error
}
