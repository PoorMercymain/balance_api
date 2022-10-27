package repository

import (
	"context"

	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/domain"
)

type user struct {
	db *db
}

func NewUser(db *db) *user {
	return &user{db: db}
}

func newError(text string) error {
	return &incorrectBalance{text}
}

type incorrectBalance struct {
	s string
}

func (e *incorrectBalance) Error() string {
	return e.s
}

func (r *user) Create(ctx context.Context, user domain.User) (domain.Id, error) {
	var id domain.Id

	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO "user" (id, username, balance) VALUES ($1, $2, $3) RETURNING id`,
		user.UserId, user.Username, user.Balance).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *user) Update(ctx context.Context, user domain.User) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE "user" SET id = $1, username = $2, balance = $3 WHERE id = $1`,
		user.UserId, user.Username, user.Balance)

	if err != nil {
		return err
	}

	return err
}

func (r *user) Delete(ctx context.Context, id domain.Id) error {
	_, err := r.db.conn.Exec(ctx, `DELETE FROM "user" WHERE id=$1`, id)

	if err != nil {
		return err
	}

	return err
}

func (r *user) Read(ctx context.Context, id domain.Id) (domain.User, error) {
	var user domain.User

	row, err := r.db.conn.Query(ctx,
		`SELECT id, username, balance FROM "user" WHERE id = $1`, id)

	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&user.UserId, &user.Username, &user.Balance)
	}

	return user, err
}

func (r *user) ReadBalance(ctx context.Context, id domain.Id) (uint32, error) {
	var balance uint32

	row, err := r.db.conn.Query(ctx,
		`SELECT balance FROM "user" WHERE id = $1`, id)

	if err != nil {
		return balance, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&balance)
	}

	return balance, err
}

func (r *user) ReserveMoney(ctx context.Context, userId domain.Id, serviceId domain.Id, orderId domain.Id, amount uint32) error {
	err := r.db.conn.QueryRow(ctx,
		`INSERT INTO reserve (user_id, order_id, money) VALUES ($1, $2, $3)`,
		userId, orderId, amount).Scan()

	if err != nil {
		return err
	}

	return err
}

func (r *user) AddMoney(ctx context.Context, id domain.Id, amount uint32) error {
	_, err := r.db.conn.Exec(ctx, `UPDATE "user" SET id = id, username = username, balance = balance + $1 WHERE id = $2`,
		amount, id)

	if err != nil {
		return err
	}

	return err
}

func (r *user) SubtractMoney(ctx context.Context, id domain.Id, amount uint32) error {
	balance, err := r.ReadBalance(ctx, id)
	if err != nil {
		return err
	}

	if balance < amount {
		err = newError("Balance editing error! User`s balance cannot be negative")
		return err
	}

	_, err = r.db.conn.Exec(ctx, `UPDATE "user" SET id = id, username = username, balance = balance - $1 WHERE id = $2`,
		amount, id)

	if err != nil {
		return err
	}

	return err
}
