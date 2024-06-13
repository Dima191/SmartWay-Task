package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *repository) AddCompany(ctx context.Context, company models.Company) (companyID int, err error) {
	q := `insert into company(name) values($1) returning company_id;`

	if err = r.pool.QueryRow(ctx, q, company.Name).Scan(&companyID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				r.logger.Warn("failed to add company. company already exists")
				return 0, models.ErrAlreadyExists
			}
		}
		r.logger.Warn("failed to add company", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	return companyID, nil
}
