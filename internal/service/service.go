package service

import (
	"context"

	"github.com/PoorMercymain/balance_api/internal/domain"
)

type service struct {
	repo domain.ServiceRepository
}

func NewService(repo domain.ServiceRepository) *service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, service domain.Service) (domain.Id, error) {
	return s.repo.Create(ctx, service)
}

func (s *service) Update(ctx context.Context, service domain.Service) error {
	return s.repo.Update(ctx, service)
}

func (s *service) Delete(ctx context.Context, id domain.Id) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) Read(ctx context.Context, id domain.Id) (domain.Service, error) {
	return s.repo.Read(ctx, id)
}
