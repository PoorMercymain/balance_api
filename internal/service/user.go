package service

import (
	"context"
	"fmt"
	"github.com/PoorMercymain/balance_api/internal/domain"
	"io/ioutil"
)

type user struct {
	repo domain.UserRepository
}

func NewUser(repo domain.UserRepository) *user {
	return &user{repo: repo}
}

func newError(text string) error {
	return &incorrectFilename{text}
}

type incorrectFilename struct {
	s string
}

func (e *incorrectFilename) Error() string {
	return e.s
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

func (s *user) ReadServiceName(ctx context.Context, id domain.Id) (string, error) {
	return s.repo.ReadServiceName(ctx, id)
}

func (s *user) ReserveMoney(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32, whoMade string) error {
	serviceName, err := s.repo.ReadServiceName(ctx, serviceId)
	if err != nil {
		return err
	}

	err = s.repo.SubtractMoney(ctx, userId, amount, whoMade, fmt.Sprintf("money reserved for order %d service %s", orderId, serviceName))
	if err != nil {
		return err
	}

	return s.repo.ReserveMoney(ctx, userId, serviceId, orderId, amount)
}

func (s *user) AddMoney(ctx context.Context, id domain.Id, amount uint32, whoMade string, reason string) error {
	return s.repo.AddMoney(ctx, id, amount, whoMade, reason)
}

func (s *user) TransactionList(ctx context.Context, id domain.Id) ([]string, error) {
	return s.repo.TransactionList(ctx, id)
}

func (s *user) MakeReport(ctx context.Context, data domain.DateForReport) (string, error) {
	return s.repo.MakeReport(ctx, data)
}

func (s *user) GetReport(ctx context.Context, filename string) ([]byte, error) {
	result := make([]byte, 0)
	if !(len(filename) > 3 || filename[len(filename)-3:] == "csv") {
		return result, newError("Incorrect filename. Expected a csv file")
	}

	result, err := ioutil.ReadFile(filename)
	if err != nil {
		return result, err
	}

	return result, err
}
