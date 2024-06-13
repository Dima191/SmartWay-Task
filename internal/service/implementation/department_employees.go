package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) DepartmentEmployees(ctx context.Context, companyID, departmentID int) (employees []models.EmployeeExtended, err error) {
	return s.repository.DepartmentEmployees(ctx, companyID, departmentID)
}
