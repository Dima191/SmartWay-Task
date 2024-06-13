package serviceimpl

import (
	rep "github.com/Dima191/SmartWay-Task/internal/repository"
	srv "github.com/Dima191/SmartWay-Task/internal/service"
	"log/slog"
)

type service struct {
	repository rep.Repository
	logger     *slog.Logger
}

func New(repository rep.Repository, logger *slog.Logger) srv.Service {
	s := &service{
		repository: repository,
		logger:     logger,
	}

	return s
}
