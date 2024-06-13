package postgresrep

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"log/slog"
)

func (r *repository) CompanyEmployees(ctx context.Context, companyID int) (employees []models.EmployeeExtended, err error) {
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
			from extended_employee 
			where company_id=$1`

	rows, err := r.pool.Query(ctx, q, companyID)
	if err != nil {
		r.logger.Warn("failed to retrieve employees for the company", slog.String("error", err.Error()))
		return nil, models.ErrInternal
	}

	for rows.Next() {
		employee := models.EmployeeExtended{}
		if err = rows.Scan(
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
			&employee.EmployeeDepartment.Phone,
		); err != nil {
			r.logger.Warn("failed to retrieve employees for the company", slog.String("error", err.Error()))
			return nil, models.ErrInternal
		}
		employees = append(employees, employee)

	}

	if err = rows.Err(); err != nil {
		r.logger.Warn("failed to retrieve employees for the company", slog.String("error", err.Error()))
		return nil, models.ErrInternal
	}

	return employees, nil
}
