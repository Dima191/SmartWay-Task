package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *repository) EmployeeByID(ctx context.Context, employeeID int) (employee models.EmployeeExtended, err error) {
	q := `select employee_id,
				 first_name,
				 second_name,
				 employee_phone,
				 company_id,
				 company_name,
				 passport_id,
				 passport_type,
				 passport_number,
				 department_id,
				 department_name,
				 department_phone
			from extended_employee where employee_id=$1`

	if err = r.pool.QueryRow(ctx, q, employeeID).Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.SecondName,
		&employee.Phone,
		&employee.EmployeeCompany.ID,
		&employee.EmployeeCompany.Name,
		&employee.EmployeePassport.ID,
		&employee.EmployeePassport.Type,
		&employee.EmployeePassport.Number,
		&employee.EmployeeDepartment.ID,
		&employee.EmployeeDepartment.Name,
		&employee.EmployeeDepartment.Phone); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.EmployeeExtended{}, models.ErrNoData
		}
		r.logger.Warn("failed to get employee", slog.String("error", err.Error()))
		return models.EmployeeExtended{}, models.ErrInternal
	}

	return employee, nil

}
