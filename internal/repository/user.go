package repository

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/PoorMercymain/balance_api/internal/domain"
	"github.com/PoorMercymain/balance_api/pkg/router"
	"github.com/jackc/pgx/v4"
	"os"
	"strconv"
	"time"
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
		`INSERT INTO reserve (user_id, order_id, service_id, money) VALUES ($1, $2, $3, $4)`,
		userId, orderId, serviceId, amount).Scan()

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return nil
}

func (r *user) AddMoney(ctx context.Context, id domain.Id, amount uint32, whoMade string, reason string) error {
	var uid domain.Id
	_, err := r.db.conn.Exec(ctx, `UPDATE "user" SET id = id, username = username, balance = balance + $1 WHERE id = $2`,
		amount, id)

	if err != nil {
		return err
	}

	err = r.db.conn.QueryRow(ctx,
		`INSERT INTO user_report (user_id, money, made_by, reason, transaction_date) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`,
		id, amount, whoMade, reason, time.Now()).Scan(&uid)

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return err
}

func (r *user) SubtractMoney(ctx context.Context, id domain.Id, amount uint32, whoMade string, reason string) error {
	var uid domain.Id
	balance, err := r.ReadBalance(ctx, id)
	if err != nil {
		return err
	}

	if balance < amount {
		err = newError("Balance editing error! User`s balance cannot be negative")
		return err
	}

	_, err = r.db.conn.Exec(ctx, `UPDATE "user" SET balance = balance - $1 WHERE id = $2`,
		amount, id)

	if err != nil {
		return err
	}

	err = r.db.conn.QueryRow(ctx,
		`INSERT INTO user_report (user_id, money, made_by, reason, transaction_date) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`,
		id, amount, whoMade, reason, time.Now()).Scan(&uid)

	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	return err
}

func (r *user) ReadServiceName(ctx context.Context, id domain.Id) (string, error) {
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

func (r *user) TransactionList(ctx context.Context, id domain.Id) ([]string, error) {
	pr := ctx.Value(router.KeyPagination).(*router.Pagination)

	page := make([]string, 0)
	var (
		pageStr         string
		money           uint32
		madeBy          string
		reason          string
		transactionDate time.Time
	)

	cmd := fmt.Sprintf(`SELECT money, made_by, reason, transaction_date 
						FROM user_report 
						WHERE user_id = $1 
						ORDER BY %s %s LIMIT $2 OFFSET $3`, pr.SortField, pr.SortDirection) //add validation

	row, err := r.db.conn.Query(ctx, cmd, id, pr.Limit, pr.Offset)
	if err != nil {
		return page, err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&money, &madeBy, &reason, &transactionDate)
		pageStr = fmt.Sprintf("money = %d, transaction made by %s, the reason is %s, transaction date is %v", money, madeBy, reason, transactionDate)
		page = append(page, pageStr)
	}

	return page, err
}

func (r *user) ReadAccountingReport(ctx context.Context, date domain.DateForReport) ([]domain.ReportContent, error) {
	reportContentSlice := make([]domain.ReportContent, 0)
	var reportContent domain.ReportContent

	row, err := r.db.conn.Query(ctx,
		`SELECT service_id, SUM(money) FROM accounting_report WHERE record_month = $1 AND record_year = $2 GROUP BY service_id`,
		date.Month, date.Year)

	if err != nil {
		return reportContentSlice, err
	}

	defer row.Close()

	for row.Next() {
		err = row.Scan(&reportContent.ServiceId, &reportContent.Total)
		reportContent.ServiceName, err = r.ReadServiceName(ctx, reportContent.ServiceId)
		if err != nil {
			return reportContentSlice, err
		}

		reportContentSlice = append(reportContentSlice, reportContent)
	}
	return reportContentSlice, err
}

func (r *user) MakeReport(ctx context.Context, date domain.DateForReport) (string, error) {
	reportContent, err := r.ReadAccountingReport(ctx, date)
	if err != nil {
		return "", err
	}

	reportContentStr := make([][]string, len(reportContent))
	for i, content := range reportContent {
		reportContentStr[i] = make([]string, 2)
		reportContentStr[i][0] = content.ServiceName
		reportContentStr[i][1] = strconv.Itoa(int(content.Total))
	}

	filename := fmt.Sprintf("report_%s.csv", time.Now().Format("2006-01-02"))
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(reportContentStr)
	return fmt.Sprintf("%s/%s", "localhost:8000", filename), err
}
