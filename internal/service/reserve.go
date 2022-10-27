package service

import (
	"context"

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

func (s *reserve) Update(ctx context.Context, id domain.Id, reserve domain.Reserve) error {
	return s.repo.Update(ctx, id, reserve)
}

func (s *reserve) Delete(ctx context.Context, id domain.Id) error {
	return s.repo.Delete(ctx, id)
}

func (s *reserve) Read(ctx context.Context, id domain.Id) (domain.Reserve, error) {
	return s.repo.Read(ctx, id)
}

func (s *reserve) ApproveRevenue(ctx context.Context, id domain.Id) error {
	return s.repo.ApproveRevenue(ctx, id)
}
