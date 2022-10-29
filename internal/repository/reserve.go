package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"

	"github.com/PoorMercymain/balance_api/internal/domain"
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
		`INSERT INTO reserve (id, user_id, order_id, service_id, money) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		reserve.ReserveId, reserve.UserId, reserve.OrderId, reserve.ServiceId, reserve.Money).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *reserve) Update(ctx context.Context, reserve domain.Reserve) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE reserve SET id = $1, user_id = $2, order_id = $3, service_id = $4, money = $5 WHERE id = $1`,
		reserve.ReserveId, reserve.UserId, reserve.OrderId, reserve.ServiceId, reserve.Money)

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

func (r *reserve) Read(ctx context.Context, id domain.Id) (domain.Reserve, error) {
	var reserve domain.Reserve

	row, err := r.db.conn.Query(ctx,
		`SELECT id, user_id, order_id, service_id, money FROM reserve WHERE id=$1`,
		id)

	if err != nil {
		return reserve, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&reserve.ReserveId, &reserve.UserId, &reserve.OrderId, &reserve.ServiceId, &reserve.Money)
	}

	return reserve, err
}

func (r *reserve) DeleteByOrderIdAndServiceId(ctx context.Context, orderId domain.Id, serviceId domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM reserve WHERE order_id=$1 AND service_id=$2`, orderId, serviceId)

	if err != nil {
		return err
	}

	return err
}

func (r *reserve) DeleteOrder(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM "order" WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return err
}

func (r *reserve) DeleteServiceFromOrderService(ctx context.Context, serviceId domain.Id, orderId domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM order_service WHERE order_id=$1 AND service_id=$2`, orderId, serviceId)
	if err != nil {
		return err
	}

	return err
}

func (r *reserve) OrderExists(ctx context.Context, orderId domain.Id) (bool, error) {
	var counter int
	err := r.db.conn.QueryRow(ctx,
		`SELECT COUNT(*) order_id FROM order_service WHERE order_id=$1`, orderId).Scan(&counter)

	if err != nil || counter == 0 {
		return false, err
	}

	return true, err
}

func (r *reserve) ApproveRevenue(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	year := time.Now().Year()
	month := int(time.Now().Month())

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO accounting_report (user_id, service_id, money, record_year, record_month) VALUES ($1, $2, $3, $4, $5)`,
		userId, serviceId, amount, year, month).Scan()

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return nil
}

func (r *reserve) ReturnMoneyFromReserve(ctx context.Context, userId domain.Id, amount uint32, whoMade string, reason string) error {
	var uid domain.Id
	_, err := r.db.conn.Exec(ctx, `UPDATE "user" SET balance = balance  + $1 WHERE id = $2`,
		amount, userId)

	if err != nil {
		return err
	}

	err = r.db.conn.QueryRow(ctx,
		`INSERT INTO user_report (user_id, money, made_by, reason) VALUES ($1, $2, $3, $4) RETURNING user_id`,
		userId, amount, whoMade, reason).Scan(&uid)

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return err
}

func (r *reserve) ReadServiceName(ctx context.Context, id domain.Id) (string, error) {
	var serviceName string

	row, err := r.db.conn.Query(ctx,
		`SELECT service_name FROM service WHERE id = $1`, id)

	if err != nil {
		return serviceName, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&serviceName)
	}

	return serviceName, err
}
