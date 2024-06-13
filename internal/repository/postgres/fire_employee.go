package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *repository) FireEmployee(ctx context.Context, employeeID int) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	})
	defer tx.Commit(ctx)
	if err != nil {
		r.logger.Warn("failed to fire employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	q := "delete from employee where employee_id = $1 returning passport_id"

	var passportID int

	if err = tx.QueryRow(ctx, q, employeeID).Scan(&passportID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ErrNoData
		}

		r.logger.Warn("failed to fire employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	q = "delete from passport where passport_id = $1"

	if _, err = tx.Exec(ctx, q, passportID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ErrNoData
		}
		r.logger.Warn("failed to fire employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	if err = tx.Commit(ctx); err != nil {
		r.logger.Warn("no employee with requested id", slog.Int("employee_id", employeeID))
		return models.ErrInternal
	}

	return nil
}
