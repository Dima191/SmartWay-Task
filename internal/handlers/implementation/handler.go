package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	srv "github.com/Dima191/SmartWay-Task/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

type handler struct {
	service srv.Service

	logger *slog.Logger
}

func (h *handler) Register(r http.Handler) error {
	router, ok := r.(*chi.Mux)
	if !ok {
		h.logger.Error("failed to convert http.Handler to chi.Mux")
		return errors.New("failed to convert http.Handler to chi.Mux")
	}

	middlewares := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.URLFormat,
	}

	router.Use(middlewares...)

	router.Post(handlers.EmployeesURL, h.AddEmployee())
	router.Patch(handlers.EmployeeURL, h.UpdateEmployee())
	router.Delete(handlers.EmployeeURL, h.FireEmployee())

	router.Get(handlers.CompaniesURL, h.Companies())
	router.Post(handlers.CompaniesURL, h.AddCompany())
	router.Patch(handlers.CompanyURL, h.UpdateCompany())

	router.Get(handlers.CompanyEmployeesURL, h.CompanyEmployees())

	router.Get(handlers.DepartmentsURL, h.Departments())
	router.Post(handlers.DepartmentsURL, h.AddDepartment())
	router.Patch(handlers.DepartmentURL, h.UpdateDepartment())

	router.Get(handlers.DepartmentEmployeesURL, h.DepartmentEmployees())

	return nil
}

func New(service srv.Service, logger *slog.Logger) handlers.Handler {
	h := &handler{
		service: service,
		logger:  logger,
	}

	return h
}
