package serviceimpl

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
)

func (s *service) AddCompany(ctx context.Context, company models.Company) (companyID int, err error) {
	return s.repository.AddCompany(ctx, company)
}
