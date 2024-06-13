package service

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

type Service interface {
	EmployeeByID(ctx context.Context, employeeID int) (employee models.EmployeeExtended, err error)

	AddEmployee(ctx context.Context, employee models.EmployeeBase) (employeeID int, err error)
	UpdateEmployee(ctx context.Context, employee models.EmployeeExtended) error
	FireEmployee(ctx context.Context, employeeID int) error

	CompanyEmployees(ctx context.Context, companyID int) (employees []models.EmployeeExtended, err error)
	Companies(ctx context.Context) (companies []models.Company, err error)
	AddCompany(ctx context.Context, company models.Company) (companyID int, err error)
	UpdateCompany(ctx context.Context, company models.Company) error

	DepartmentEmployees(ctx context.Context, companyID, departmentID int) (employees []models.EmployeeExtended, err error)
	Departments(ctx context.Context, companyID int) (departments []models.Department, err error)
	AddDepartment(ctx context.Context, department models.Department) (departmentID int, err error)
	UpdateDepartment(ctx context.Context, department models.Department) error
}
