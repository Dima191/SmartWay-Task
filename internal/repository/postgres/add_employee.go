package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *repository) AddEmployee(ctx context.Context, employee models.EmployeeBase) (employeeID int, err error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	})
	defer tx.Commit(ctx)
	if err != nil {
		r.logger.Warn("failed to add employee", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	q := `insert into passport(type, number) values ($1,$2) RETURNING passport_id`
	var passportID int

	if err = tx.QueryRow(ctx, q,
		employee.EmployeePassport.Type,
		employee.EmployeePassport.Number,
	).Scan(&passportID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				r.logger.Warn("failed to add passport. passport already exists", slog.String("error", pgErr.Error()))
				return 0, models.ErrPassportAlreadyExists
			}
		}
		r.logger.Warn("failed to add passport", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	q = `insert into employee(employee_id,
										first_name,
										second_name,
										phone,
										passport_id) values ($1,$2,$3,$4,$5) RETURNING employee_id`

	if err = tx.QueryRow(ctx, q,
		employee.ID,
		employee.FirstName,
		employee.SecondName,
		employee.Phone,
		passportID,
	).Scan(&employeeID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				r.logger.Warn("failed to add employee. employee already exists", slog.String("error", pgErr.Error()))
				return 0, models.ErrEmployeeAlreadyExists
			}
		}
		r.logger.Warn("failed to add employee", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	if err = tx.Commit(ctx); err != nil {
		r.logger.Warn("failed to add employee", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	return employeeID, nil
}
