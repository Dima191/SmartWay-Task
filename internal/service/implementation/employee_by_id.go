package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) EmployeeByID(ctx context.Context, employeeID int) (employee models.EmployeeExtended, err error) {
	return s.repository.EmployeeByID(ctx, employeeID)
}
