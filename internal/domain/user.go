package domain

import (
	"context"
)

type User struct {
	UserId   Id     `json:"user_id"`
	Username string `json:"username"`
	Balance  uint32 `json:"balance"`
}

type DateForReport struct {
	Year  uint32 `json:"year"`
	Month uint32 `json:"month"`
}

type ReportContent struct {
	ServiceId   Id
	ServiceName string
	Total       uint32
}

type UserRepository interface {
	Create(ctx context.Context, user User) (Id, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (User, error)
	ReadBalance(ctx context.Context, id Id) (uint32, error)
	ReserveMoney(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32) error
	AddMoney(ctx context.Context, id Id, amount uint32, whoMade string, reason string) error
	SubtractMoney(ctx context.Context, id Id, amount uint32, whoMade string, reason string) error
	ReadServiceName(ctx context.Context, id Id) (string, error)
	TransactionList(ctx context.Context, id Id) ([]string, error)
	MakeReport(ctx context.Context, date DateForReport) (string, error)
	ReadAccountingReport(ctx context.Context, date DateForReport) ([]ReportContent, error)
}

type UserService interface {
	Create(ctx context.Context, user User) (Id, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id Id) error
	Read(ctx context.Context, id Id) (User, error)
	ReadBalance(ctx context.Context, id Id) (uint32, error)
	ReserveMoney(ctx context.Context, userId Id, serviceId Id, orderId Id, amount uint32, whoMade string) error
	AddMoney(ctx context.Context, id Id, amount uint32, whoMade string, reason string) error
	ReadServiceName(ctx context.Context, id Id) (string, error)
	TransactionList(ctx context.Context, id Id) ([]string, error)
	MakeReport(ctx context.Context, date DateForReport) (string, error)
	GetReport(ctx context.Context, filename string) ([]byte, error)
}
