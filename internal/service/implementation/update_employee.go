package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) UpdateEmployee(ctx context.Context, employee models.EmployeeExtended) error {
	return s.repository.UpdateEmployee(ctx, employee)
}
