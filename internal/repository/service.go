package repository

import (
	"context"
	"github.com/PoorMercymain/balance_api/internal/domain"
)

type service struct {
	db *db
}

func NewService(db *db) *service {
	return &service{db: db}
}

func (r *service) Create(ctx context.Context, service domain.Service) (domain.Id, error) {
	var id domain.Id

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO service (id, service_name, price) VALUES ($1, $2, $3) RETURNING id`,
		service.ServiceId, service.Name, service.Price).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *service) Update(ctx context.Context, service domain.Service) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE service SET id = $1, service_name = $2, price = $3 WHERE id = $1`,
		service.ServiceId, service.Name, service.Price)

	if err != nil {
		return err
	}

	return err
}

func (r *service) Delete(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM service WHERE id=$1`, id)

	if err != nil {
		return err
	}

	return err
}

func (r *service) Read(ctx context.Context, id domain.Id) (domain.Service, error) {
	var service domain.Service

	row, err := r.db.conn.Query(ctx,
		`SELECT id, service_name, price FROM service WHERE id = $1`, id)

	if err != nil {
		return service, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&service.ServiceId, &service.Name, &service.Price)
	}

	return service, err
}
