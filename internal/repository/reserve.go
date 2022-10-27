package repository

import (
	"context"

	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
)

type reserve struct {
	db *db
}

func NewReserve(db *db) *reserve {
	return &reserve{db: db}
}

func (r *reserve) Create(ctx context.Context, reserve domain.Reserve) (domain.Id, error) {
	var id domain.Id

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO reserve (id, user_id, order_id, money) VALUES ($1, $2, $3, $4) RETURNING id`,
		reserve.ReserveId, reserve.UserId, reserve.OrderId, reserve.Money).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *reserve) Update(ctx context.Context, reserve domain.Reserve) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE reserve SET id = $1, user_id = $2, order_id = $3, money = $4 WHERE id = $1`,
		reserve.ReserveId, reserve.UserId, reserve.OrderId, reserve.Money)

	if err != nil {
		return err
	}

	return err
}

func (r *reserve) Delete(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM reserve WHERE id=$1`, id)

	if err != nil {
		return err
	}

	return err
}

func (r *reserve) DeleteByOrderId(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM reserve WHERE order_id=$1`, id)

	if err != nil {
		return err
	}

	return err
}

func (r *reserve) Read(ctx context.Context, id domain.Id) (domain.Reserve, error) {
	var reserve domain.Reserve

	row, err := r.db.conn.Query(ctx,
		`SELECT id, user_id, order_id, money FROM reserve WHERE id = $1`, id)

	if err != nil {
		return reserve, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&reserve.ReserveId, &reserve.UserId, &reserve.OrderId, &reserve.Money)
	}

	return reserve, err
}

func (r *reserve) ApproveRevenue(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO report (user_id, service_id, revenue) VALUES ($1, $2, $3)`,
		userId, serviceId, amount).Scan()

	if err != nil {
		return err
	}

	return err
}
