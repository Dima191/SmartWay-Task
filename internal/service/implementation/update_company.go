package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) UpdateCompany(ctx context.Context, company models.Company) error {
	return s.repository.UpdateCompany(ctx, company)
}
