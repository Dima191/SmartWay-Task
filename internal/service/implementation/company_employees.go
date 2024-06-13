package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) CompanyEmployees(ctx context.Context, companyID int) (employees []models.EmployeeExtended, err error) {
	return s.repository.CompanyEmployees(ctx, companyID)
}
