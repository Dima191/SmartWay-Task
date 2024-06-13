package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"strings"
)

func (r *repository) UpdateEmployee(ctx context.Context, employee models.EmployeeExtended) error {
	tx, err := r.pool.Begin(ctx)
	defer tx.Commit(ctx)
	if err != nil {
		r.logger.Warn("failed to update employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	q := `update employee set
							first_name = $1,
							second_name = $2,
							phone = $3
		  where employee_id = $4 returning passport_id`

	var passportID int

	if err = tx.QueryRow(ctx, q,
		employee.FirstName,
		employee.SecondName,
		employee.Phone,
		employee.ID).Scan(&passportID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return models.ErrAlreadyExists
			}
		}

		r.logger.Warn("failed to update employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	switch {
	case employee.EmployeeCompany.ID != -1 && employee.EmployeeDepartment.ID != -1:
		q = `update employee set
								company_id = $1,
								department_id = $2
			  where employee_id = $3 returning passport_id`

		if err = tx.QueryRow(ctx, q,
			employee.EmployeeCompany.ID,
			employee.EmployeeDepartment.ID,
			employee.ID).Scan(&passportID); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == pgerrcode.ForeignKeyViolation {
					if strings.Contains(pgErr.ConstraintName, "company") {
						r.logger.Warn("failed to add employee to unknown company", slog.String("error", pgErr.Error()))
						return models.ErrUnknownCompanyReference
					}

					r.logger.Warn("failed to add employee to unknown department", slog.String("error", pgErr.Error()))
					return models.ErrUnknownDepartmentReference
				}
			}

			r.logger.Warn("failed to update employee", slog.String("error", err.Error()))
			return models.ErrInternal
		}
	case employee.EmployeeCompany.ID == -1 && employee.EmployeeDepartment.ID == -1:
	default:
		return models.ErrCompanyDepProvide
	}

	q = `update passport set
							type = $1,
							number = $2
	     where passport_id = $3`

	if _, err = tx.Exec(ctx, q,
		employee.EmployeePassport.Type,
		employee.EmployeePassport.Number,
		passportID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				r.logger.Warn("failed to update employee", slog.String("error", pgErr.Error()))
				return models.ErrAlreadyExists
			}
		}
		r.logger.Warn("failed to update employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	if err = tx.Commit(ctx); err != nil {
		r.logger.Warn("failed to update employee", slog.String("error", err.Error()))
		return models.ErrInternal
	}

	return nil
}
