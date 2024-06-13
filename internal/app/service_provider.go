package app

import (
	"context"
	rep "github.com/Dima191/SmartWay-Task/internal/repository"
	postgresrep "github.com/Dima191/SmartWay-Task/internal/repository/postgres"
	srv "github.com/Dima191/SmartWay-Task/internal/service"
	serviceimpl "github.com/Dima191/SmartWay-Task/internal/service/implementation"
	"log/slog"
)

type serviceProvider struct {
	repository rep.Repository
	service    srv.Service

	connectionString string
	logger           *slog.Logger
}

func (sp *serviceProvider) Repository(ctx context.Context) (rep.Repository, error) {
	if sp.repository == nil {
		repository, err := postgresrep.New(ctx, sp.connectionString, sp.logger)
		if err != nil {
			return nil, err
		}
		sp.repository = repository
	}

	return sp.repository, nil
}

func (sp *serviceProvider) Service(ctx context.Context) (srv.Service, error) {
	if sp.service == nil {
		repository, err := sp.Repository(ctx)
		if err != nil {
			return nil, err
		}

		service := serviceimpl.New(repository, sp.logger)
		sp.service = service
	}
	return sp.service, nil
}

func newServiceProvider(connectionString string, logger *slog.Logger) *serviceProvider {
	sp := &serviceProvider{
		connectionString: connectionString,
		logger:           logger,
	}
	return sp
}
