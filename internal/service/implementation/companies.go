package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) Companies(ctx context.Context) (companies []models.Company, err error) {
	return s.repository.Companies(ctx)
}
