package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/google/uuid"
)

func (s *service) AddEmployee(ctx context.Context, employee models.EmployeeBase) (employeeID int, err error) {
	employee.ID = int(uuid.New().ID())
	return s.repository.AddEmployee(ctx, employee)
}
