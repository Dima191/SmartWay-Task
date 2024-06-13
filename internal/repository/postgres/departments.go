package postgresrep

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"log/slog"
)

func (r *repository) Departments(ctx context.Context, companyID int) (departments []models.Department, err error) {
	q := `select department_id, company_id, name, phone from department where company_id = $1`
	rows, err := r.pool.Query(ctx, q, companyID)
	if err != nil {
		r.logger.Warn("failed to get list of departments", slog.String("error", err.Error()))
		return nil, models.ErrInternal
	}

	for rows.Next() {
		department := models.Department{}
		if err = rows.Scan(&department.ID, &department.CompanyID, &department.Name, &department.Phone); err != nil {
			r.logger.Warn("failed to get list of departments", slog.String("error", err.Error()))
			return nil, models.ErrInternal
		}

		departments = append(departments, department)
	}

	return departments, nil
}
