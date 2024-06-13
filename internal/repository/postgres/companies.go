package postgresrep

import (
	"context"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"log/slog"
)

func (r *repository) Companies(ctx context.Context) (companies []models.Company, err error) {
	q := "select company_id, name from company"

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		r.logger.Warn("failed to get list of companies", slog.String("error", err.Error()))
		return nil, models.ErrInternal
	}

	for rows.Next() {
		company := models.Company{}
		if err = rows.Scan(&company.ID, &company.Name); err != nil {
			r.logger.Warn("failed to get list of companies", slog.String("error", err.Error()))
			return nil, models.ErrInternal
		}
		companies = append(companies, company)
	}

	return companies, nil
}
