package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *repository) AddDepartment(ctx context.Context, department models.Department) (departmentID int, err error) {
	q := `insert into department (company_id, name, phone) values($1, $2, $3) returning department_id`
	if err = r.pool.QueryRow(ctx, q, department.CompanyID, department.Name, department.Phone).Scan(&departmentID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				r.logger.Warn("failed to add department. department already exists", slog.String("error", pgErr.Error()))
				return 0, models.ErrAlreadyExists
			case pgerrcode.ForeignKeyViolation:
				r.logger.Warn("failed to add department to unknown company", slog.String("error", pgErr.Error()))
				return 0, models.ErrUnknownReference
			}
		}
		r.logger.Warn("failed to add department", slog.String("error", err.Error()))
		return 0, models.ErrInternal
	}

	return departmentID, nil
}
