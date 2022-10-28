package service

import (
	"context"
	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
)

type user struct {
	repo domain.UserRepository
}

func NewUser(repo domain.UserRepository) *user {
	return &user{repo: repo}
}

func (s *user) Create(ctx context.Context, user domain.User) (domain.Id, error) {
	return s.repo.Create(ctx, user)
}

func (s *user) Update(ctx context.Context, user domain.User) error {
	return s.repo.Update(ctx, user)
}

func (s *user) Delete(ctx context.Context, id domain.Id) error {
	return s.repo.Delete(ctx, id)
}

func (s *user) Read(ctx context.Context, id domain.Id) (domain.User, error) {
	return s.repo.Read(ctx, id)
}

func (s *user) ReadBalance(ctx context.Context, id domain.Id) (uint32, error) {
	return s.repo.ReadBalance(ctx, id)
}

func (s *user) ReserveMoney(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	err := s.repo.SubtractMoney(ctx, userId, amount)
	if err != nil {
		return err
	}

	return s.repo.ReserveMoney(ctx, userId, serviceId, orderId, amount)
}

func (s *user) AddMoney(ctx context.Context, id domain.Id, amount uint32) error {
	return s.repo.AddMoney(ctx, id, amount)
}
