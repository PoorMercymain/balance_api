package repository

import (
	"context"

	"github.com/PoorMercymain/balance_api/internal/domain"
)

type order struct {
	db *db
}

func NewOrder(db *db) *order {
	return &order{db: db}
}

func (r *order) Create(ctx context.Context, order domain.Order) (domain.Id, error) {
	var id domain.Id

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO "order" (id, user_id) VALUES ($1, $2) RETURNING id`,
		order.OrderId, order.UserId).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *order) Update(ctx context.Context, order domain.Order) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE "order" SET id = $1, user_id = $2 WHERE id = $1`, order.OrderId, order.UserId)

	if err != nil {
		return err
	}

	return err
}

func (r *order) Delete(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM "order" WHERE id=$1`, id)

	if err != nil {
		return err
	}

	return err
}

func (r *order) Read(ctx context.Context, id domain.Id) (domain.Order, error) {
	var order domain.Order

	row, err := r.db.conn.Query(ctx,
		`SELECT id, user_id FROM "order" WHERE id = $1`, id)

	if err != nil {
		return order, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&order.OrderId, &order.UserId)
	}

	return order, err
}

func (r *order) AddService(ctx context.Context, orderId domain.Id, serviceId domain.Id) error {
	var id domain.Id

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO order_service (order_id, service_id) VALUES ($1, $2)`,
		orderId, serviceId).Scan(&id)

	if err != nil {
		return err
	}

	return err
}
