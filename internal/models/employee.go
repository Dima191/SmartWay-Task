package models

type EmployeeBase struct {
	ID               int      `json:"employee_id,omitempty"`
	FirstName        string   `json:"first_name"`
	SecondName       string   `json:"second_name"`
	Phone            string   `json:"employee_phone"`
	EmployeePassport Passport `json:"passport"`
}

type EmployeeExtended struct {
	EmployeeBase
	EmployeeCompany    Company    `json:"employee_company"`
	EmployeeDepartment Department `json:"department"`
}
