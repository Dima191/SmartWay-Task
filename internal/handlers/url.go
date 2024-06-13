package handlers

import "fmt"

var (
	EmployeesURL  = "/api/v1/employees"
	EmployeeIDKey = "employee_id"
	EmployeeURL   = fmt.Sprintf("%s/{%s}", EmployeesURL, EmployeeIDKey)

	CompaniesURL        = "/api/v1/companies"
	CompanyIDKey        = "company_id"
	CompanyURL          = fmt.Sprintf("%s/{%s}", CompaniesURL, CompanyIDKey)
	CompanyEmployeesURL = fmt.Sprintf("%s/employees", CompanyURL)

	DepartmentsURL         = fmt.Sprintf("%s/departments", CompanyURL)
	DepartmentIDKey        = "department_id"
	DepartmentURL          = fmt.Sprintf("%s/{%s}", DepartmentsURL, DepartmentIDKey)
	DepartmentEmployeesURL = fmt.Sprintf("%s/employees", DepartmentURL)
)
