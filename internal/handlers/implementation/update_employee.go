package handlerimpl

import (
	"errors"
	"github.com/Dima191/SmartWay-Task/internal/handlers"
	"github.com/Dima191/SmartWay-Task/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *handler) UpdateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr := chi.URLParam(r, handlers.EmployeeIDKey)
		if employeeIDStr == "" {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}

		employeeID, err := strconv.Atoi(employeeIDStr)
		if err != nil {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}

		employee, err := h.service.EmployeeByID(r.Context(), employeeID)
		if err != nil {
			if errors.Is(err, models.ErrNoData) {
				http.Error(w, "can not update a non-existing employee", http.StatusBadRequest)
				return
			}
			http.Error(w, "failed to update employee", http.StatusInternalServerError)
			return
		}

		if err = render.DecodeJSON(r.Body, &employee); err != nil {
			h.logger.Error("failed to decode body", slog.String("error", err.Error()))
			http.Error(w, "failed to decode request", http.StatusBadRequest)
			return
		}

		employee.ID = employeeID

		if err = h.service.UpdateEmployee(r.Context(), employee); err != nil {
			switch {
			case errors.Is(err, models.ErrUnknownCompanyReference):
				http.Error(w, "failed to add employee to a non-existent company", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrUnknownDepartmentReference):
				http.Error(w, "failed to add employee to a non-existent department", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrAlreadyExists):
				http.Error(w, "failed to update employee. employee should be unique", http.StatusBadRequest)
				return
			case errors.Is(err, models.ErrCompanyDepProvide):
				http.Error(w, "failed to update employee. provide both values: company and department", http.StatusBadRequest)
				return
			}
			http.Error(w, "failed to update employee", http.StatusInternalServerError)
			return
		}
	}
}
