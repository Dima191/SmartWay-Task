package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *repository) UpdateDepartment(ctx context.Context, department models.Department) error {
	q := `update department SET company_id = $1, name = $2, phone = $3 where department_id = $4`

	if _, err := r.pool.Exec(ctx, q, department.CompanyID, department.Name, department.Phone, department.ID); err != nil {
		r.logger.Warn("failed to update department", slog.String("error", err.Error()))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return models.ErrAlreadyExists
			}

			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return models.ErrUnknownCompanyReference
			}
		}

		return models.ErrInternal
	}

	return nil
}
