package app

import (
	"context"
	"fmt"
	"github.com/Dima191/SmartWay-Task/internal/config"
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	handlerimpl "github.com/Dima191/SmartWay-Task/internal/handlers/implementation"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type App struct {
	httpServer http.Server
	router     *chi.Mux
	handler    handlers.Handler

	sp  *serviceProvider
	cfg *config.Config

	configPath string

	logger *slog.Logger
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider(a.cfg.DBConnectionString, a.logger)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.New(a.configPath)
	if err != nil {
		return err
	}

	a.cfg = cfg
	return nil
}

func (a *App) initRouter(_ context.Context) error {
	a.router = chi.NewRouter()
	return nil
}

func (a *App) initHandler(ctx context.Context) error {
	service, err := a.sp.Service(ctx)
	if err != nil {
		return err
	}

	a.handler = handlerimpl.New(service, a.logger)
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initRouter,
		a.initHandler,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runHTTPServer() error {
	if err := a.handler.Register(a.router); err != nil {
		a.logger.Error("failed to register handler", slog.String("error", err.Error()))
		return err
	}

	a.httpServer = http.Server{
		Addr:         fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler:      a.router,
		ReadTimeout:  a.cfg.ReadTimeout,
		WriteTimeout: a.cfg.WriteTimeout,
		IdleTimeout:  a.cfg.IdleTimeout,
	}

	a.logger.Info("app started", slog.String("address", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)))

	return a.httpServer.ListenAndServe()

}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) Stop() error {
	if err := a.httpServer.Close(); err != nil {
		return err
	}
	if a.sp.repository != nil {
		a.sp.repository.CloseConnection()
	}
	return nil
}

func New(ctx context.Context, configPath string, logger *slog.Logger) (*App, error) {
	logger.Info("creating new xds app")
	a := &App{
		configPath: configPath,
		logger:     logger,
	}

	logger.Info("initializing all dependencies for xds app")
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	logger.Info("all dependencies for xds app initialized")
	return a, nil
}
