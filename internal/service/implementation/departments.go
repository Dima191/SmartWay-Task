package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) Departments(ctx context.Context, companyID int) (departments []models.Department, err error) {
	return s.repository.Departments(ctx, companyID)
}
