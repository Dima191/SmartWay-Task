package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) UpdateDepartment(ctx context.Context, department models.Department) error {
	return s.repository.UpdateDepartment(ctx, department)
}
