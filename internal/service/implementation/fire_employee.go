package serviceimpl

import "context"

func (s *service) FireEmployee(ctx context.Context, employeeID int) error {
	return s.repository.FireEmployee(ctx, employeeID)
}
