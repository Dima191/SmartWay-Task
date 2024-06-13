package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) AddDepartment(ctx context.Context, department models.Department) (departmentID int, err error) {
	return s.repository.AddDepartment(ctx, department)
}
