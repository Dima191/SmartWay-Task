package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *repository) UpdateCompany(ctx context.Context, company models.Company) error {
	q := `update company set name = $1 where company_id = $2`

	if _, err := r.pool.Exec(ctx, q, company.Name, company.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				r.logger.Warn("failed to update company", slog.String("error", pgErr.Error()))
				return models.ErrAlreadyExists
			}
		}
	}

	return nil
}
