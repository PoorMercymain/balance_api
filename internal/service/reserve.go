package service

import (
	"context"
	"fmt"

	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
)

type reserve struct {
	repo domain.ReserveRepository
}

func NewReserve(repo domain.ReserveRepository) *reserve {
	return &reserve{repo: repo}
}

func (s *reserve) Create(ctx context.Context, reserve domain.Reserve) (domain.Id, error) {
	return s.repo.Create(ctx, reserve)
}

func (s *reserve) Update(ctx context.Context, reserve domain.Reserve) error {
	return s.repo.Update(ctx, reserve)
}

func (s *reserve) Delete(ctx context.Context, id domain.Id) error {
	return s.repo.Delete(ctx, id)
}

func (s *reserve) Read(ctx context.Context, id domain.Id) (domain.Reserve, error) {
	return s.repo.Read(ctx, id)
}

func (s *reserve) ApproveRevenue(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	err := s.repo.DeleteServiceFromOrderService(ctx, serviceId, orderId)
	if err != nil {
		return err
	}
	fmt.Println("service from order-service")

	err = s.repo.DeleteByOrderIdAndServiceId(ctx, orderId, serviceId)
	if err != nil {
		return err
	}

	exists, err := s.repo.OrderExists(ctx, orderId)

	if err != nil {
		return err
	}
	fmt.Println("order exists")
	fmt.Println(exists)

	if !exists {
		fmt.Println("удаляем ордер")
		s.repo.DeleteOrder(ctx, orderId)
	}

	return s.repo.ApproveRevenue(ctx, userId, serviceId, orderId, amount)
}

func (s *reserve) ReturnMoneyFromReserve(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	err := s.repo.DeleteByOrderIdAndServiceId(ctx, orderId, serviceId)
	if err != nil {
		return err
	}

	return s.repo.ReturnMoneyFromReserve(ctx, userId, amount)
}
