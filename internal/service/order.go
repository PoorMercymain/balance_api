package service

import (
	"context"

	"github.com/PoorMercymain/balance_api/internal/domain"
)

type order struct {
	repo domain.OrderRepository
}

func NewOrder(repo domain.OrderRepository) *order {
	return &order{repo: repo}
}

func (s *order) Create(ctx context.Context, order domain.Order) (domain.Id, error) {
	return s.repo.Create(ctx, order)
}

func (s *order) Update(ctx context.Context, order domain.Order) error {
	return s.repo.Update(ctx, order)
}

func (s *order) Delete(ctx context.Context, id domain.Id) error {
	return s.repo.Delete(ctx, id)
}

func (s *order) Read(ctx context.Context, id domain.Id) (domain.Order, error) {
	return s.repo.Read(ctx, id)
}

func (s *order) AddService(ctx context.Context, orderId domain.Id, serviceId domain.Id) error {
	return s.repo.AddService(ctx, orderId, serviceId)
}
