package models

type Department struct {
	ID        int    `json:"department_id"`
	CompanyID int    `json:"company_id,omitempty"`
	Name      string `json:"department_name"`
	Phone     string `json:"department_phone"`
}
